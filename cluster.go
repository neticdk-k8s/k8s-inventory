package inventory

type Cluster struct {
	// Like `kubectl version -o json`
	Version     string `json:"semver_version"`
	FullVersion string `json:"version"`
	GitCommit   string `json:"git_commit"`
	BuildDate   string `json:"build_date"`

	// Where/on what is the cluster running
	ClusterProvider        string `json:"cluster_provider"`
	InfrastructureProvider string `json:"infrastructure_provider"`

	// Secure Cloud Stack Information
	OperatorName                  string `json:"operator_name" validate:"required"`
	OperatorSubscriptionID        int    `json:"operator_subscription_id" validate:"required"`
	ProviderName                  string `json:"provider_name" validate:"required"`
	ProviderSubscriptionID        int    `json:"provider_subscription_id" validate:"required"`
	CustomerName                  string `json:"customer_name" validate:"required"`
	CustomerID                    int    `json:"customer_id" validate:"required"`
	BillingSubject                string `json:"billing_subject" validate:"oneof=operator provider tenant special"`
	BillingGranularity            string `json:"billing_granularity" validate:"oneof=cluster tenant"`
	ClusterName                   string `json:"cluster_name" validate:"required"`
	ClusterFQDN                   string `json:"cluster_fqdn" validate:"omitempty,fqdn"`
	ClusterType                   string `json:"cluster_type" validate:"required,oneof=dedicated shared custom"`
	ClusterDescription            string `json:"cluster_description"`
	ClusterResilienceZone         string `json:"cluster_resilience_zone"`
	InfrastructureEnvironmentType string `json:"infrastructure_environment_type" validate:"required,oneof=dedicated shared hybrid other"`
	EnvironmentName               string `json:"environment_name"`
	HasTechnicalOperations        bool   `json:"has_technical_operations"`
	HasTechnicalManagement        bool   `json:"has_technical_management"`
	HasApplicationOperations      bool   `json:"has_application_operations"`
	HasApplicationManagement      bool   `json:"has_application_management"`
	HasCapacityManagement         bool   `json:"has_capacity_management"`
	HasCustomOperations           bool   `json:"has_custom_operations"`
	CustomOperationsURL           string `json:"custom_operations_url" validate:"required_if=HasCustomOperations true,omitempty,url"`
}

func NewCluster() *Cluster {
	return &Cluster{}
}
