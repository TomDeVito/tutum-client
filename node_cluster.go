package tutum

type NodeClusterAPI interface {
	List() (*ListNodeClusterResponse, error)
}

type NodeCluster struct {
	Name  string `json:"name,omitempty"`
	Uuid  string `json:"uuid,omitempty"`
	State string `json:"state,omitempty"`
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
