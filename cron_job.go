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

type CronJobStatus struct {
	LastScheduleTime   *metav1.Time `json:"lastScheduleTime,omitempty"`
	LastSuccessfulTime *metav1.Time `json:"lastSuccessfulTime,omitempty"`
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
