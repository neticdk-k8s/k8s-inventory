package inventory

import (
	"database/sql/driver"
	"encoding/json"
	"errors"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type CronJobSpec struct {
	Schedule          string       `json:"schedule"`
	ConcurrencyPolicy string       `json:"concurrencyPolicy"`
	JobTemplate       *PodTemplate `json:"jobTemplate"`
}

func (cs *CronJobSpec) Value() (driver.Value, error) {
	bytes, err := json.Marshal(cs)
	return bytes, err
}

func (cs *CronJobSpec) Scan(val interface{}) error {
	b, ok := val.([]byte)
	if !ok {
		return errors.New("type assertion to []byte failed")
	}

	return json.Unmarshal(b, &cs)
}

type CronJobStatus struct {
	LastScheduleTime   *metav1.Time `json:"lastScheduleTime,omitempty"`
	LastSuccessfulTime *metav1.Time `json:"lastSuccessfulTime,omitempty"`
}

func (cs *CronJobStatus) Value() (driver.Value, error) {
	bytes, err := json.Marshal(cs)
	return bytes, err
}

func (cs *CronJobStatus) Scan(val interface{}) error {
	b, ok := val.([]byte)
	if !ok {
		return errors.New("type assertion to []byte failed")
	}

	return json.Unmarshal(b, &cs)
}

func NewCronJob() *Workload {
	return &Workload{
		TypeMeta: TypeMeta{
			Kind:         "CronJob",
			APIGroup:     "batch",
			APIVersion:   "v1",
			ResourceType: "cronjobs",
		},
		ObjectMeta: NewObjectMeta(metav1.ObjectMeta{}),
		Spec: CronJobSpec{
			JobTemplate: NewPodTemplate(),
		},
		Status: CronJobStatus{},
	}
}
