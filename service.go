package tutum

import (
	"bytes"
	"encoding/json"
	"fmt"
)

type ServiceAPI interface {
	List() (*ListServicesResponse, error)
	Fetch(uuid string) (*Service, error)
	Create(service *Service) (*Service, error)
}

type Service struct {
	Uuid                 string             `json:"uuid,omitempty"`
	ResourceUri          string             `json:"resource_uri,omitempty"`
	Image                string             `json:"image,omitempty"`
	ImageName            string             `json:"image_name,omitempty"`
	ImageTag             string             `json:"image_tag,omitempty"`
	Name                 string             `json:"name,omitempty"`
	PublicDns            string             `json:"public_dns,omitempty"`
	State                string             `json:"state,omitempty"`
	Synchronized         bool               `json:"synchronized,omitempty"`
	DeployedDateTime     string             `json:"deployed_datetime,omitempty"`
	StartedDateTime      string             `json:"started_datetime,omitempty"`
	StoppedDateTime      string             `json:"stopped_datetime,omitempty"`
	DestroyedDateTime    string             `json:"destroyed_datetime,omitempty"`
	TargetNumContainers  int32              `json:"target_num_containers,omitempty"`
	CurrentNumContainers int32              `json:"current_num_containers,omitempty"`
	RunningNumContainers int32              `json:"running_num_containers,omitempty"`
	StoppedNumContainers int32              `json:"stopped_num_containers,omitempty"`
	Stack                string             `json:"stack,omitempty"`
	Containers           []string           `json:"containers,omitempty"`
	ContainerPorts       []ContainerPorts   `json:"container_ports,omitempty"`
	ContainerEnvVars     []ContainerEnvVars `json:"container_envvars,omitempty"`
	Entrypoint           string             `json:"entrypoint,omitempty"`
	RunCommand           string             `json:"run_command,omitempty"`
	SequentialDeployment bool               `json:"sequential_deployment,omitempty"`
	CpuShares            int32              `json:"cpu_shares,omitempty"`
	Memory               int32              `json:"memory,omitempty"`
	LinkedFromService    []LinkService      `json:"linked_from_service,omitempty"`
	LinkedToService      []LinkService      `json:"linked_to_service,omitempty"`
	AutoRestart          string             `json:"autorestart,omitempty"`
	AutoDestroy          string             `json:"autodestroy,omitempty"`
	AutoRedeploy         bool               `json:"autoredeploy,omitempty"`
	Roles                []string           `json:"roles,omitempty"`
	Privileged           bool               `json:"privileged,omitempty"`
	DeploymentStrategy   string             `json:"deployment_strategy,omitempty"`
	Tags                 []string           `json:"tags,omitempty"`
	//Bindings
	//LinkVariables
}

type ContainerPorts struct {
	EndpointUri string `json:"endpoint_uri,omitempty"`
	InnerPort   int32  `json:"inner_port,omitempty"`
	OuterPort   int32  `json:"outer_port,omitempty"`
	PortName    string `json:"port_name,omitempty"`
	Protocol    string `json:"protocol,omitempty"`
	Published   bool   `json:"published,omitempty"`
}

type ContainerEnvVars struct {
	Key   string `json:"key,omitempty"`
	Value string `json:"Value,omitempty"`
}

type LinkService struct {
	FromService string `json:"from_service,omitempty"`
	Name        string `json:"name,omitempty"`
	ToService   string `json:"to_service,omitempty"`
}

type serviceAPI struct {
	*baseAPI
}

func NewServiceAPI(baseURL, username, apiKey string) ServiceAPI {
	return &serviceAPI{
		baseAPI: newBaseAPI(baseURL, username, apiKey),
	}
}

func (me *serviceAPI) List() (*ListServicesResponse, error) {
	response := &ListServicesResponse{}
	if err := me.get("/service/", response); err != nil {
		return nil, err
	}
	return response, nil
}

func (me *serviceAPI) Fetch(uuid string) (*Service, error) {
	response := &Service{}
	if err := me.get(fmt.Sprintf("/service/%s/", uuid), response); err != nil {
		return nil, err
	}
	return response, nil
}

func (me *serviceAPI) Create(service *Service) (*Service, error) {
	response := &Service{}

	serviceBytes, err := json.Marshal(service)
	if err != nil {
		return nil, err
	}

	if err := me.post("/service/", bytes.NewReader(serviceBytes), response); err != nil {
		return nil, err
	}

	return response, nil
}
