package inventory

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type Deployment struct {
	Name              string                `json:"name" db:"name"`
	Namespace         string                `json:"namespace" db:"namespace"`
	CreationTimestamp metav1.Time           `json:"creation_timestamp" db:"creation_timestamp"`
	Labels            KubernetesLabels      `json:"labels" db:"labels"`
	Annotations       KubernetesAnnotations `json:"annotations" db:"annotations"`
	Replicas          *int32                `json:"replicas" db:"replicas"`
	Strategy          string                `json:"strategy" db:"strategy"`
	Template          *PodTemplate          `json:"template"`
}

func NewDeployment() *Deployment {
	return &Deployment{
		Labels:      make(KubernetesLabels, 0),
		Annotations: make(KubernetesAnnotations, 0),
		Template:    NewPodTemplate(),
	}
}
