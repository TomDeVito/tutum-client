package tutum

import (
	"fmt"
)

type NodeAPI interface {
	List() (*ListNodesResponse, error)
	Fetch(uuid string) (*Node, error)
}

type Node struct {
	DeployedDateTime  string    `json:"deployed_datetime,omitempty"`
	DestroyedDateTime string    `json:"destroyed_datetime,omitempty"`
	DockerExecDriver  string    `json:"docker_execdriver,omitempty"`
	DockerGraphDriver string    `json:"docker_graphdriver"`
	DockerVersion     string    `json:"docker_version,omitempty"`
	ExternalFQDN      string    `json:"external_fqdn,omitempty"`
	LastSeen          string    `json:"last_seen,omitempty"`
	NodeCluster       string    `json:"node_cluster,omitempty"`
	NodeType          string    `json:"node_type,omitempty"`
	PublicIP          string    `json:"public_ip,omitempty"`
	Region            string    `json:"region,omitempty"`
	ResourceUri       string    `json:"resource_uri,omitempty"`
	State             string    `json:"state,omitempty"`
	Tags              []NodeTag `json:"tags,omitempty"`
	Uuid              string    `json:"uuid,omitempty"`
}

type NodeTag struct {
	Name string `json:"name,omitempty"`
}

type nodeAPI struct {
	*baseAPI
}

func NewNodeAPI(baseURL, username, apiKey string) NodeAPI {
	return &nodeAPI{
		baseAPI: newBaseAPI(baseURL, username, apiKey),
	}
}

func (me *nodeAPI) List() (*ListNodesResponse, error) {
	response := &ListNodesResponse{}
	if err := me.get("/node/", response); err != nil {
		return nil, err
	}
	return response, nil
}

func (me *nodeAPI) Fetch(uuid string) (*Node, error) {
	response := &Node{}
	if err := me.get(fmt.Sprintf("/node/%s/", uuid), response); err != nil {
		return nil, err
	}
	return response, nil
}
