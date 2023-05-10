package inventory

import metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

type Job struct {
	TypeMeta
	ObjectMeta

	Spec   JobSpec   `json:"spec" db:"spec"`
	Status JobStatus `json:"status" db:"status"`
}

type JobSpec struct {
	Parallelism  *int32       `json:"parallelism"`
	Completions  *int32       `json:"completions"`
	BackoffLimit *int32       `json:"backoff_limit"`
	Template     *PodTemplate `json:"template"`
}

type JobStatus struct {
	StartTime      *metav1.Time `json:"start_time"`
	CompletionTime *metav1.Time `json:"completion_time"`
	Active         int32        `json:"active"`
	Ready          *int32       `json:"ready"`
	Succeeded      int32        `json:"succeeded"`
	Failed         int32        `json:"failed"`
}

func NewJob() *Job {
	return &Job{
		TypeMeta: TypeMeta{
			Kind:         "Job",
			APIGroup:     "batch",
			APIVersion:   "v1",
			ResourceType: "jobs",
		},
		ObjectMeta: NewObjectMeta(),
		Spec: JobSpec{
			Template: NewPodTemplate(),
		},
		Status: JobStatus{},
	}
}
