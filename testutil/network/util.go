package network

import (
	"context"
	"encoding/json"
	"fmt"
	"math/rand"
	"path/filepath"
	"testing"
	"time"

	"github.com/spf13/pflag"
	"github.com/stretchr/testify/require"
	tmos "github.com/tendermint/tendermint/libs/os"
	"github.com/tendermint/tendermint/node"
	"github.com/tendermint/tendermint/p2p"
	pvm "github.com/tendermint/tendermint/privval"
	"github.com/tendermint/tendermint/proxy"
	"github.com/tendermint/tendermint/rpc/client/local"
	"github.com/tendermint/tendermint/types"
	tmtime "github.com/tendermint/tendermint/types/time"

	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/cosmos/cosmos-sdk/server/api"
	servergrpc "github.com/cosmos/cosmos-sdk/server/grpc"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	"github.com/cosmos/cosmos-sdk/x/genutil"
	genutiltypes "github.com/cosmos/cosmos-sdk/x/genutil/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	core "github.com/terra-money/core/types"
	oracletypes "github.com/terra-money/core/x/oracle/types"
)

func startInProcess(t *testing.T, cfg Config, val *Validator) error {
	logger := val.Ctx.Logger
	tmCfg := val.Ctx.Config
	tmCfg.Instrumentation.Prometheus = false

	nodeKey, err := p2p.LoadOrGenNodeKey(tmCfg.NodeKeyFile())
	if err != nil {
		return err
	}

	app := cfg.AppConstructor(*val)

	genDocProvider := node.DefaultGenesisDocProviderFunc(tmCfg)
	tmNode, err := node.NewNode(
		tmCfg,
		pvm.LoadOrGenFilePV(tmCfg.PrivValidatorKeyFile(), tmCfg.PrivValidatorStateFile()),
		nodeKey,
		proxy.NewLocalClientCreator(app),
		genDocProvider,
		node.DefaultDBProvider,
		node.DefaultMetricsProvider(tmCfg.Instrumentation),
		logger.With("module", val.Moniker),
	)
	if err != nil {
		return err
	}

	if err := tmNode.Start(); err != nil {
		return err
	}

	val.tmNode = tmNode

	if val.RPCAddress != "" {
		val.RPCClient = local.New(tmNode)
	}

	// We'll need a RPC client if the validator exposes a gRPC or REST endpoint.
	if val.APIAddress != "" || val.AppConfig.GRPC.Enable {
		val.ClientCtx = val.ClientCtx.
			WithClient(val.RPCClient)

		// Add the tx service in the gRPC router.
		app.RegisterTxService(val.ClientCtx)

		// Add the tendermint queries service in the gRPC router.
		app.RegisterTendermintService(val.ClientCtx)
	}

	setupOraclePseudoFeeder(t, cfg, val)

	if val.APIAddress != "" {
		apiSrv := api.New(val.ClientCtx, logger.With("module", "api-server"))
		app.RegisterAPIRoutes(apiSrv, val.AppConfig.API)

		errCh := make(chan error)

		go func() {
			if err := apiSrv.Start(*val.AppConfig); err != nil {
				errCh <- err
			}
		}()

		select {
		case err := <-errCh:
			return err
		case <-time.After(5 * time.Second): // assume server started successfully
		}

		val.api = apiSrv
	}

	if val.AppConfig.GRPC.Enable {
		grpcSrv, err := servergrpc.StartGRPCServer(val.ClientCtx, app, val.AppConfig.GRPC.Address)
		if err != nil {
			return err
		}

		val.grpc = grpcSrv

		if val.AppConfig.GRPCWeb.Enable {
			errCh1 := make(chan error)
			go func() {
				grpcWeb, err := servergrpc.StartGRPCWeb(grpcSrv, *val.AppConfig)
				if err != nil {
					errCh1 <- err
				}

				val.grpcWeb = grpcWeb
			}()
			select {
			case err := <-errCh1:
				return err
			case <-time.After(5 * time.Second): // assume server started successfully
			}

		}
	}

	return nil
}

type OraclePseudoFeederInfo struct {
	params      *oracletypes.QueryParamsResponse
	lastPrevote int64
	lastSalt    string
}

