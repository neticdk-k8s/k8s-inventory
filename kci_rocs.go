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
	Name              string            `json:"name"`
	Labels            map[string]string `json:"labels"`
	Annotations       map[string]string `json:"annotations"`
	CreationTimestamp metav1.Time       `json:"creation_timestamp"`
	Engine            string            `json:"engine"`
	Host              string            `json:"host"`
	Port              uint16            `json:"port"`
	Status            string            `json:"status"`
}

func NewKCIRocksDBInstance() *KCIRocksDBInstance {
	return &KCIRocksDBInstance{
		Labels:      make(map[string]string, 0),
		Annotations: make(map[string]string, 0),
	}
}
