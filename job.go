package inventory

import (
	"database/sql/driver"
	"encoding/json"
	"errors"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type JobSpec struct {
	Parallelism  *int32       `json:"parallelism"`
	Completions  *int32       `json:"completions"`
	BackoffLimit *int32       `json:"backoffLimit"`
	Template     *PodTemplate `json:"template"`
}

func (js *JobSpec) Value() (driver.Value, error) {
	bytes, err := json.Marshal(js)
	return bytes, err
}

func (js *JobSpec) Scan(val interface{}) error {
	b, ok := val.([]byte)
	if !ok {
		return errors.New("type assertion to []byte failed")
	}

	return json.Unmarshal(b, &js)
}

type JobStatus struct {
	StartTime      *metav1.Time `json:"startTime"`
	CompletionTime *metav1.Time `json:"completionTime,omitempty"`
	Active         int32        `json:"active"`
	Ready          *int32       `json:"ready"`
	Succeeded      int32        `json:"succeeded"`
	Failed         int32        `json:"failed"`
}

func (js *JobStatus) Value() (driver.Value, error) {
	bytes, err := json.Marshal(js)
	return bytes, err
}

func (js *JobStatus) Scan(val interface{}) error {
	b, ok := val.([]byte)
	if !ok {
		return errors.New("type assertion to []byte failed")
	}

	return json.Unmarshal(b, &js)
}

func NewJob() *Workload {
	return &Workload{
		TypeMeta: TypeMeta{
			Kind:         "Job",
			APIGroup:     "batch",
			APIVersion:   "v1",
			ResourceType: "jobs",
		},
		ObjectMeta: NewObjectMeta(metav1.ObjectMeta{}),
		Spec: JobSpec{
			Template: NewPodTemplate(),
		},
		Status: JobStatus{},
	}
}
