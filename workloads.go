package inventory

type (
	Spec   interface{}
	Status interface{}
)

type Workload struct {
	TypeMeta
	ObjectMeta

	Spec   Spec   `json:"spec"`
	Status Status `json:"status"`
}

func NewWorkloads() []*Workload {
	return make([]*Workload, 0)
}
