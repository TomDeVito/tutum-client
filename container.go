package tutum

type ContainerAPI interface {
	List() (*ListContainersResponse, error)
}

type Container struct {
	Application          string    `json:"application"`
	Autodestroy          string    `json:"autodestroy"`
	Autoreplace          string    `json:"autoreplace"`
	Autorestart          string    `json:"autorestart"`
	ContainerPorts       []CCPInfo `json:"container_ports"`
	ContainerSize        string    `json:"container_size"`
	CurrentNumContainers int       `json:"current_num_containers"`
	DeployedDatetime     string    `json:"deployed_datetime"`
	DestroyedDatetime    string    `json:"destroyed_datetime"`
	Entrypoint           string    `json:"entrypoint"`
	ExitCode             int       `json:"exit_code"`
	ExitCodeMessage      string    `json:"exit_code_message"`
	ImageName            string    `json:"image_name"`
	ImageTag             string    `json:"image_tag"`
	Name                 string    `json:"name"`
	PublicDns            string    `json:"public_dns"`
	ResourceUri          string    `json:"resource_uri"`
	RunCommand           string    `json:"run_command"`
	StartedDatetime      string    `json:"started_datetime"`
	State                string    `json:"state"`
	StoppedDatetime      string    `json:"stopped_datetime"`
	UniqueName           string    `json:"unique_name"`
	Uuid                 string    `json:"uuid"`
}

type CCPInfo struct {
	Container string `json:"container"`
	InnerPort int    `json:"inner_port"`
	OuterPort int    `json:"outer_port"`
	Protocol  string `json:"protocol"`
}

type containerAPI struct {
	*baseAPI
}

func NewContainerAPI(baseURL, username, apiKey string) ContainerAPI {
	return &containerAPI{
		baseAPI: newBaseAPI(baseURL, username, apiKey),
	}
}

func (me *containerAPI) List() (*ListContainersResponse, error) {
	response := &ListContainersResponse{}
	if err := me.get("/container/", response); err != nil {
		return nil, err
	}
	return response, nil
}
