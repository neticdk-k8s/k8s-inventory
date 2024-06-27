package inventory

import metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

type Namespaces = Set[*Namespace]

type Namespace struct {
	PartialObject

	Spec   NamespaceSpec   `json:"spec"`
	Status NamespaceStatus `json:"status"`
}

type NamespaceSpec struct{}
type NamespaceStatus struct{}

func NewNamespace() *Namespace {
	return &Namespace{
		PartialObject: PartialObject{
			TypeMeta: TypeMeta{
				Kind:         "Namespace",
				APIGroup:     "core",
				APIVersion:   "v1",
				ResourceType: "namespaces",
			},
			ObjectMeta: NewObjectMeta(metav1.ObjectMeta{}),
		},
		Spec:   NamespaceSpec{},
		Status: NamespaceStatus{},
	}
}
