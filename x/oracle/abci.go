package oracle

import (
	"github.com/terra-project/core/x/oracle/internal/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/staking/exported"

	core "github.com/terra-project/core/types"
)

// EndBlocker is called at the end of every block
func EndBlocker(ctx sdk.Context, k Keeper) {
	params := k.GetParams(ctx)

	// Not yet time for a tally
	if !core.IsPeriodLastBlock(ctx, params.VotePeriod) {
		return
	}

	// Build valid votes counter map over all active validators
	validVotesCounterMap := make(map[string]int64)
	k.StakingKeeper.IterateValidators(ctx, func(index int64, validator exported.ValidatorI) bool {
		validVotesCounterMap[validator.GetOperator().String()] = int64(0)
		return false
	})

	whitelist := make(map[string]bool)
	for _, denom := range k.Whitelist(ctx) {
		whitelist[denom] = true
	}

	// Clear exchange rates
	for denom := range whitelist {
		k.DeleteLunaExchangeRate(ctx, denom)
	}

	winnerMap := make(map[string]types.Claim)

	// Organize votes to ballot by denom
	voteMap := k.OrganizeBallotByDenom(ctx)

	// Build abstain validator map
	abstainOperatorMap := make(map[string]bool)
	abstainBallot := voteMap[core.MicroLunaDenom]
	for _, vote := range abstainBallot {
		abstainOperatorMap[vote.Voter.String()] = true
	}

	// Remove abstain key
	delete(voteMap, core.MicroLunaDenom)

	// Iterate through ballots and update exchange rates; drop if not enough votes have been achieved.
	for denom, ballot := range voteMap {

		// If denom is not in the whitelist, or the ballot for it has failed, then skip
		if _, exists := whitelist[denom]; !exists || !ballotIsPassing(ctx, ballot, k) {
			continue
		}

		// Get weighted median exchange rates, and faithful respondants
		ballotMedian, ballotWinningClaims := tally(ctx, ballot, k)

		// Set the exchange rate
		k.SetLunaExchangeRate(ctx, denom, ballotMedian)

		// Collect claims of ballot winners
		for _, ballotWinningClaim := range ballotWinningClaims {
			key := ballotWinningClaim.Recipient.String()
			prevClaim, exists := winnerMap[key]
			if !exists {
				winnerMap[key] = ballotWinningClaim
			} else {
				prevClaim.Weight += ballotWinningClaim.Weight
				winnerMap[key] = prevClaim
			}

			// Increase valid votes counter
			validVotesCounterMap[key]++
		}

		// Emit abci events
		ctx.EventManager().EmitEvent(
			sdk.NewEvent(types.EventTypeExchangeRateUpdate,
				sdk.NewAttribute(types.AttributeKeyDenom, denom),
				sdk.NewAttribute(types.AttributeKeyExchangeRate, ballotMedian.String()),
			),
		)
	}

	//---------------------------
	// Do miss counting & slashing

	whitelistLen := int64(len(whitelist))
	for operatorBechAddr, count := range validVotesCounterMap {
		// Skip abstain operators & valid voters
		if (count == whitelistLen) || (count == 0 && abstainOperatorMap[operatorBechAddr]) {
			continue
		}

		// Increase miss counter
		operator, _ := sdk.ValAddressFromBech32(operatorBechAddr) // error never occur
		k.SetMissCounter(ctx, operator, k.GetMissCounter(ctx, operator)+1)
	}

	// Reset miss counters of all validators at the last block of slash window
	if core.IsPeriodLastBlock(ctx, params.VotePeriod*params.SlashWindow) {
		k.SlashAndResetMissCounters(ctx)
	}

	// Distribute rewards to ballot winners
	k.RewardBallotWinners(ctx, winnerMap)

	// Clear the ballot
	clearBallots(k, ctx, params)

	return
}

// clearBallots clears all tallied prevotes and votes from the store
func clearBallots(k Keeper, ctx sdk.Context, params Params) {
	// Clear all prevotes
	k.IterateExchangeRatePrevotes(ctx, func(prevote types.ExchangeRatePrevote) (stop bool) {
		if ctx.BlockHeight() > prevote.SubmitBlock+params.VotePeriod {
			k.DeleteLunaExchangeRatePrevote(ctx, prevote)
		}

		return false
	})

	// Clear all votes
	k.IterateExchangeRateVotes(ctx, func(vote types.ExchangeRateVote) (stop bool) {
		k.DeleteLunaExchangeRateVote(ctx, vote)
		return false
	})
}
