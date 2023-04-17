package inventory

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type StorageClass struct {
	Name              string            `json:"name"`
	CreationTimestamp metav1.Time       `json:"creation_timestamp"`
	Labels            map[string]string `json:"labels"`
	Annotations       map[string]string `json:"annotations"`
	Provisioner       string            `json:"provisioner"`
}

func NewStorageClass() *StorageClass {
	return &StorageClass{
		Labels:      make(map[string]string, 0),
		Annotations: make(map[string]string, 0),
	}
}
