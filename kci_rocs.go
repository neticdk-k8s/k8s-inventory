package inventory

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type KCIRocks struct {
	DBInstances []*DBInstance `json:"db_instances"`
}

func NewKCIRocks() *KCIRocks {
	return &KCIRocks{
		DBInstances: make([]*DBInstance, 0),
	}
}

type DBInstance struct {
	Name              string            `json:"name"`
	Labels            map[string]string `json:"labels"`
	Annotations       map[string]string `json:"annotations"`
	CreationTimestamp metav1.Time       `json:"creation_timestamp"`
	Engine            string            `json:"engine"`
	Host              string            `json:"host"`
	Port              uint16            `json:"port"`
	Status            string            `json:"status"`
}

func NewDBInstance() *DBInstance {
	return &DBInstance{
		Labels:      make(map[string]string, 0),
		Annotations: make(map[string]string, 0),
	}
}
