/*
 * Gaia-Lite for Cosmos
 *
 * A REST interface for state queries, transaction generation and broadcasting.
 *
 * API version: 3.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package openapi
// InlineResponse200ApplicationVersion struct for InlineResponse200ApplicationVersion
type InlineResponse200ApplicationVersion struct {
	BuildTags string `json:"build_tags,omitempty"`
	ClientName string `json:"client_name,omitempty"`
	Commit string `json:"commit,omitempty"`
	Go string `json:"go,omitempty"`
	Name string `json:"name,omitempty"`
	ServerName string `json:"server_name,omitempty"`
	Version string `json:"version,omitempty"`
}
