package rest_test

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	grpctypes "github.com/cosmos/cosmos-sdk/types/grpc"
	"github.com/terra-money/core/testutil"
	"github.com/terra-money/core/x/market/types"
)

func (s *IntegrationTestSuite) TestGRPCQueryParameters() {
	val := s.network.Validators[0]
	resp, err := testutil.GetRequestWithHeaders(
		fmt.Sprintf("%s/terra/market/v1beta1/params", val.APIAddress),
		map[string]string{
			grpctypes.GRPCBlockHeightHeader: "1",
		},
	)
	s.Require().NoError(err)
	a := types.QueryParamsResponse{}
	err = val.ClientCtx.JSONCodec.UnmarshalJSON(resp, &a)
	s.Require().NoError(err)
	b := types.QueryParamsResponse{}
	b.Params.BurnBasePool, _ = sdk.NewDecFromStr("7000000000000.000000000000000000")
	b.Params.MintBasePool, _ = sdk.NewDecFromStr("7000000000000.000000000000000000")
	b.Params.MinStabilitySpread, _ = sdk.NewDecFromStr("0.005000000000000000")
	b.Params.PoolRecoveryPeriod = 200
	s.Require().Equal(a.String(), b.String())
}

func (s *IntegrationTestSuite) TestGRPCQueryMintPoolDeltaFn() {
	val := s.network.Validators[0]
	resp, err := testutil.GetRequestWithHeaders(
		fmt.Sprintf("%s/terra/market/v1beta1/mint_pool_delta", val.APIAddress),
		map[string]string{
			grpctypes.GRPCBlockHeightHeader: "1",
		},
	)
	s.Require().NoError(err)
	a := types.QueryMintPoolDeltaResponse{}
	err = val.ClientCtx.JSONCodec.UnmarshalJSON(resp, &a)
	s.Require().NoError(err)
	b := types.QueryMintPoolDeltaResponse{}
	b.MintPoolDelta, _ = sdk.NewDecFromStr("0")
	s.Require().Equal(a.String(), b.String())
}

func (s *IntegrationTestSuite) TestGRPCQueryBurnPoolDeltaFn() {
	val := s.network.Validators[0]
	resp, err := testutil.GetRequestWithHeaders(
		fmt.Sprintf("%s/terra/market/v1beta1/burn_pool_delta", val.APIAddress),
		map[string]string{
			grpctypes.GRPCBlockHeightHeader: "1",
		},
	)
	s.Require().NoError(err)
	a := types.QueryBurnPoolDeltaResponse{}
	err = val.ClientCtx.JSONCodec.UnmarshalJSON(resp, &a)
	s.Require().NoError(err)
	b := types.QueryBurnPoolDeltaResponse{}
	b.BurnPoolDelta, _ = sdk.NewDecFromStr("0")
	s.Require().Equal(a.String(), b.String())
}

func (s *IntegrationTestSuite) TestGRPCQuerySwapFn() {
	val := s.network.Validators[0]
	resp, err := testutil.GetRequestWithHeaders(
		fmt.Sprintf("%s/terra/market/v1beta1/swap?offer_coin=1000000uluna&ask_denom=usdr", val.APIAddress),
		map[string]string{
			grpctypes.GRPCBlockHeightHeader: "16",
		},
	)
	s.Require().NoError(err)
	a := types.QuerySwapResponse{}
	err = val.ClientCtx.JSONCodec.UnmarshalJSON(resp, &a)
	s.Require().NoError(err)
	b := types.QuerySwapResponse{
		ReturnCoin: sdk.NewCoin("usdr", sdk.NewInt(268650)),
	}
	s.Require().Equal(a.String(), b.String())
}
