package inventory

type Calico struct {
	TypeMeta
	ObjectMeta

	Spec   CalicoSpec   `json:"spec" db:"spec"`
	Status CalicoStatus `json:"status" db:"status"`
}

type CalicoSpec struct {
	Version string `json:"version"`
}

type CalicoStatus struct{}

func NewCalico() *Calico {
	return &Calico{
		ObjectMeta: NewObjectMeta(),
		Spec:       CalicoSpec{},
		Status:     CalicoStatus{},
	}
}
