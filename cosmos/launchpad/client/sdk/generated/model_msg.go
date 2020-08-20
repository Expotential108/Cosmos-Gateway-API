/*
 * Gaia-Lite for Cosmos
 *
 * A REST interface for state queries, transaction generation and broadcasting.
 *
 * API version: 3.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi
// Msg struct for Msg
type Msg struct {
	Type string `json:"type,omitempty"`
	Value MsgValue `json:"value,omitempty"`
}
