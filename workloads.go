package inventory

type (
	Spec   interface{}
	Status interface{}
)

type Workload struct {
	PartialObject

	Spec   Spec   `json:"spec"`
	Status Status `json:"status"`

	RootOwner *RootOwner `json:"rootOwner,omitempty"`
}

func NewWorkloads() []*Workload {
	return make([]*Workload, 0)
}
