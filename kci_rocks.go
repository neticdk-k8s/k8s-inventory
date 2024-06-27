package inventory

import metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

type KCIRocks struct {
	DBInstances KCIRocksDBInstances `json:"dbInstances"`
}

func NewKCIRocks() *KCIRocks {
	return &KCIRocks{
		DBInstances: make([]*KCIRocksDBInstance, 0),
	}
}

func (k *KCIRocks) AddDBInstance(i *KCIRocksDBInstance) {
	k.DBInstances = k.DBInstances.Add(i)
}

func (k *KCIRocks) DeleteDBInstance(i *KCIRocksDBInstance) {
	k.DBInstances = k.DBInstances.Delete(i)
}

type KCIRocksDBInstance struct {
	PartialObject

	Spec   KCIRocksDBInstanceSpec   `json:"spec"`
	Status KCIRocksDBInstanceStatus `json:"status"`
}

type KCIRocksDBInstances = Set[*KCIRocksDBInstance]

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
		PartialObject: PartialObject{
			TypeMeta: TypeMeta{
				Kind:         "DBInstance",
				APIGroup:     "kci.rocks",
				APIVersion:   "v1alpha1",
				ResourceType: "dbinstances",
			},
			ObjectMeta: NewObjectMeta(metav1.ObjectMeta{}),
		},
		Spec:   KCIRocksDBInstanceSpec{},
		Status: KCIRocksDBInstanceStatus{},
	}
}
