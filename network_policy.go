package inventory

import (
	"encoding/json"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type NetworkPolicies = Set[*NetworkPolicy]

type NetworkPolicy struct {
	PartialObject

	Spec NetworkPolicySpec `json:"spec"`
}

type NetworkPolicyPort struct {
	Protocol *string      `json:"protocol,omitempty"`
	Port     *IntOrString `json:"port,omitempty"`
	EndPort  *int32       `json:"endPort,omitempty"`
}

func (npp NetworkPolicyPort) MarshalJSON() ([]byte, error) {
	if npp.Port == nil {
		return json.Marshal(struct {
			Protocol *string      `json:"protocol,omitempty"`
			Port     *IntOrString `json:"port,omitempty"`
			EndPort  *int32       `json:"endPort,omitempty"`
		}{
			Protocol: npp.Protocol,
			Port:     nil,
			EndPort:  npp.EndPort,
		})
	}
	if npp.Port.Type == 0 { // int
		return json.Marshal(struct {
			Protocol *string `json:"protocol,omitempty"`
			Port     *int32  `json:"port,omitempty"`
			EndPort  *int32  `json:"endPort,omitempty"`
		}{
			Protocol: npp.Protocol,
			Port:     &npp.Port.IntVal,
			EndPort:  npp.EndPort,
		})
	}
	// str
	return json.Marshal(struct {
		Protocol *string `json:"protocol,omitempty"`
		Port     *string `json:"port,omitempty"`
		EndPort  *int32  `json:"endPort,omitempty"`
	}{
		Protocol: npp.Protocol,
		Port:     &npp.Port.StrVal,
		EndPort:  npp.EndPort,
	})
}

func (npp *NetworkPolicyPort) UnmarshalJSON(data []byte) error {
	if string(data) == "null" || string(data) == `""` {
		return nil
	}

	var jsonNPP struct {
		Protocol *string     `json:"protocol,omitempty"`
		Port     interface{} `json:"port,omitempty"`
		EndPort  *int32      `json:"endPort,omitempty"`
	}

	if err := json.Unmarshal(data, &jsonNPP); err != nil {
		return err
	}

	var port *IntOrString

	portVal := jsonNPP.Port

	switch v := portVal.(type) {
	case float64:
		port = &IntOrString{
			Type:   0,
			IntVal: int32(v),
		}
	case string:
		port = &IntOrString{
			Type:   1,
			StrVal: v,
		}
	}

	*npp = NetworkPolicyPort{
		Protocol: jsonNPP.Protocol,
		Port:     port,
		EndPort:  jsonNPP.EndPort,
	}

	return nil
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

func NewNetworkPolicy() *NetworkPolicy {
	return &NetworkPolicy{
		PartialObject: PartialObject{
			TypeMeta: TypeMeta{
				Kind:         "NetworkPolicy",
				APIGroup:     "networking",
				APIVersion:   "v1",
				ResourceType: "networkpolicies",
			},
			ObjectMeta: NewObjectMeta(metav1.ObjectMeta{}),
		},
		Spec: NetworkPolicySpec{
			PodSelector: LabelSelector{},
			Ingress:     make([]NetworkPolicyIngressRule, 0),
			Egress:      make([]NetworkPolicyEgressRule, 0),
		},
	}
}
