package inventory

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type Calico struct {
	Labels            map[string]string `json:"labels"`
	CreationTimestamp metav1.Time       `json:"creation_timestamp"`
	Version           string            `json:"version"`
}

func NewCalico() *Calico {
	return &Calico{
		Labels: make(map[string]string, 0),
	}
}
