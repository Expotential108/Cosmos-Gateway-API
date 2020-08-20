/*
 * Gaia-Lite for Cosmos
 *
 * A REST interface for state queries, transaction generation and broadcasting.
 *
 * API version: 3.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

// InlineResponse2005Value struct for InlineResponse2005Value
type InlineResponse2005Value struct {
	AccountNumber string    `json:"account_number,omitempty"`
	Address       string    `json:"address,omitempty"`
	Coins         []Coin    `json:"coins,omitempty"`
	PublicKey     PublicKey `json:"public_key,omitempty"`
	Sequence      string    `json:"sequence,omitempty"`
}
