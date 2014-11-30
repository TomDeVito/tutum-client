package tutum

import "fmt"

type ProviderAPI interface {
	List() (*ListProvidersResponse, error)
	Fetch(name string) (*Provider, error)
}

type Provider struct {
	Available   bool     `json:"available"`
	Label       string   `json:"label"`
	Name        string   `json:"name"`
	Regions     []string `json:"regions"`
	ResourceURI string   `json:"resource_uri"`
}

type providerAPI struct {
	*baseAPI
}

func NewProviderAPI(baseURL, username, apiKey string) ProviderAPI {
	return &providerAPI{
		baseAPI: newBaseAPI(baseURL, username, apiKey),
	}
}

func (me *providerAPI) List() (*ListProvidersResponse, error) {
	response := &ListProvidersResponse{}
	if err := me.get("/provider/", response); err != nil {
		return nil, err
	}
	return response, nil
}

func (me *providerAPI) Fetch(name string) (*Provider, error) {
	response := &Provider{}
	if err := me.get(fmt.Sprintf("/provider/%s/", name), response); err != nil {
		return nil, err
	}
	return response, nil
}
