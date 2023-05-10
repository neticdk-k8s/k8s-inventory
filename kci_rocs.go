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

	Engine string `json:"engine" db:"engine"`
	Host   string `json:"host" db:"host"`
	Port   uint16 `json:"port" db:"port"`
	Status string `json:"status" db:"status"`
}

func NewKCIRocksDBInstance() *KCIRocksDBInstance {
	return &KCIRocksDBInstance{
		ObjectMeta: NewObjectMeta(),
	}
}
