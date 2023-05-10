package inventory

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
		ObjectMeta: NewObjectMeta(),
		Spec:       RabbitMQClusterSpec{},
		Status:     RabbitMQClusterStatus{},
	}
}
