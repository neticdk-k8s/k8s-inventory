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
	IsControlPlane bool   `json:"is_control_plane"`
	TopologyRegion string `json:"topology_region"`
	TopologyZone   string `json:"topology_zone"`
	Provider       string `json:"provider"`

	KubeProxyVersion        string `json:"kube_proxy_version"`
	KubeletVersion          string `json:"kubelet_version"`
	KernelVersion           string `json:"kernel_version"`
	CRIName                 string `json:"cri_name"`
	CRIVersion              string `json:"cri_version"`
	ContainerRuntimeVersion string `json:"container_runtime_version"`

	MemoryCapacityBytes    int64 `json:"memory_capacity_bytes"`
	MemoryAllocatableBytes int64 `json:"memory_allocatable_bytes"`
	CPUCapacityMillis      int64 `json:"cpu_capacity_mi"`
	CPUAllocatableMillis   int64 `json:"cpu_allocatable_mi"`
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
