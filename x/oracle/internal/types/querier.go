package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// Defines the prefix of each query path
const (
	QueryParameters       = "parameters"
	QueryExchangeRate     = "exchangeRate"
	QueryExchangeRates    = "exchangeRates"
	QueryActives          = "actives"
	QueryPrevotes         = "prevotes"
	QueryVotes            = "votes"
	QueryFeederDelegation = "feederDelegation"
)

// QueryExchangeRateParams defines the params for the following queries:
// - 'custom/oracle/exchangeRate'
type QueryExchangeRateParams struct {
	Denom string
}

// NewQueryExchangeRateParams returns params for exchangeRate query
func NewQueryExchangeRateParams(denom string) QueryExchangeRateParams {
	return QueryExchangeRateParams{denom}
}

// QueryPrevotesParams defines the params for the following queries:
// - 'custom/oracle/prevotes'
type QueryPrevotesParams struct {
	Voter sdk.ValAddress
	Denom string
}

// NewQueryPrevotesParams returns params for exchangeRate prevotes query
func NewQueryPrevotesParams(voter sdk.ValAddress, denom string) QueryPrevotesParams {
	return QueryPrevotesParams{voter, denom}
}

// QueryVotesParams defines the params for the following queries:
// - 'custom/oracle/votes'
type QueryVotesParams struct {
	Voter sdk.ValAddress
	Denom string
}

// NewQueryVotesParams returns params for exchangeRate votes query
func NewQueryVotesParams(voter sdk.ValAddress, denom string) QueryVotesParams {
	return QueryVotesParams{voter, denom}
}

// QueryFeederDelegationParams defeins the params for the following queries:
// - 'custom/oracle/feederDelegation'
type QueryFeederDelegationParams struct {
	Validator sdk.ValAddress
}

// NewQueryFeederDelegationParams returns params for feeder delegation query
func NewQueryFeederDelegationParams(validator sdk.ValAddress) QueryFeederDelegationParams {
	return QueryFeederDelegationParams{validator}
}
