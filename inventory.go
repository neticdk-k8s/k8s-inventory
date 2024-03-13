package inventory

import (
	"time"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type Inventory struct {
	Cluster             *Cluster         `json:"cluster"`
	Nodes               Nodes            `json:"nodes"`
	Namespaces          Namespaces       `json:"namespaces"`
	Workloads           []*Workload      `json:"workloads"`
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
		Workloads:        NewWorkloads(),
		Storage:          NewStorage(),
		CustomResources:  NewCustomResources(),
		CollectedAt:      metav1.Time{Time: time.Now()},
		CollectionErrors: make([]string, 0),
	}
}
