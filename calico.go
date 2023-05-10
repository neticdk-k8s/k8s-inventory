package inventory

type Calico struct {
	TypeMeta
	ObjectMeta

	Spec CalicoSpec `json:"spec" db:"spec"`
}

type CalicoSpec struct {
	Version string `json:"version"`
}

func NewCalico() *Calico {
	return &Calico{
		ObjectMeta: NewObjectMeta(),
		Spec:       CalicoSpec{},
	}
}
