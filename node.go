package inventory

import (
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type Nodes []*Node

type Node struct {
	TypeMeta
	ObjectMeta

	Spec   NodeSpec   `json:"spec"`
	Status NodeStatus `json:"status"`

	Role           string `json:"role"`
	IsControlPlane bool   `json:"isControlPlane"`
	TopologyRegion string `json:"topologyRegion"`
	TopologyZone   string `json:"topologyZone"`
	Provider       string `json:"provider"`

	KubeProxyVersion        string `json:"kubeProxyVersion"`
	KubeletVersion          string `json:"kubeletVersion"`
	KernelVersion           string `json:"kernelVersion"`
	CRIName                 string `json:"criName"`
	CRIVersion              string `json:"criVersion"`
	ContainerRuntimeVersion string `json:"containerRuntimeVersion"`

	MemoryCapacityBytes    int64 `json:"memoryCapacityBytes"`
	MemoryAllocatableBytes int64 `json:"memoryAllocatableBytes"`
	CPUCapacityMillis      int64 `json:"cpuCapacityMillis"`
	CPUAllocatableMillis   int64 `json:"cpuAllocatableMillis"`
}

type NodeSpec struct {
	PodCIDRs      []string       `json:"podCIDRs,omitempty"`
	ProviderID    string         `json:"providerID,omitempty"`
	Unschedulable bool           `json:"unschedulable,omitempty"`
	Taints        []corev1.Taint `json:"taints,omitempty"`
}

type NodeStatus struct {
	NodeInfo corev1.NodeSystemInfo `json:"nodeInfo"`
}

func NewNode() *Node {
	return &Node{
		TypeMeta: TypeMeta{
			Kind:         "Node",
			APIGroup:     "core",
			APIVersion:   "v1",
			ResourceType: "nodes",
		},
		ObjectMeta: NewObjectMeta(metav1.ObjectMeta{}),
		Spec:       NodeSpec{},
		Status:     NodeStatus{},
	}
}