func setupOraclePseudoFeeder(t *testing.T, cfg Config, val *Validator) {
	ti := time.NewTicker(500 * time.Millisecond)
	done := make(chan bool)
	val.pseudoOracleFeederControl = &done
	rand.Seed(time.Now().UnixNano())

	go func() {
		info := OraclePseudoFeederInfo{
			lastPrevote: -1,
			lastSalt:    "",
		}
		for {
			select {
			case <-done:
				return
			case <-ti.C:
				if val.tmNode.BlockStore().Height() > 0 {
					// Grab oracle parameters if not initialized
					if info.params == nil {
						queryClient := oracletypes.NewQueryClient(val.ClientCtx)
						res, err := queryClient.Params(context.Background(), &oracletypes.QueryParamsRequest{})
						require.NoError(t, err)
						info.params = res
					}

					blockHeight := val.tmNode.BlockStore().Height()
					oracleVotePeriod := info.params.Params.VotePeriod
					currentVotePeriod := blockHeight / int64(oracleVotePeriod)

					if currentVotePeriod > info.lastPrevote {
						salt := fmt.Sprintf("%x", rand.Intn(99999))
						if len(salt) > 4 {
							salt = salt[len(salt)-4:]
						}

						var msgs []sdk.Msg

						exchangeRatesStr := fmt.Sprintf("1000.23%s,0.29%s,0.27%s", core.MicroKRWDenom, core.MicroUSDDenom, core.MicroSDRDenom)
						hash := oracletypes.GetAggregateVoteHash(salt, exchangeRatesStr, val.ValAddress)
						aggregateExchangeRatePrevoteMsg := oracletypes.NewMsgAggregateExchangeRatePrevote(hash, val.Address, val.ValAddress)

						if info.lastSalt != "" {
							aggregateExchangeRateVoteMsg :=
								oracletypes.NewMsgAggregateExchangeRateVote(info.lastSalt, exchangeRatesStr, val.Address, val.ValAddress)
							msgs = append(msgs, aggregateExchangeRateVoteMsg)
						}

						msgs = append(msgs, aggregateExchangeRatePrevoteMsg)

						builder := val.ClientCtx.TxConfig.NewTxBuilder()
						builder.SetMsgs(msgs...)
						builder.SetGasLimit(100000)

						txFactory := tx.Factory{}
						txf := tx.NewFactoryCLI(val.ClientCtx, &pflag.FlagSet{})
						acct, err := txf.AccountRetriever().GetAccount(val.ClientCtx, val.Address)
						require.NoError(t, err)
						txFactory = txFactory.
							WithChainID(cfg.ChainID).
							WithKeybase(val.ClientCtx.Keyring).
							WithTxConfig(cfg.TxConfig).
							WithSequence(acct.GetSequence())
						err = tx.Sign(txFactory, val.Moniker, builder, true)
						require.NoError(t, err)

						txBytes, err := val.ClientCtx.TxConfig.TxEncoder()(builder.GetTx())
						require.NoError(t, err)
						_, err = val.ClientCtx.BroadcastTxSync(txBytes)
						require.NoError(t, err)
						info.lastPrevote = currentVotePeriod
						info.lastSalt = salt
					}
				}
			}
		}
	}()
}

func collectGenFiles(cfg Config, vals []*Validator, outputDir string) error {
	genTime := tmtime.Now()

	for i := 0; i < cfg.NumValidators; i++ {
		tmCfg := vals[i].Ctx.Config

		nodeDir := filepath.Join(outputDir, vals[i].Moniker, "simd")
		gentxsDir := filepath.Join(outputDir, "gentxs")

		tmCfg.Moniker = vals[i].Moniker
		tmCfg.SetRoot(nodeDir)

		initCfg := genutiltypes.NewInitConfig(cfg.ChainID, gentxsDir, vals[i].NodeID, vals[i].PubKey)

		genFile := tmCfg.GenesisFile()
		genDoc, err := types.GenesisDocFromFile(genFile)
		if err != nil {
			return err
		}

		appState, err := genutil.GenAppStateFromConfig(cfg.Codec, cfg.TxConfig,
			tmCfg, initCfg, *genDoc, banktypes.GenesisBalancesIterator{})
		if err != nil {
			return err
		}

		// overwrite each validator's genesis file to have a canonical genesis time
		if err := genutil.ExportGenesisFileWithTime(genFile, cfg.ChainID, nil, appState, genTime); err != nil {
			return err
		}
	}

	return nil
}

func initGenFiles(cfg Config, genAccounts []authtypes.GenesisAccount, genBalances []banktypes.Balance, genFiles []string) error {

	// set the accounts in the genesis state
	var authGenState authtypes.GenesisState
	cfg.Codec.MustUnmarshalJSON(cfg.GenesisState[authtypes.ModuleName], &authGenState)

	accounts, err := authtypes.PackAccounts(genAccounts)
	if err != nil {
		return err
	}

	authGenState.Accounts = append(authGenState.Accounts, accounts...)
	cfg.GenesisState[authtypes.ModuleName] = cfg.Codec.MustMarshalJSON(&authGenState)

	// set the balances in the genesis state
	var bankGenState banktypes.GenesisState
	cfg.Codec.MustUnmarshalJSON(cfg.GenesisState[banktypes.ModuleName], &bankGenState)

	bankGenState.Balances = append(bankGenState.Balances, genBalances...)
	cfg.GenesisState[banktypes.ModuleName] = cfg.Codec.MustMarshalJSON(&bankGenState)

	appGenStateJSON, err := json.MarshalIndent(cfg.GenesisState, "", "  ")
	if err != nil {
		return err
	}

	genDoc := types.GenesisDoc{
		ChainID:    cfg.ChainID,
		AppState:   appGenStateJSON,
		Validators: nil,
	}

	// generate empty genesis files for each validator and save
	for i := 0; i < cfg.NumValidators; i++ {
		if err := genDoc.SaveAs(genFiles[i]); err != nil {
			return err
		}
	}

	return nil
}

func writeFile(name string, dir string, contents []byte) error {
	writePath := filepath.Join(dir)
	file := filepath.Join(writePath, name)

	err := tmos.EnsureDir(writePath, 0755)
	if err != nil {
		return err
	}

	err = tmos.WriteFile(file, contents, 0644)
	if err != nil {
		return err
	}

	return nil
}
