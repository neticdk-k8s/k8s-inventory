package inventory

type Cluster struct {
	// FQDN
	FQDN string `json:"fqdn" validate:"omitempty,fqdn" db:"fqdn"`

	// Like `kubectl version -o json`
	Version     string `json:"semver_version" db:"version"`
	FullVersion string `json:"version"        db:"full_version"`
	GitCommit   string `json:"git_commit"     db:"git_commit"`
	BuildDate   string `json:"build_date"     db:"build_date"`

	// Where/on what is the cluster running
	KubernetesProvider     string `json:"kubernetes_provider"     db:"kubernetes_provider"`
	InfrastructureProvider string `json:"infrastructure_provider" db:"infrastructure_provider"`

	// Secure Cloud Stack Information
	ProviderName string `json:"provider_name" validate:"required" db:"provider_name"`
	Name         string `json:"name"          validate:"required" db:"name"`
}

func NewCluster() *Cluster {
	return &Cluster{}
}
