package inventory

type Nodes []*Node

type Node struct {
	Name                    string `json:"name" db:"name"`
	Role                    string `json:"role" db:"role"`
	KubeProxyVersion        string `json:"kube_proxy_version" db:"kube_proxy_version"`
	KubeletVersion          string `json:"kubelet_version" db:"kubelet_version"`
	KernelVersion           string `json:"kernel_version" db:"kernel_version"`
	CRIName                 string `json:"cri_name" db:"cri_name"`
	CRIVersion              string `json:"cri_version" db:"cri_version"`
	ContainerRuntimeVersion string `json:"container_runtime_version" db:"container_runtime_version"`
	IsControlPlane          bool   `json:"is_control_plane" db:"is_control_plane"`
	Provider                string `json:"provider" db:"provider"`
	TopologyRegion          string `json:"topology_region" db:"topology_region"`
	TopologyZone            string `json:"topology_zone" db:"topology_zone"`
	MemoryCapacityBytes     int64  `json:"memory_capacity_bytes" db:"memory_capacity_bytes"`
	MemoryAllocatableBytes  int64  `json:"memory_allocatable_bytes" db:"memory_allocatable_bytes"`
	CPUCapacityMillis       int64  `json:"cpu_capacity_mi" db:"cpu_capacity_millis"`
	CPUAllocatableMillis    int64  `json:"cpu_allocatable_mi" db:"cpu_allocatable_millis"`
}

func NewNode() *Node {
	return &Node{}
}
