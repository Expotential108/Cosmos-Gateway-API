/*
 * Gaia-Lite for Cosmos
 *
 * A REST interface for state queries, transaction generation and broadcasting.
 *
 * API version: 3.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi
// BlockLastCommit struct for BlockLastCommit
type BlockLastCommit struct {
	BlockId BlockId `json:"block_id,omitempty"`
	Precommits []BlockLastCommitPrecommits `json:"precommits,omitempty"`
}
