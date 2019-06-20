package update

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth"

	"github.com/terra-project/core/update/plan"
	"github.com/terra-project/core/x/oracle"
)

// EndBlocker
func EndBlocker(
	ctx sdk.Context,
	accountKeeper auth.AccountKeeper,
	oracleKeeper oracle.Keeper) (resTags sdk.Tags) {

	plan.Update230000(ctx, accountKeeper, oracleKeeper)

	return
}