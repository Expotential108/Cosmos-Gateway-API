/*
 * Gaia-Lite for Cosmos
 *
 * A REST interface for state queries, transaction generation and broadcasting.
 *
 * API version: 3.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi
// InlineResponse2008 struct for InlineResponse2008
type InlineResponse2008 struct {
	MinDeposit []Coin `json:"min_deposit,omitempty"`
	MaxDepositPeriod string `json:"max_deposit_period,omitempty"`
}