/*
 * Gaia-Lite for Cosmos
 *
 * A REST interface for state queries, transaction generation and broadcasting.
 *
 * API version: 3.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi
// PaginatedQueryTxs struct for PaginatedQueryTxs
type PaginatedQueryTxs struct {
	TotalCount float32 `json:"total_count,omitempty"`
	Count float32 `json:"count,omitempty"`
	PageNumber float32 `json:"page_number,omitempty"`
	PageTotal float32 `json:"page_total,omitempty"`
	Limit float32 `json:"limit,omitempty"`
	Txs []TxQuery `json:"txs,omitempty"`
}
