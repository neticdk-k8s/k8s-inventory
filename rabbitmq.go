package inventory

import metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

type RabbitMQ struct {
	Clusters []*RabbitMQCluster `json:"clusters"`
}

func NewRabbitMQ() *RabbitMQ {
	return &RabbitMQ{
		Clusters: make([]*RabbitMQCluster, 0),
	}
}

type RabbitMQCluster struct {
	TypeMeta
	ObjectMeta

	Spec   RabbitMQClusterSpec   `json:"spec" db:"spec"`
	Status RabbitMQClusterStatus `json:"status" db:"status"`
}

type RabbitMQClusterSpec struct {
	Replicas *int32 `json:"replicas"`
	Image    string `json:"image"`
}

type RabbitMQClusterStatus struct{}

func NewRabbitMQCluster() *RabbitMQCluster {
	return &RabbitMQCluster{
		TypeMeta: TypeMeta{
			Kind:         "RabbitmqCluster",
			APIGroup:     "rabbitmq.com",
			APIVersion:   "v1beta1",
			ResourceType: "rabbitmqclusters",
		},
		ObjectMeta: NewObjectMeta(metav1.ObjectMeta{}),
		Spec:       RabbitMQClusterSpec{},
		Status:     RabbitMQClusterStatus{},
	}
}
