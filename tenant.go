package inventory

type Tenants []*Tenant

type Tenant struct {
	Name                     string `json:"name" validate:"required"`
	SubscriptionID           int    `json:"subscription_id" validate:"required"`
	Namespace                string `json:"namespace" validate:"required"`
	BusinessUnitID           string `json:"business_unit_id"`
	HasApplicationManagement bool   `json:"has_application_management"`
	HasApplicationOperations bool   `json:"has_application_operations"`
	HasCapacityManagement    bool   `json:"has_capacity_management"`
}

func NewTenant() *Tenant {
	return &Tenant{}
}
