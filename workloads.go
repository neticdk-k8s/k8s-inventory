package inventory

type Spec interface{}
type Status interface{}

type Workload struct {
	TypeMeta
	ObjectMeta

	Spec   Spec   `json:"spec" db:"spec"`
	Status Status `json:"status" db:"status"`
}

func NewWorkloads() []*Workload {
	return make([]*Workload, 0)
}
