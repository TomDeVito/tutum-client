package tutum

type TutumAPI interface {
	Action() ActionAPI
	Container() ContainerAPI
	Node() NodeAPI
	NodeCluster() NodeClusterAPI
	NodeType() NodeTypeAPI
	Provider() ProviderAPI
	Region() RegionAPI
	Service() ServiceAPI
}

type tutumAPI struct {
	username string
	apiKey   string
	baseURL  string
}

func NewTutumAPI(username, apiKey, baseURL string) TutumAPI {
	return &tutumAPI{
		username: username,
		apiKey:   apiKey,
		baseURL:  baseURL,
	}
}

func (me *tutumAPI) Action() ActionAPI {
	return struct{}{}
}

func (me *tutumAPI) Container() ContainerAPI {
	return NewContainerAPI(me.baseURL, me.username, me.apiKey)
}

func (me *tutumAPI) Node() NodeAPI {
	return NewNodeAPI(me.baseURL, me.username, me.apiKey)
}

func (me *tutumAPI) NodeCluster() NodeClusterAPI {
	return NewNodeClusterAPI(me.baseURL, me.username, me.apiKey)
}

func (me *tutumAPI) NodeType() NodeTypeAPI {
	return struct{}{}
}

func (me *tutumAPI) Provider() ProviderAPI {
	return NewProviderAPI(me.baseURL, me.username, me.apiKey)
}

func (me *tutumAPI) Region() RegionAPI {
	return struct{}{}
}

func (me *tutumAPI) Service() ServiceAPI {
	return NewServiceAPI(me.baseURL, me.username, me.apiKey)
}
