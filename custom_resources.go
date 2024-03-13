package inventory

type CustomResources struct {
	Velero             *Velero                   `json:"velero"`
	KCIRocks           *KCIRocks                 `json:"kciRocks"`
	RabbitMQ           *RabbitMQ                 `json:"rabbitMQ"`
	CalicoCluster      *CalicoClusterInformation `json:"calicoCluster"`
	HasContour         bool                      `json:"hasContour"`
	HasVelero          bool                      `json:"hasVelero"`
	HasKCIRocks        bool                      `json:"hasKCIRocks"`
	HasRabbitMQ        bool                      `json:"hasRabbitMQ"`
	HasCalico          bool                      `json:"hasCalico"`
	HasExternalSecrets bool                      `json:"hasExternalSecrets"`
	HasCertManager     bool                      `json:"hasCertManager"`
	HasGitOpsToolkit   bool                      `json:"hasGitOpsToolkit"`
	HasPrometheus      bool                      `json:"hasPrometheus"`
}

func NewCustomResources() *CustomResources {
	return &CustomResources{
		Velero:        NewVelero(),
		KCIRocks:      NewKCIRocks(),
		RabbitMQ:      NewRabbitMQ(),
		CalicoCluster: NewCalicoClusterInformation(),
	}
}
