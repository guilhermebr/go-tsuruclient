/*
 * Tsuru
 *
 * Open source, extensible and Docker-based Platform as a Service (PaaS)
 *
 * API version: 1.6.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package tsuru

// App plan
type Plan struct {
	Name string `json:"name,omitempty"`

	Memory int64 `json:"memory,omitempty"`

	Swap int64 `json:"swap,omitempty"`

	Cpushare int32 `json:"cpushare,omitempty"`

	Default_ bool `json:"default,omitempty"`
}