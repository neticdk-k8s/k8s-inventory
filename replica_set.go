package inventory

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type ReplicaSet struct {
	Name              string            `json:"name" db:"name"`
	Namespace         string            `json:"namespace" db:"namespace"`
	CreationTimestamp metav1.Time       `json:"creation_timestamp" db:"creation_timestamp"`
	Labels            map[string]string `json:"labels"`
	Annotations       map[string]string `json:"annotations"`
	Replicas          *int32            `json:"replicas" db:"replicas"`
	Template          *PodTemplate      `json:"template"`
	OwnerName         string            `json:"owner_name" db:"owner_name"`
	OwnerKind         string            `json:"owner_kind" db:"owner_kind"`
}

func NewReplicaSet() *ReplicaSet {
	return &ReplicaSet{
		Labels:      make(map[string]string, 0),
		Annotations: make(map[string]string, 0),
		Template:    NewPodTemplate(),
	}
}
