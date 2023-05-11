package inventory

import metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

type DaemonSetSpec struct {
	Template       *PodTemplate `json:"template"`
	UpdateStrategy string       `json:"update_strategy"`
}

type DaemonSetStatus struct {
	CurrentNumberScheduled int32 `json:"current_number_scheduled"`
	NumberMisscheduled     int32 `json:"number_misscheduled"`
	DesiredNumberScheduled int32 `json:"desired_number_scheduled"`
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
