package inventory

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type Deployment struct {
	Name              string            `json:"name" db:"name"`
	Namespace         string            `json:"namespace" db:"namespace"`
	CreationTimestamp metav1.Time       `json:"creation_timestamp" db:"creation_timestamp"`
	Labels            map[string]string `json:"labels"`
	Annotations       map[string]string `json:"annotations"`
	Replicas          *int32            `json:"replicas" db:"replicas"`
	Strategy          string            `json:"strategy" db:"strategy"`
	Template          *PodTemplate      `json:"template"`
}

func NewDeployment() *Deployment {
	return &Deployment{
		Labels:      make(map[string]string, 0),
		Annotations: make(map[string]string, 0),
		Template:    NewPodTemplate(),
	}
}
