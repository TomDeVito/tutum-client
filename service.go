package tutum

type ServiceAPI interface {
	List() (*ListServicesResponse, error)
	//Fetch(name string) (*Service, error)
}

type Service struct {
	Uuid                 string             `json:"uuid"`
	ResourceUri          string             `json:"resource_uri"`
	ImageName            string             `json:"image_name"`
	ImageTag             string             `json:"image_tag"`
	Name                 string             `json:"name"`
	PublicDns            string             `json:"public_dns"`
	State                string             `json:"state"`
	Synchronized         bool               `json:"synchronized"`
	DeployedDateTime     string             `json:"deployed_datetime"`
	StartedDateTime      string             `json:"started_datetime"`
	StoppedDateTime      string             `json:"stopped_datetime"`
	DestroyedDateTime    string             `json:"destroyed_datetime"`
	TargetNumContainers  int32              `json:"target_num_containers"`
	CurrentNumContainers int32              `json:"current_num_containers"`
	RunningNumContainers int32              `json:"running_num_containers"`
	StoppedNumContainers int32              `json:"stopped_num_containers"`
	Stack                string             `json:"stack"`
	Containers           []string           `json:"containers"`
	ContainerPorts       []ContainerPorts   `json:"container_ports"`
	ContainerEnvVars     []ContainerEnvVars `json:"container_envvars"`
	Entrypoint           string             `json:"entrypoint"`
	RunCommand           string             `json:"run_command"`
	SequentialDeployment bool               `json:"sequential_deployment"`
	CpuShares            int32              `json:"cpu_shares"`
	Memory               int32              `json:"memory"`
	LinkedFromService    []LinkService      `json:"linked_from_service"`
	LinkedToService      []LinkService      `json:"linked_to_service"`
	AutoRestart          string             `json:"autorestart"`
	AutoDestroy          string             `json:"autodestroy"`
	AutoRedeploy         bool               `json:"autoredeploy"`
	Roles                []string           `json:"roles"`
	Privileged           bool               `json:"privileged"`
	DeploymentStrategy   string             `json:"deployment_strategy"`
	Tags                 []string           `json:"tags"`
	//Bindings
	//LinkVariables
}

type ContainerPorts struct {
	EndpointUri string `json:"endpoint_uri"`
	InnerPort   int32  `json:"inner_port"`
	OuterPort   int32  `json:"outer_port"`
	PortName    string `json:"port_name"`
	Protocol    string `json:"protocol"`
	Published   bool   `json:"published"`
}

type ContainerEnvVars struct {
	Key   string `json:"key"`
	Value string `json:"Value"`
}

type LinkService struct {
	FromService string `json:"from_service"`
	Name        string `json:"name"`
	ToService   string `json:"to_service"`
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
