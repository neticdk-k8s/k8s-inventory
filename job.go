package inventory

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type JobSpec struct {
	Parallelism  *int32       `json:"parallelism"`
	Completions  *int32       `json:"completions"`
	BackoffLimit *int32       `json:"backoffLimit"`
	Template     *PodTemplate `json:"template"`
}

type JobStatus struct {
	StartTime      *metav1.Time `json:"startTime"`
	CompletionTime *metav1.Time `json:"completionTime,omitempty"`
	Active         int32        `json:"active"`
	Ready          *int32       `json:"ready"`
	Succeeded      int32        `json:"succeeded"`
	Failed         int32        `json:"failed"`
}

func NewJob() *Workload {
	return &Workload{
		PartialObject: PartialObject{
			TypeMeta: TypeMeta{
				Kind:         "Job",
				APIGroup:     "batch",
				APIVersion:   "v1",
				ResourceType: "jobs",
			},
			ObjectMeta: NewObjectMeta(metav1.ObjectMeta{}),
		},
		Spec: JobSpec{
			Template: NewPodTemplate(),
		},
		Status: JobStatus{},
	}
}
