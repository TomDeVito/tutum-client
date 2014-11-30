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
