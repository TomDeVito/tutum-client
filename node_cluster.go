package tutum

import (
	"bytes"
	"encoding/json"
	"fmt"
)

type NodeClusterAPI interface {
	List() (*ListNodeClusterResponse, error)
	Create(nodeCluster *NodeCluster) (*NodeCluster, error)
	Deploy(uuid string) (*NodeCluster, error)
	Fetch(uuid string) (*NodeCluster, error)
}

type NodeCluster struct {
	Name           string   `json:"name,omitempty"`
	Uuid           string   `json:"uuid,omitempty"`
	Region         string   `json:"region,omitempty"`
	State          string   `json:"state,omitempty"`
	NodeType       string   `json:"node_type,omitempty"`
	Nodes          []string `json:"nodes,omitempty"`
	TargetNumNodes int32    `json:"target_num_nodes,omitempty"`
	Tags           []Tag    `json:"tags,omitempty"`
}

type nodeClusterAPI struct {
	*baseAPI
}

func NewNodeClusterAPI(baseURL, username, apiKey string) NodeClusterAPI {
	return &nodeClusterAPI{
		baseAPI: newBaseAPI(baseURL, username, apiKey),
	}
}

func (me *nodeClusterAPI) List() (*ListNodeClusterResponse, error) {
	response := &ListNodeClusterResponse{}
	if err := me.get("/nodecluster/", response); err != nil {
		return nil, err
	}
	return response, nil
}

func (me *nodeClusterAPI) Create(nodeCluster *NodeCluster) (*NodeCluster, error) {
	response := &NodeCluster{}

	nodeClusterBytes, err := json.Marshal(nodeCluster)
	if err != nil {
		return nil, err
	}

	if err := me.post("/nodecluster/", bytes.NewReader(nodeClusterBytes), response); err != nil {
		return nil, err
	}

	return response, nil
}

func (me *nodeClusterAPI) Deploy(uuid string) (*NodeCluster, error) {
	response := &NodeCluster{}

	if err := me.post(fmt.Sprintf("/nodecluster/%s/deploy/", uuid), nil, response); err != nil {
		return nil, err
	}

	return response, nil
}

func (me *nodeClusterAPI) Fetch(uuid string) (*NodeCluster, error) {
	response := &NodeCluster{}

	if err := me.get(fmt.Sprintf("/nodecluster/%s/", uuid), response); err != nil {
		return nil, err
	}

	return response, nil
}
