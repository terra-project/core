// nolint
// autogenerated code using github.com/rigelrozanski/multitool
// aliases generated for the following subdirectories:
// ALIASGEN: github.com/terra-money/core/x/market/internal/types/
// ALIASGEN: github.com/terra-money/core/x/market/internal/keeper/
package market

import (
	"github.com/terra-money/core/x/market/internal/keeper"
	"github.com/terra-money/core/x/market/internal/types"
)

const (
	ModuleName          = types.ModuleName
	StoreKey            = types.StoreKey
	RouterKey           = types.RouterKey
	QuerierRoute        = types.QuerierRoute
	DefaultParamspace   = types.DefaultParamspace
	QuerySwap           = types.QuerySwap
	QueryTerraPoolDelta = types.QueryTerraPoolDelta
	QueryParameters     = types.QueryParameters
)

var (
	// functions aliases
	RegisterCodec       = types.RegisterCodec
	ErrNoEffectivePrice = types.ErrNoEffectivePrice
	ErrInvalidOfferCoin = types.ErrInvalidOfferCoin
	ErrRecursiveSwap    = types.ErrRecursiveSwap
	NewGenesisState     = types.NewGenesisState
	DefaultGenesisState = types.DefaultGenesisState
	ValidateGenesis     = types.ValidateGenesis
	NewMsgSwap          = types.NewMsgSwap
	NewMsgSwapSend      = types.NewMsgSwapSend
	DefaultParams       = types.DefaultParams
	NewQuerySwapParams  = types.NewQuerySwapParams
	ParamKeyTable       = types.ParamKeyTable
	NewKeeper           = keeper.NewKeeper
	NewQuerier          = keeper.NewQuerier

	// variable aliases
	ModuleCdc                       = types.ModuleCdc
	TerraPoolDeltaKey               = types.TerraPoolDeltaKey
	ParamStoreKeyBasePool           = types.ParamStoreKeyBasePool
	ParamStoreKeyPoolRecoveryPeriod = types.ParamStoreKeyPoolRecoveryPeriod
	ParamStoreKeyMinSpread          = types.ParamStoreKeyMinStabilitySpread
	DefaultBasePool                 = types.DefaultBasePool
	DefaultPoolRecoveryPeriod       = types.DefaultPoolRecoveryPeriod
	DefaultMinSpread                = types.DefaultMinStabilitySpread
)

type (
	SupplyKeeper    = types.SupplyKeeper
	OracleKeeper    = types.OracleKeeper
	GenesisState    = types.GenesisState
	MsgSwap         = types.MsgSwap
	MsgSwapSend     = types.MsgSwapSend
	Params          = types.Params
	QuerySwapParams = types.QuerySwapParams
	Keeper          = keeper.Keeper
)
