package tutum

import (
	"fmt"
)

type ContainerAPI interface {
	List() (*ListContainersResponse, error)
	Fetch(uuid string) (*Container, error)
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
	Node                 string    `json:"node"`
	PublicDns            string    `json:"public_dns"`
	ResourceUri          string    `json:"resource_uri"`
	RunCommand           string    `json:"run_command"`
	Service              string    `json:"service"`
	StartedDatetime      string    `json:"started_datetime"`
	State                string    `json:"state"`
	StoppedDatetime      string    `json:"stopped_datetime"`
	UniqueName           string    `json:"unique_name"`
	Uuid                 string    `json:"uuid"`
}

type CCPInfo struct {
	EndpointUri string `json:"endpoint_uri,omitempty"`
	InnerPort   int32  `json:"inner_port,omitempty"`
	OuterPort   int32  `json:"outer_port,omitempty"`
	PortName    string `json:"port_name,omitempty"`
	Protocol    string `json:"protocol,omitempty"`
	Published   bool   `json:"published,omitempty"`
	////	Container string `json:"container"`
	//	InnerPort int    `json:"inner_port"`
	//	OuterPort int    `json:"outer_port"`
	//	Protocol  string `json:"protocol"`
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

func (me *containerAPI) Fetch(uuid string) (*Container, error) {
	response := &Container{}

	if err := me.get(fmt.Sprintf("/container/%s/", uuid), response); err != nil {
		return nil, err
	}
	return response, nil
}
