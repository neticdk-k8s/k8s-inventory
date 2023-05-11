package inventory

import metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

type ReplicaSetSpec struct {
	Replicas *int32       `json:"replicas"`
	Template *PodTemplate `json:"template"`
}

type ReplicaSetStatus struct {
	Replicas             int32 `json:"replicas"`
	FullyLabeledReplicas int32 `json:"fully_labeled_replicas"`
	ReadyReplicas        int32 `json:"ready_replicas"`
	AvailableReplicas    int32 `json:"available_replicas"`
}

func NewReplicaSet() *Workload {
	return &Workload{
		TypeMeta: TypeMeta{
			Kind:         "ReplicaSet",
			APIGroup:     "apps",
			APIVersion:   "v1",
			ResourceType: "replicasets",
		},
		ObjectMeta: NewObjectMeta(metav1.ObjectMeta{}),
		Spec: ReplicaSetSpec{
			Template: NewPodTemplate(),
		},
		Status: ReplicaSetStatus{},
	}
}
