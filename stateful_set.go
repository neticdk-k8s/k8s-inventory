package inventory

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type StatefulSetSpec struct {
	Replicas       *int32       `json:"replicas"`
	Template       *PodTemplate `json:"template"`
	ServiceName    string       `json:"serviceName"`
	UpdateStrategy string       `json:"updateStrategy"`
}

type StatefulSetStatus struct {
	Replicas          int32 `json:"replicas"`
	ReadyReplicas     int32 `json:"readyReplicas"`
	CurrentReplicas   int32 `json:"currentReplicas"`
	UpdatedReplicas   int32 `json:"updatedReplicas"`
	AvailableReplicas int32 `json:"availableReplicas"`
}

func NewStatefulSet() *Workload {
	return &Workload{
		TypeMeta: TypeMeta{
			Kind:         "StatefulSet",
			APIGroup:     "apps",
			APIVersion:   "v1",
			ResourceType: "statefulsets",
		},
		ObjectMeta: NewObjectMeta(metav1.ObjectMeta{}),
		Spec: StatefulSetSpec{
			Template: NewPodTemplate(),
		},
		Status: StatefulSetStatus{},
	}
}
