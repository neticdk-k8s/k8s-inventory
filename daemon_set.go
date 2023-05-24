package inventory

import (
	"database/sql/driver"
	"encoding/json"
	"errors"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type DaemonSetSpec struct {
	Template       *PodTemplate `json:"template"`
	UpdateStrategy string       `json:"update_strategy"`
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
	CurrentNumberScheduled int32 `json:"current_number_scheduled"`
	NumberMisscheduled     int32 `json:"number_misscheduled"`
	DesiredNumberScheduled int32 `json:"desired_number_scheduled"`
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
