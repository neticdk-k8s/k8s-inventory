package inventory

type Namespaces []*Namespace

type Namespace struct {
	TypeMeta
	ObjectMeta
}

func NewNamespace() *Namespace {
	return &Namespace{
		ObjectMeta: NewObjectMeta(),
	}
}
