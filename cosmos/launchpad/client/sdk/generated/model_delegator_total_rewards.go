/*
 * Gaia-Lite for Cosmos
 *
 * A REST interface for state queries, transaction generation and broadcasting.
 *
 * API version: 3.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

// DelegatorTotalRewards struct for DelegatorTotalRewards
type DelegatorTotalRewards struct {
	Rewards []DelegationDelegatorReward `json:"rewards,omitempty"`
	Total   []Coin                      `json:"total,omitempty"`
}
