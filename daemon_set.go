package inventory

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type DaemonSet struct {
	Name              string            `json:"name" db:"name"`
	Namespace         string            `json:"namespace" db:"namespace"`
	CreationTimestamp metav1.Time       `json:"creation_timestamp" db:"creation_timestamp"`
	Labels            map[string]string `json:"labels" db:"labels"`
	Annotations       map[string]string `json:"annotations" db:"annotations"`
	Template          *PodTemplate      `json:"template"`
	UpdateStrategy    string            `json:"update_strategy" db:"update_strategy"`
}

func NewDaemonSet() *DaemonSet {
	return &DaemonSet{
		Labels:      make(map[string]string, 0),
		Annotations: make(map[string]string, 0),
		Template:    NewPodTemplate(),
	}
}
