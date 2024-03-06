package inventory

import metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

type NetworkPolicies []*NetworkPolicy

type NetworkPolicy struct {
	TypeMeta
	ObjectMeta

	Spec NetworkPolicySpec `json:"spec" db:"spec"`
}

type (
	NetworkPolicySpec struct{}
)

func NewNetworkPolicy() *NetworkPolicy {
	return &NetworkPolicy{
		TypeMeta: TypeMeta{
			Kind:         "NetworkPolicy",
			APIGroup:     "core",
			APIVersion:   "v1",
			ResourceType: "namespaces",
		},
		ObjectMeta: NewObjectMeta(metav1.ObjectMeta{}),
		Spec:       NetworkPolicySpec{},
	}
}
