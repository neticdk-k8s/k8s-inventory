package inventory

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type ReplicaSet struct {
	Name              string            `json:"name"`
	Namespace         string            `json:"namespace"`
	CreationTimestamp metav1.Time       `json:"creation_timestamp"`
	Labels            map[string]string `json:"labels"`
	Annotations       map[string]string `json:"annotations"`
	Replicas          *int32            `json:"replicas"`
	Template          *PodTemplate      `json:"template"`
	OwnerName         string            `json:"owner_name"`
	OwnerKind         string            `json:"owner_kind"`
}

func NewReplicaSet() *ReplicaSet {
	return &ReplicaSet{
		Labels:      make(map[string]string, 0),
		Annotations: make(map[string]string, 0),
		Template:    NewPodTemplate(),
	}
}
