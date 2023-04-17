package inventory

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type RabbitMQ struct {
	Clusters []*RabbitMQCluster `json:"clusters"`
}

func NewRabbitMQ() *RabbitMQ {
	return &RabbitMQ{
		Clusters: make([]*RabbitMQCluster, 0),
	}
}

type RabbitMQCluster struct {
	Name              string            `json:"name"`
	Namespace         string            `json:"namespace"`
	Labels            map[string]string `json:"labels"`
	Annotations       map[string]string `json:"annotations"`
	CreationTimestamp metav1.Time       `json:"creation_timestamp"`
	Image             string            `json:"image"`
}

func NewRabbitMQCluster() *RabbitMQCluster {
	return &RabbitMQCluster{
		Labels:      make(map[string]string, 0),
		Annotations: make(map[string]string, 0),
	}
}
