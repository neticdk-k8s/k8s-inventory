package inventory

import (
	"database/sql/driver"
	"encoding/json"
	"errors"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type ReplicaSetSpec struct {
	Replicas *int32       `json:"replicas"`
	Template *PodTemplate `json:"template"`
}

func (rs *ReplicaSetSpec) Value() (driver.Value, error) {
	bytes, err := json.Marshal(rs)
	return bytes, err
}

func (rs *ReplicaSetSpec) Scan(val interface{}) error {
	b, ok := val.([]byte)
	if !ok {
		return errors.New("type assertion to []byte failed")
	}

	return json.Unmarshal(b, &rs)
}

type ReplicaSetStatus struct {
	Replicas             int32 `json:"replicas"`
	FullyLabeledReplicas int32 `json:"fullyLabeledReplicas"`
	ReadyReplicas        int32 `json:"readyReplicas"`
	AvailableReplicas    int32 `json:"availableReplicas"`
}

func (rs *ReplicaSetStatus) Value() (driver.Value, error) {
	bytes, err := json.Marshal(rs)
	return bytes, err
}

func (rs *ReplicaSetStatus) Scan(val interface{}) error {
	b, ok := val.([]byte)
	if !ok {
		return errors.New("type assertion to []byte failed")
	}

	return json.Unmarshal(b, &rs)
}

func NewReplicaSet() *Workload {
	return &Workload{
		TypeMeta: TypeMeta{
			Kind:         "ReplicaSet",
			APIGroup:     "apps",
			APIVersion:   "v1",
			ResourceType: "replicasets",
		},
		ObjectMeta: NewObjectMeta(metav1.ObjectMeta{}),
		Spec: ReplicaSetSpec{
			Template: NewPodTemplate(),
		},
		Status: ReplicaSetStatus{},
	}
}
