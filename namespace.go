package inventory

type Namespaces []*Namespace

type Namespace struct {
	TypeMeta
	ObjectMeta

	Spec   NamespaceSpec   `json:"spec" db:"spec"`
	Status NamespaceStatus `json:"status" db:"status"`
}

type NamespaceSpec struct{}
type NamespaceStatus struct{}

func NewNamespace() *Namespace {
	return &Namespace{
		TypeMeta: TypeMeta{
			Kind:         "Namespace",
			APIGroup:     "core",
			APIVersion:   "v1",
			ResourceType: "namespaces",
		},
		ObjectMeta: NewObjectMeta(),
		Spec:       NamespaceSpec{},
		Status:     NamespaceStatus{},
	}
}
