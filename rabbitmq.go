package inventory

import metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

type RabbitMQ struct {
	Clusters RabbitMQClusters `json:"clusters"`
}

func NewRabbitMQ() *RabbitMQ {
	return &RabbitMQ{
		Clusters: RabbitMQClusters{},
	}
}

func (r *RabbitMQ) AddCluster(c *RabbitMQCluster) {
	r.Clusters = r.Clusters.Add(c)
}

func (r *RabbitMQ) DeleteCluster(c *RabbitMQCluster) {
	r.Clusters = r.Clusters.Delete(c)
}

type RabbitMQCluster struct {
	PartialObject

	Spec   RabbitMQClusterSpec   `json:"spec"`
	Status RabbitMQClusterStatus `json:"status"`
}

type RabbitMQClusters = Set[*RabbitMQCluster]

type RabbitMQClusterSpec struct {
	Replicas *int32 `json:"replicas"`
	Image    string `json:"image"`
}

type RabbitMQClusterStatus struct{}

func NewRabbitMQCluster() *RabbitMQCluster {
	return &RabbitMQCluster{
		PartialObject: PartialObject{
			TypeMeta: TypeMeta{
				Kind:         "RabbitmqCluster",
				APIGroup:     "rabbitmq.com",
				APIVersion:   "v1beta1",
				ResourceType: "rabbitmqclusters",
			},
			ObjectMeta: NewObjectMeta(metav1.ObjectMeta{}),
		},
		Spec:   RabbitMQClusterSpec{},
		Status: RabbitMQClusterStatus{},
	}
}
