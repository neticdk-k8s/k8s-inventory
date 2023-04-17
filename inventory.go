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
	Namespaces          Namespaces       `json:"namespacxs"`
	Tenants             Tenants          `json:"tenants"`
	Workloads           *Workload        `json:"workloads"`
	Storage             *Storage         `json:"storage"`
	CustomResources     *CustomResources `json:"custom_resources"`
	CollectedAt         metav1.Time      `json:"collected_at" validate:"datetime"`
	CollectionSucceeded bool             `json:"collection_succeeded"`
}

func NewInventory() *Inventory {
	return &Inventory{
		CollectedAt:     metav1.Time{Time: time.Now()},
		Cluster:         NewCluster(),
		CustomResources: NewCustomResources(),
		Namespaces:      Namespaces{},
		Nodes:           Nodes{},
		Workloads:       NewWorkload(),
		Storage:         NewStorage(),
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
