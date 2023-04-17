package inventory

type CustomResources struct {
	Velero             *Velero   `json:"velero"`
	KCIRocks           *KCIRocks `json:"kci_rocks"`
	RabbitMQ           *RabbitMQ `json:"rabbitmq"`
	Calico             *Calico   `json:"calico"`
	HasContour         bool      `json:"has_contour"`
	HasVelero          bool      `json:"has_velero"`
	HasKCIRocks        bool      `json:"has_kci_rocks"`
	HasRabbitMQ        bool      `json:"has_rabbit_mq"`
	HasCalico          bool      `json:"has_calico"`
	HasExternalSecrets bool      `json:"has_external_secrets"`
	HasCertManager     bool      `json:"has_cert_manager"`
	HasGitOpsToolkit   bool      `json:"has_git_ops_toolkit"`
	HasPrometheus      bool      `json:"has_prometheus"`
}

func NewCustomResources() *CustomResources {
	return &CustomResources{
		Velero:   NewVelero(),
		KCIRocks: NewKCIRocks(),
		RabbitMQ: NewRabbitMQ(),
		Calico:   NewCalico(),
	}
}
