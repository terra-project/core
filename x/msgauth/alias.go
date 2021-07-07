// nolint
// autogenerated code using github.com/rigelrozanski/multitool
// aliases generated for the following subdirectories:
// ALIASGEN: github.com/terra-money/core/x/msgauth/internal/keeper
// ALIASGEN: github.com/terra-money/core/x/msgauth/internal/types
package msgauth

import (
	"github.com/terra-money/core/x/msgauth/internal/keeper"
	"github.com/terra-money/core/x/msgauth/internal/types"
)

const (
	EventGrantAuthorization    = types.EventGrantAuthorization
	EventRevokeAuthorization   = types.EventRevokeAuthorization
	EventExecuteAuthorization  = types.EventExecuteAuthorization
	AttributeKeyGrantType      = types.AttributeKeyGrantType
	AttributeKeyGranteeAddress = types.AttributeKeyGranteeAddress
	AttributeKeyGranterAddress = types.AttributeKeyGranterAddress
	AttributeValueCategory     = types.AttributeValueCategory
	ModuleName                 = types.ModuleName
	StoreKey                   = types.StoreKey
	RouterKey                  = types.RouterKey
	QuerierRoute               = types.QuerierRoute
	QueryGrant                 = types.QueryGrant
	QueryGrants                = types.QueryGrants
)

var (
	// functions aliases
	NewKeeper                    = keeper.NewKeeper
	NewQuerier                   = keeper.NewQuerier
	SetupTestInput               = keeper.SetupTestInput
	NewAuthorizationGrant        = types.NewAuthorizationGrant
	RegisterCodec                = types.RegisterCodec
	NewGenericAuthorization      = types.NewGenericAuthorization
	NewGenesisState              = types.NewGenesisState
	DefaultGenesisState          = types.DefaultGenesisState
	ValidateGenesis              = types.ValidateGenesis
	GetGrantKey                  = types.GetGrantKey
	GetGrantTimeKey              = types.GetGrantTimeKey
	ExtractAddressesFromGrantKey = types.ExtractAddressesFromGrantKey
	NewMsgGrantAuthorization     = types.NewMsgGrantAuthorization
	NewMsgRevokeAuthorization    = types.NewMsgRevokeAuthorization
	NewMsgExecAuthorized         = types.NewMsgExecAuthorized
	NewQueryGrantParams          = types.NewQueryGrantParams
	NewQueryGrantsParams         = types.NewQueryGrantsParams
	NewSendAuthorization         = types.NewSendAuthorization

	// variable aliases
	ModuleCdc        = types.ModuleCdc
	ErrInvalidPeriod = types.ErrInvalidPeriod
	GrantKey         = types.GrantKey
	GrantQueueKey    = types.GrantQueueKey
)

type (
	Keeper                 = keeper.Keeper
	Authorization          = types.Authorization
	AuthorizationGrant     = types.AuthorizationGrant
	GGMPair                = types.GGMPair
	GenericAuthorization   = types.GenericAuthorization
	AuthorizationEntry     = types.AuthorizationEntry
	GenesisState           = types.GenesisState
	MsgGrantAuthorization  = types.MsgGrantAuthorization
	MsgRevokeAuthorization = types.MsgRevokeAuthorization
	MsgExecAuthorized      = types.MsgExecAuthorized
	QueryGrantParams       = types.QueryGrantParams
	QueryGrantsParams      = types.QueryGrantsParams
	SendAuthorization      = types.SendAuthorization
)
