package inventory

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type StatefulSet struct {
	Name              string                `json:"name" db:"name"`
	Namespace         string                `json:"namespace" db:"namespace"`
	CreationTimestamp metav1.Time           `json:"creation_timestamp" db:"creation_timestamp"`
	Labels            KubernetesLabels      `json:"labels" db:"labels"`
	Annotations       KubernetesAnnotations `json:"annotations" db:"annotations"`
	Replicas          *int32                `json:"replicas" db:"replicas"`
	Template          *PodTemplate          `json:"template"`
	ServiceName       string                `json:"service_name" db:"service_name"`
	UpdateStrategy    string                `json:"update_strategy" db:"update_strategy"`
}

func NewStatefulSet() *StatefulSet {
	return &StatefulSet{
		Labels:      make(KubernetesLabels, 0),
		Annotations: make(KubernetesAnnotations, 0),
		Template:    NewPodTemplate(),
	}
}
