package inventory

type Cluster struct {
	// Like `kubectl version -o json`
	Version     string `json:"semver_version" db:"version"`
	FullVersion string `json:"version" db:"full_version"`
	GitCommit   string `json:"git_commit" db:"git_commit"`
	BuildDate   string `json:"build_date" db:"build_date"`

	// Where/on what is the cluster running
	KubernetesProvider     string `json:"kubernetes_provider" db:"kubernetes_provider"`
	InfrastructureProvider string `json:"infrastructure_provider" db:"infrastructure_provider"`

	// Secure Cloud Stack Information
	OperatorName                  string `json:"operator_name" validate:"required" db:"operator_name"`
	OperatorSubscriptionID        int    `json:"operator_subscription_id" validate:"required" db:"operator_subscription_id"`
	ProviderName                  string `json:"provider_name" validate:"required" db:"provider_name"`
	ProviderSubscriptionID        int    `json:"provider_subscription_id" validate:"required" db:"provider_subscription_id"`
	CustomerName                  string `json:"customer_name" validate:"required" db:"customer_name"`
	CustomerID                    int    `json:"customer_id" validate:"required" db:"customer_id"`
	BillingSubject                string `json:"billing_subject" validate:"oneof=operator provider tenant special" db:"billing_subject"`
	BillingGranularity            string `json:"billing_granularity" validate:"oneof=cluster tenant" db:"billing_granularity"`
	Name                          string `json:"name" validate:"required" db:"name"`
	FQDN                          string `json:"fqdn" validate:"omitempty,fqdn" db:"fqdn"`
	ClusterType                   string `json:"cluster_type" validate:"required,oneof=dedicated shared custom" db:"cluster_type"`
	Description                   string `json:"description" db:"description"`
	ResilienceZone                string `json:"resilience_zone" db:"resilience_zone"`
	InfrastructureEnvironmentType string `json:"infrastructure_environment_type" validate:"required,oneof=dedicated shared hybrid other" db:"infrastructure_environment_type"`
	EnvironmentName               string `json:"environment_name" db:"environment_name"`
	HasTechnicalOperations        bool   `json:"has_technical_operations" db:"has_technical_operations"`
	HasTechnicalManagement        bool   `json:"has_technical_management" db:"has_technical_management"`
	HasApplicationOperations      bool   `json:"has_application_operations" db:"has_application_operations"`
	HasApplicationManagement      bool   `json:"has_application_management" db:"has_application_management"`
	HasCapacityManagement         bool   `json:"has_capacity_management" db:"has_capacity_management"`
	HasCustomOperations           bool   `json:"has_custom_operations" db:"has_custom_operations"`
	CustomOperationsURL           string `json:"custom_operations_url" validate:"required_if=HasCustomOperations true,omitempty,url" db:"custom_operations_url"`
}

func NewCluster() *Cluster {
	return &Cluster{}
}
