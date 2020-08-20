/*
 * Gaia-Lite for Cosmos
 *
 * A REST interface for state queries, transaction generation and broadcasting.
 *
 * API version: 3.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi

// Vote struct for Vote
type Vote struct {
	Voter      string `json:"voter,omitempty"`
	ProposalId string `json:"proposal_id,omitempty"`
	Option     string `json:"option,omitempty"`
}
