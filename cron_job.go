package inventory

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type CronJob struct {
	Name              string            `json:"name"`
	Namespace         string            `json:"namespace"`
	CreationTimestamp metav1.Time       `json:"creation_timestamp"`
	Labels            map[string]string `json:"labels"`
	Annotations       map[string]string `json:"annotations"`
	Schedule          string            `json:"schedule"`
	ConcurrencyPolicy string            `json:"concurrency_policy"`
	Template          *PodTemplate      `json:"template"`
}

func NewCronJob() *CronJob {
	return &CronJob{
		Labels:      make(map[string]string, 0),
		Annotations: make(map[string]string, 0),
		Template:    NewPodTemplate(),
	}
}
