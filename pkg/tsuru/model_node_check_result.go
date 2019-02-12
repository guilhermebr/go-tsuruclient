/*
 * Tsuru
 *
 * Open source, extensible and Docker-based Platform as a Service (PaaS)
 *
 * API version: 1.6
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package tsuru

type NodeCheckResult struct {
	Name       interface{} `json:"name,omitempty"`
	Err        interface{} `json:"err,omitempty"`
	Successful interface{} `json:"successful,omitempty"`
}
