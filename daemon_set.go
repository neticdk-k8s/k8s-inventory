package inventory

import (
	"database/sql/driver"
	"encoding/json"
	"errors"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type DaemonSetSpec struct {
	Template       *PodTemplate `json:"template"`
	UpdateStrategy string       `json:"updateStrategy"`
}

func (ds *DaemonSetSpec) Value() (driver.Value, error) {
	bytes, err := json.Marshal(ds)
	return bytes, err
}

func (ds *DaemonSetSpec) Scan(val interface{}) error {
	b, ok := val.([]byte)
	if !ok {
		return errors.New("type assertion to []byte failed")
	}

	return json.Unmarshal(b, &ds)
}

type DaemonSetStatus struct {
	CurrentNumberScheduled int32 `json:"currentNumberScheduled"`
	NumberMisscheduled     int32 `json:"numberMisscheduled"`
	DesiredNumberScheduled int32 `json:"desiredNumberScheduled"`
}

func (ds *DaemonSetStatus) Value() (driver.Value, error) {
	bytes, err := json.Marshal(ds)
	return bytes, err
}

func (ds *DaemonSetStatus) Scan(val interface{}) error {
	b, ok := val.([]byte)
	if !ok {
		return errors.New("type assertion to []byte failed")
	}

	return json.Unmarshal(b, &ds)
}

func NewDaemonSet() *Workload {
	return &Workload{
		TypeMeta: TypeMeta{
			Kind:         "DaemonSet",
			APIGroup:     "apps",
			APIVersion:   "v1",
			ResourceType: "daemonsets",
		},
		ObjectMeta: NewObjectMeta(metav1.ObjectMeta{}),
		Spec: DaemonSetSpec{
			Template: NewPodTemplate(),
		},
		Status: DaemonSetStatus{},
	}
}
