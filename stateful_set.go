package inventory

import (
	"database/sql/driver"
	"encoding/json"
	"errors"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type StatefulSetSpec struct {
	Replicas       *int32       `json:"replicas"`
	Template       *PodTemplate `json:"template"`
	ServiceName    string       `json:"serviceName"`
	UpdateStrategy string       `json:"updateStrategy"`
}

func (ss *StatefulSetSpec) Value() (driver.Value, error) {
	bytes, err := json.Marshal(ss)
	return bytes, err
}

func (ss *StatefulSetSpec) Scan(val interface{}) error {
	b, ok := val.([]byte)
	if !ok {
		return errors.New("type assertion to []byte failed")
	}

	return json.Unmarshal(b, &ss)
}

type StatefulSetStatus struct {
	Replicas          int32 `json:"replicas"`
	ReadyReplicas     int32 `json:"readyReplicas"`
	CurrentReplicas   int32 `json:"currentReplicas"`
	UpdatedReplicas   int32 `json:"updatedReplicas"`
	AvailableReplicas int32 `json:"availableReplicas"`
}

func (ss *StatefulSetStatus) Value() (driver.Value, error) {
	bytes, err := json.Marshal(ss)
	return bytes, err
}

func (ss *StatefulSetStatus) Scan(val interface{}) error {
	b, ok := val.([]byte)
	if !ok {
		return errors.New("type assertion to []byte failed")
	}

	return json.Unmarshal(b, &ss)
}

func NewStatefulSet() *Workload {
	return &Workload{
		TypeMeta: TypeMeta{
			Kind:         "StatefulSet",
			APIGroup:     "apps",
			APIVersion:   "v1",
			ResourceType: "statefulsets",
		},
		ObjectMeta: NewObjectMeta(metav1.ObjectMeta{}),
		Spec: StatefulSetSpec{
			Template: NewPodTemplate(),
		},
		Status: StatefulSetStatus{},
	}
}
