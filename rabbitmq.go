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

	Image string `json:"image" db:"image"`
}

func NewRabbitMQCluster() *RabbitMQCluster {
	return &RabbitMQCluster{
		ObjectMeta: NewObjectMeta(),
	}
}
