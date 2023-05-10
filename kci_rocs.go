package inventory

type KCIRocks struct {
	DBInstances []*KCIRocksDBInstance `json:"db_instances"`
}

func NewKCIRocks() *KCIRocks {
	return &KCIRocks{
		DBInstances: make([]*KCIRocksDBInstance, 0),
	}
}

type KCIRocksDBInstance struct {
	TypeMeta
	ObjectMeta

	Spec   KCIRocksDBInstanceSpec   `json:"spec" db:"spec"`
	Status KCIRocksDBInstanceStatus `json:"status" db:"status"`
}

type KCIRocksDBInstanceSpec struct {
	Engine string `json:"engine"`
	Host   string `json:"host"`
	Port   uint16 `json:"port"`
}

type KCIRocksDBInstanceStatus struct {
	Phase string `json:"status"`
}

func NewKCIRocksDBInstance() *KCIRocksDBInstance {
	return &KCIRocksDBInstance{
		TypeMeta: TypeMeta{
			Kind:         "DBInstance",
			APIGroup:     "kci.rocks",
			APIVersion:   "v1alpha1",
			ResourceType: "dbinstances",
		},
		ObjectMeta: NewObjectMeta(),
		Spec:       KCIRocksDBInstanceSpec{},
		Status:     KCIRocksDBInstanceStatus{},
	}
}
