package inventory

type Tenants []*Tenant

type Tenant struct {
	Name                     string `json:"name" validate:"required" db:"name"`
	SubscriptionID           string `json:"subscription_id" validate:"required" db:"subscription_id"`
	Namespace                string `json:"namespace" validate:"required" db:"namespace"`
	BusinessUnitID           string `json:"business_unit_id" db:"business_unit_id"`
	HasApplicationManagement bool   `json:"has_application_management" db:"has_application_management"`
	HasApplicationOperations bool   `json:"has_application_operations" db:"has_application_operations"`
	HasCapacityManagement    bool   `json:"has_capacity_management" db:"has_capacity_management"`
}

func NewTenant() *Tenant {
	return &Tenant{}
}
