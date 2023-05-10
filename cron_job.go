package inventory

import metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

type CronJob struct {
	TypeMeta
	ObjectMeta

	Spec   CronJobSpec   `json:"spec" db:"spec"`
	Status CronJobStatus `json:"status" db:"status"`
}

type CronJobSpec struct {
	Schedule          string       `json:"schedule"`
	ConcurrencyPolicy string       `json:"concurrency_policy"`
	JobTemplate       *PodTemplate `json:"job_template"`
}

type CronJobStatus struct {
	LastScheduleTime   *metav1.Time `json:"last_schedule_time,omitempty"`
	LastSuccessfulTime *metav1.Time `json:"last_successful_time,omitempty"`
}

func NewCronJob() *CronJob {
	return &CronJob{
		TypeMeta: TypeMeta{
			Kind:         "CronJob",
			APIGroup:     "batch",
			APIVersion:   "v1",
			ResourceType: "cronjobs",
		},
		ObjectMeta: NewObjectMeta(),
		Spec: CronJobSpec{
			JobTemplate: NewPodTemplate(),
		},
		Status: CronJobStatus{},
	}
}
