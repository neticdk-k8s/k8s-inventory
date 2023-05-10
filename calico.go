package inventory

type Calico struct {
	TypeMeta
	ObjectMeta

	Version string `json:"version"`
}

func NewCalico() *Calico {
	return &Calico{
		ObjectMeta: NewObjectMeta(),
	}
}
