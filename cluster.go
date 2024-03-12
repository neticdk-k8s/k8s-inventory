package inventory

type Cluster struct {
	// FQDN
	FQDN string `json:"fqdn" validate:"omitempty,fqdn"`

	// Like `kubectl version -o json`
	Version     string `json:"semver_version"`
	FullVersion string `json:"version"`
	GitCommit   string `json:"git_commit"`
	BuildDate   string `json:"build_date"`

	// Where/on what is the cluster running
	KubernetesProvider     string `json:"kubernetes_provider"`
	InfrastructureProvider string `json:"infrastructure_provider"`

	// Secure Cloud Stack Information
	ProviderName string `json:"provider_name" validate:"required"`
	Name         string `json:"name"          validate:"required"`
}

func NewCluster() *Cluster {
	return &Cluster{}
}
