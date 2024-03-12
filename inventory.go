package inventory

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"time"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type Inventory struct {
	Cluster             *Cluster         `json:"cluster"`
	Nodes               Nodes            `json:"nodes"`
	Namespaces          Namespaces       `json:"namespaces"`
	Workloads           []*Workload      `json:"workloads"`
	Pods                []*Pod           `json:"pods"`
	Storage             *Storage         `json:"storage"`
	NetworkPolicies     NetworkPolicies  `json:"network_policies"`
	CustomResources     *CustomResources `json:"custom_resources"`
	CollectedAt         metav1.Time      `json:"collected_at"`
	CollectionSucceeded bool             `json:"collection_succeeded"`
	CollectionErrors    []string         `json:"collection_errors"`
	ClientVersion       string           `json:"client_version"`
	ClientCommit        string           `json:"client_commit"`
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

func (i *Inventory) Value() (driver.Value, error) {
	return json.Marshal(i)
}

func (i *Inventory) Scan(value interface{}) error {
	b, ok := value.([]byte)
	if !ok {
		return errors.New("type assertion to []byte failed")
	}

	return json.Unmarshal(b, &i)
}
