package inventory

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type Namespaces []*Namespace

type Namespace struct {
	Name              string            `json:"name" db:"name"`
	CreationTimestamp metav1.Time       `json:"creation_timestamp" db:"creation_timestamp"`
	Labels            map[string]string `json:"labels"`
	Annotations       map[string]string `json:"annotations"`
}

func NewNamespace() *Namespace {
	return &Namespace{
		Labels:      make(map[string]string, 0),
		Annotations: make(map[string]string, 0),
	}
}
