/*
 * Tsuru
 *
 * Open source, extensible and Docker-based Platform as a Service (PaaS)
 *
 * API version: 1.6
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package tsuru

type VolumeBindId struct {
	// App the volume is bound to.
	App interface{} `json:"app,omitempty"`
	// Volume mountpoint.
	Mountpoint interface{} `json:"mountpoint,omitempty"`
	// Volume name.
	Volume interface{} `json:"volume,omitempty"`
}
