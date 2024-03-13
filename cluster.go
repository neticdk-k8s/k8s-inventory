package inventory

type Cluster struct {
	// FQDN
	FQDN string `json:"fqdn" validate:"omitempty,fqdn"`

	// Like `kubectl version -o json`
	Version     string `json:"semverVersion"`
	FullVersion string `json:"version"`
	GitCommit   string `json:"gitCommit"`
	BuildDate   string `json:"buildDate"`

	// Where/on what is the cluster running
	KubernetesProvider     string `json:"kubernetesProvider"`
	InfrastructureProvider string `json:"infrastructureProvider"`

	// Secure Cloud Stack Information
	ProviderName string `json:"providerName" validate:"required"`
	Name         string `json:"name"          validate:"required"`
}

func NewCluster() *Cluster {
	return &Cluster{}
}
