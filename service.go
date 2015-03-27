package tutum

import (
	"bytes"
	"encoding/json"
	"fmt"
)

type ServiceAPI interface {
	List() (*ListServicesResponse, error)
	Fetch(uuid string) (*Service, error)
	Start(uuid string) (*Service, error)
	Create(service *Service) (*Service, error)
	Terminate(uuid string) (*Service, error)
}

type Service struct {
	AutoDestroy          string            `json:"autodestroy,omitempty"`
	AutoRedeploy         bool              `json:"autoredeploy,omitempty"`
	AutoRestart          string            `json:"autorestart,omitempty"`
	Bindings             []Binding         `json:"bindings,omitempty"`
	ContainerEnvVars     []ContainerEnvVar `json:"container_envvars,omitempty"`
	ContainerPorts       []ContainerPorts  `json:"container_ports,omitempty"`
	Containers           []string          `json:"containers,omitempty"`
	CpuShares            int32             `json:"cpu_shares,omitempty"`
	CurrentNumContainers int32             `json:"current_num_containers,omitempty"`
	DeployedDateTime     string            `json:"deployed_datetime,omitempty"`
	DeploymentStrategy   string            `json:"deployment_strategy,omitempty"`
	DestroyedDateTime    string            `json:"destroyed_datetime,omitempty"`
	Entrypoint           string            `json:"entrypoint,omitempty"`
	Image                string            `json:"image,omitempty"`
	ImageName            string            `json:"image_name,omitempty"`
	ImageTag             string            `json:"image_tag,omitempty"`
	LinkedFromService    []LinkService     `json:"linked_from_service,omitempty"`
	LinkedToService      []LinkService     `json:"linked_to_service,omitempty"`
	Memory               int32             `json:"memory,omitempty"`
	Name                 string            `json:"name,omitempty"`
	Privileged           bool              `json:"privileged,omitempty"`
	PublicDns            string            `json:"public_dns,omitempty"`
	ResourceUri          string            `json:"resource_uri,omitempty"`
	Roles                []string          `json:"roles,omitempty"`
	RunCommand           string            `json:"run_command,omitempty"`
	RunningNumContainers int32             `json:"running_num_containers,omitempty"`
	SequentialDeployment bool              `json:"sequential_deployment,omitempty"`
	Stack                string            `json:"stack,omitempty"`
	StartedDateTime      string            `json:"started_datetime,omitempty"`
	State                string            `json:"state,omitempty"`
	StoppedDateTime      string            `json:"stopped_datetime,omitempty"`
	StoppedNumContainers int32             `json:"stopped_num_containers,omitempty"`
	Synchronized         bool              `json:"synchronized,omitempty"`
	Tags                 []Tag             `json:"tags,omitempty"`
	TargetNumContainers  int32             `json:"target_num_containers,omitempty"`
	Uuid                 string            `json:"uuid,omitempty"`
}

type ContainerPorts struct {
	EndpointUri string `json:"endpoint_uri,omitempty"`
	InnerPort   int32  `json:"inner_port,omitempty"`
	OuterPort   int32  `json:"outer_port,omitempty"`
	PortName    string `json:"port_name,omitempty"`
	Protocol    string `json:"protocol,omitempty"`
	Published   bool   `json:"published,omitempty"`
}

type ContainerEnvVar struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

type Binding struct {
	HostPath      string `json:"host_path"`
	ContainerPath string `json:"container_path"`
	Rewritable    bool   `json:"rewritable"`
	VolumesFrom   string `json:"volumesFrom"`
}

type LinkService struct {
	FromService string `json:"from_service,omitempty"`
	Name        string `json:"name,omitempty"`
	ToService   string `json:"to_service,omitempty"`
}

type Tag struct {
	Name        string `json:"name,omitempty"`
	ResourceUri string `json:"resource_uri,omitempty"`
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

func (me *serviceAPI) Start(uuid string) (*Service, error) {
	response := &Service{}
	if err := me.post(fmt.Sprintf("/service/%s/start/", uuid), nil, response); err != nil {
		return nil, err
	}

	return response, nil
}

func (me *serviceAPI) Terminate(uuid string) (*Service, error) {
	response := &Service{}

	if err := me.del(fmt.Sprintf("/service/%s/", uuid), response); err != nil {
		return nil, err
	}

	return response, nil
}
