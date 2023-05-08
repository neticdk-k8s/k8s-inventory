package inventory

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type KCIRocks struct {
	DBInstances []*KCIRocksDBInstance `json:"db_instances"`
}

func NewKCIRocks() *KCIRocks {
	return &KCIRocks{
		DBInstances: make([]*KCIRocksDBInstance, 0),
	}
}

type KCIRocksDBInstance struct {
	Name              string                `json:"name" db:"name"`
	Labels            KubernetesLabels      `json:"labels" db:"labels"`
	Annotations       KubernetesAnnotations `json:"annotations" db:"annotations"`
	CreationTimestamp metav1.Time           `json:"creation_timestamp" db:"creation_timestamp"`
	Engine            string                `json:"engine" db:"engine"`
	Host              string                `json:"host" db:"host"`
	Port              uint16                `json:"port" db:"port"`
	Status            string                `json:"status" db:"status"`
}

func NewKCIRocksDBInstance() *KCIRocksDBInstance {
	return &KCIRocksDBInstance{
		Labels:      make(KubernetesLabels, 0),
		Annotations: make(KubernetesAnnotations, 0),
	}
}
