package util

import sdk "github.com/cosmos/cosmos-sdk/types"

// IsWaitingForSoftfork returns whether current block
// height is bigger than reserved softfork block height
func IsWaitingForSoftfork(ctx sdk.Context, version uint8) bool {
	if version == 1 {
		// Expected time:
		// MAINNET
		// Fri Jan 01 2021 18:00:00 GMT+0900 (KST)
		// Fri Jan 01 2021 09:00:00 GMT+0000 (UTC)
		// Fri Jan 01 2021 01:00:00 GMT-0800 (PST)
		//
		// TEQUILA
		// Fri Nov 27 2020 12:00:00 GMT+0900 (KST)
		// Fri Nov 27 2020 03:00:00 GMT+0000 (UTC)
		// Fri Nov 26 2020 19:00:00 GMT-0800 (KST)

		return (ctx.ChainID() == "columbus-4" && ctx.BlockHeight() < 1200000) ||
			(ctx.ChainID() == "tequila-0004" && ctx.BlockHeight() < 1350000)
	} else if version == 2 {
		// Expected time:
		// MAINNET
		// Tue Mar 30 2021 18:00:00 GMT+0900 (KST)
		// Tue Mar 30 2021 09:00:00 GMT+0000 (UTC)
		// Tue Mar 30 2021 01:00:00 GMT-0800 (PST)
		//
		// TEQUILA
		// ASAP
		return (ctx.ChainID() == "columbus-4" && ctx.BlockHeight() < 2380000) ||
			(ctx.ChainID() == "tequila-0004" && ctx.BlockHeight() < 3052265)
	}

	return false
}

// IsSoftforkHeight return whether current block
// height is the targeted softfork height
func IsSoftforkHeight(ctx sdk.Context, _version uint8) bool {
	return (ctx.ChainID() == "columbus-4" && ctx.BlockHeight() == 1200000) ||
		(ctx.ChainID() == "tequila-0004" && ctx.BlockHeight() == 1350000)
}
