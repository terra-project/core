package rest_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/suite"
	"github.com/terra-money/core/testutil/network"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/rest"
	"github.com/terra-money/core/x/market/types"
)

type IntegrationTestSuite struct {
	suite.Suite

	cfg     network.Config
	network *network.Network
}

func (s *IntegrationTestSuite) SetupSuite() {
	s.T().Log("setting up integration test suite")

	cfg := network.DefaultConfig()
	genesisState := cfg.GenesisState
	cfg.NumValidators = 1

	var marketGenesis types.GenesisState
	s.Require().NoError(cfg.Codec.UnmarshalJSON(genesisState[types.ModuleName], &marketGenesis))
	var err error
	marketGenesis.Params.BurnBasePool, _ = sdk.NewDecFromStr("7000000000000.000000000000000000")
	marketGenesis.Params.MintBasePool, _ = sdk.NewDecFromStr("7000000000000.000000000000000000")
	marketGenesis.Params.MinStabilitySpread, _ = sdk.NewDecFromStr("0.005000000000000000")
	marketGenesis.Params.PoolRecoveryPeriod = 200
	marketGenesisBz, err := cfg.Codec.MarshalJSON(&marketGenesis)
	s.Require().NoError(err)
	genesisState[types.ModuleName] = marketGenesisBz
	cfg.GenesisState = genesisState

	s.cfg = cfg
	s.network = network.New(s.T(), cfg)

	_, err = s.network.WaitForHeight(2)
	s.Require().NoError(err)
}

func (s *IntegrationTestSuite) TearDownSuite() {
	s.T().Log("tearing down integration test suite")
	s.network.Cleanup()
}

func (s *IntegrationTestSuite) TestParametersFn() {
	val := s.network.Validators[0]
	resp, err := rest.GetRequest(fmt.Sprintf("%s/market/parameters?height=1", val.APIAddress))
	s.Require().NoError(err)
	bz, err := rest.ParseResponseWithHeight(val.ClientCtx.LegacyAmino, resp)
	s.Require().NoError(err)
	a := types.Params{}
	b := types.Params{}
	b.BurnBasePool, _ = sdk.NewDecFromStr("7000000000000.000000000000000000")
	b.MintBasePool, _ = sdk.NewDecFromStr("7000000000000.000000000000000000")
	b.MinStabilitySpread, _ = sdk.NewDecFromStr("0.005000000000000000")
	b.PoolRecoveryPeriod = 200
	s.Require().NoError(val.ClientCtx.LegacyAmino.UnmarshalJSON(bz, &a))
	s.Require().Equal(a.String(), b.String())
}

func (s *IntegrationTestSuite) TestQueryMintPoolDeltaFn() {
	val := s.network.Validators[0]
	resp, err := rest.GetRequest(fmt.Sprintf("%s/market/mint_pool_delta?height=1", val.APIAddress))
	s.Require().NoError(err)
	bz, err := rest.ParseResponseWithHeight(val.ClientCtx.LegacyAmino, resp)
	s.Require().NoError(err)
	a := sdk.Dec{}
	b := sdk.NewDec(0)
	s.Require().NoError(val.ClientCtx.LegacyAmino.UnmarshalJSON(bz, &a))
	s.Require().Equal(a.String(), b.String())
}

func (s *IntegrationTestSuite) TestQueryBurnPoolDeltaFn() {
	val := s.network.Validators[0]
	resp, err := rest.GetRequest(fmt.Sprintf("%s/market/burn_pool_delta?height=1", val.APIAddress))
	s.Require().NoError(err)
	bz, err := rest.ParseResponseWithHeight(val.ClientCtx.LegacyAmino, resp)
	s.Require().NoError(err)
	a := sdk.Dec{}
	b := sdk.NewDec(0)
	s.Require().NoError(val.ClientCtx.LegacyAmino.UnmarshalJSON(bz, &a))
	s.Require().Equal(a.String(), b.String())
}

/*func (s *IntegrationTestSuite) TestQuerySwapFn() {
	val := s.network.Validators[0]
	resp, err := rest.GetRequest(
		fmt.Sprintf("%s/market/swap?offer_coin=1000000uluna&ask_denom=usdr&height=1", val.APIAddress),
	)
	s.Require().NoError(err)
	bz, err := rest.ParseResponseWithHeight(val.ClientCtx.LegacyAmino, resp)
	s.Require().NoError(err)
	s.T().Log(bz)
	a := sdk.Coin{}
	//b := sdk.NewDec(0)
	s.Require().NoError(val.ClientCtx.LegacyAmino.UnmarshalJSON(bz, &a))
	//s.Require().Equal(a.String(), b.String())
	s.T().Log(a.String())
}*/

func TestIntegrationTestSuite(t *testing.T) {
	suite.Run(t, new(IntegrationTestSuite))
}
