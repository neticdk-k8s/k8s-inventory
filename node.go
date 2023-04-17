package inventory

type Nodes []*Node

type Node struct {
	Name                    string `json:"name"`
	Role                    string `json:"role"`
	KubeProxyVersion        string `json:"kube_proxy_version"`
	KubeletVersion          string `json:"kubelet_version"`
	KernelVersion           string `json:"kernel_version"`
	CRIName                 string `json:"cri_name"`
	CRIVersion              string `json:"cri_version"`
	ContainerRuntimeVersion string `json:"container_runtime_version"`
	IsControlPlane          bool   `json:"is_control_plane"`
	Provider                string `json:"provider"`
	TopologyRegion          string `json:"topology_region"`
	TopologyZone            string `json:"topology_zone"`
	MemoryCapacityBytes     int64  `json:"memory_capacity_bytes"`
	MemoryAllocatableBytes  int64  `json:"memory_allocatable_bytes"`
	CPUCapacityMillis       int64  `json:"cpu_capacity_mi"`
	CPUAllocatableMillis    int64  `json:"cpu_allocatable_mi"`
}

func NewNode() *Node {
	return &Node{}
}
