/*
 * Tsuru
 *
 * Open source, extensible and Docker-based Platform as a Service (PaaS)
 *
 * API version: 1.6
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package tsuru

type Unit struct {
	Id          interface{} `json:"id,omitempty"`
	Name        interface{} `json:"name,omitempty"`
	Appname     interface{} `json:"appname,omitempty"`
	Processname interface{} `json:"processname,omitempty"`
	Type_       interface{} `json:"type,omitempty"`
	Ip          interface{} `json:"ip,omitempty"`
	Status      interface{} `json:"status,omitempty"`
	Address     interface{} `json:"address,omitempty"`
}
