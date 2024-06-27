package inventory

import (
	"time"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type Inventory struct {
	Cluster             *Cluster         `json:"cluster"`
	Nodes               Nodes            `json:"nodes"`
	Namespaces          Namespaces       `json:"namespaces"`
	Workloads           Workloads        `json:"workloads"`
	Storage             *Storage         `json:"storage"`
	NetworkPolicies     NetworkPolicies  `json:"networkPolicies"`
	CustomResources     *CustomResources `json:"customResources"`
	CollectedAt         metav1.Time      `json:"collectedAt"`
	CollectionSucceeded bool             `json:"collectionSucceeded"`
	CollectionErrors    []string         `json:"collectionErrors"`
	ClientVersion       string           `json:"clientVersion"`
	ClientCommit        string           `json:"clientCommit"`
}

func NewInventory() *Inventory {
	return &Inventory{
		Cluster:          NewCluster(),
		Nodes:            Nodes{},
		Namespaces:       Namespaces{},
		NetworkPolicies:  NetworkPolicies{},
		Workloads:        Workloads{},
		Storage:          NewStorage(),
		CustomResources:  NewCustomResources(),
		CollectedAt:      metav1.Time{Time: time.Now()},
		CollectionErrors: make([]string, 0),
	}
}

func (i *Inventory) AddWorkload(w *Workload) {
	i.Workloads = i.Workloads.Add(w)
}

func (i *Inventory) DeleteWorkload(w *Workload) {
	i.Workloads = i.Workloads.Delete(w)
}

func (i *Inventory) AddNode(n *Node) {
	i.Nodes = i.Nodes.Add(n)
}

func (i *Inventory) DeleteNode(n *Node) {
	i.Nodes = i.Nodes.Delete(n)
}

func (i *Inventory) AddNamespace(n *Namespace) {
	i.Namespaces = i.Namespaces.Add(n)
}

func (i *Inventory) DeleteNamespace(n *Namespace) {
	i.Namespaces = i.Namespaces.Delete(n)
}

func (i *Inventory) AddNetworkPolicy(n *NetworkPolicy) {
	i.NetworkPolicies = i.NetworkPolicies.Add(n)
}

func (i *Inventory) DeleteNetworkPolicy(n *NetworkPolicy) {
	i.NetworkPolicies = i.NetworkPolicies.Delete(n)
}
