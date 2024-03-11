package inventory

import (
	"database/sql/driver"
	"encoding/json"
	"errors"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type NetworkPolicies []*NetworkPolicy

type NetworkPolicy struct {
	TypeMeta
	ObjectMeta

	Spec NetworkPolicySpec `json:"spec" db:"spec"`
}

type NetworkPolicyPort struct {
	Protocol *string `json:"protocol,omitempty"`
	Port     *string `json:"port,omitempty"`
	EndPort  *string `json:"endPort,omitempty"`
}

type IPBlock struct {
	CIDR   string   `json:"cidr,omitempty"`
	Except []string `json:"except,omitempty"`
}
type NetworkPolicyPeer struct {
	PodSelector       *LabelSelector `json:"podSelector,omitempty"`
	NamespaceSelector *LabelSelector `json:"namespaceSelector,omitempty"`
	IPBlock           *IPBlock       `json:"ipBlock,omitempty"`
}

type NetworkPolicyIngressRule struct {
	Ports []NetworkPolicyPort `json:"ports,omitempty"`
	From  []NetworkPolicyPeer `json:"from,omitempty"`
}

type NetworkPolicyEgressRule struct {
	Ports []NetworkPolicyPort `json:"ports,omitempty"`
	To    []NetworkPolicyPeer `json:"from,omitempty"`
}

type NetworkPolicySpec struct {
	PodSelector LabelSelector              `json:"podSelector"`
	Ingress     []NetworkPolicyIngressRule `json:"ingress,omitempty"`
	Egress      []NetworkPolicyEgressRule  `json:"egress,omitempty"`
	PolicyTypes []string                   `json:"policyTypes,omitempty"`
}

func (nps *NetworkPolicySpec) Value() (driver.Value, error) {
	bytes, err := json.Marshal(nps)
	return bytes, err
}

func (nps *NetworkPolicySpec) Scan(val interface{}) error {
	b, ok := val.([]byte)
	if !ok {
		return errors.New("type assertion to []byte failed")
	}

	return json.Unmarshal(b, &nps)
}

func NewNetworkPolicy() *NetworkPolicy {
	return &NetworkPolicy{
		TypeMeta: TypeMeta{
			Kind:         "NetworkPolicy",
			APIGroup:     "networking",
			APIVersion:   "v1",
			ResourceType: "networkpolicies",
		},
		ObjectMeta: NewObjectMeta(metav1.ObjectMeta{}),
		Spec: NetworkPolicySpec{
			PodSelector: LabelSelector{},
			Ingress:     make([]NetworkPolicyIngressRule, 0),
			Egress:      make([]NetworkPolicyEgressRule, 0),
		},
	}
}
