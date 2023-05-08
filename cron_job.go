package inventory

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type CronJob struct {
	Name              string                `json:"name" db:"name"`
	Namespace         string                `json:"namespace" db:"namespace"`
	CreationTimestamp metav1.Time           `json:"creation_timestamp" db:"creation_timestamp"`
	Labels            KubernetesLabels      `json:"labels" db:"labels"`
	Annotations       KubernetesAnnotations `json:"annotations" db:"annotations"`
	Schedule          string                `json:"schedule" db:"schedule"`
	ConcurrencyPolicy string                `json:"concurrency_policy" db:"concurrency_policy"`
	Template          *PodTemplate          `json:"template"`
}

func NewCronJob() *CronJob {
	return &CronJob{
		Labels:      make(KubernetesLabels, 0),
		Annotations: make(KubernetesAnnotations, 0),
		Template:    NewPodTemplate(),
	}
}
