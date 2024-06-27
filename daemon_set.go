package inventory

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type DaemonSetSpec struct {
	Template       *PodTemplate `json:"template"`
	UpdateStrategy string       `json:"updateStrategy"`
}

type DaemonSetStatus struct {
	CurrentNumberScheduled int32 `json:"currentNumberScheduled"`
	NumberMisscheduled     int32 `json:"numberMisscheduled"`
	DesiredNumberScheduled int32 `json:"desiredNumberScheduled"`
}

func NewDaemonSet() *Workload {
	return &Workload{
		PartialObject: PartialObject{
			TypeMeta: TypeMeta{
				Kind:         "DaemonSet",
				APIGroup:     "apps",
				APIVersion:   "v1",
				ResourceType: "daemonsets",
			},
			ObjectMeta: NewObjectMeta(metav1.ObjectMeta{}),
		},
		Spec: DaemonSetSpec{
			Template: NewPodTemplate(),
		},
		Status: DaemonSetStatus{},
	}
}
