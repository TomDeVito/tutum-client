package tutum

type MetaResponse struct {
	Limit      int         `json:"limit"`
	Next       interface{} `json:"next"`
	Previous   interface{} `json:"previous"`
	Offset     int         `json:"offset"`
	TotalCount int         `json:"total_count"`
}

type ListProvidersResponse struct {
	Meta    *MetaResponse `json:"meta"`
	Objects []Provider    `json:"objects"`
}

type ListServicesResponse struct {
	Meta    *MetaResponse `json:"meta"`
	Objects []Service     `json:"objects"`
}

type ListContainersResponse struct {
	Meta    *MetaResponse `json:"meta"`
	Objects []Container   `json:"objects"`
}

type ListNodesResponse struct {
	Meta    *MetaResponse `json:"meta"`
	Objects []Node        `json:"objects"`
}

type ListNodeClusterResponse struct {
	Meta    *MetaResponse `json:"meta"`
	Objects []NodeCluster `json:"objects"`
}
