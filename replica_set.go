package inventory

type ReplicaSet struct {
	TypeMeta
	ObjectMeta

	Spec   ReplicaSetSpec   `json:"spec" db:"spec"`
	Status ReplicaSetStatus `json:"status" db:"status"`
}

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

func NewReplicaSet() *ReplicaSet {
	return &ReplicaSet{
		TypeMeta: TypeMeta{
			Kind:         "ReplicaSet",
			APIGroup:     "apps",
			APIVersion:   "v1",
			ResourceType: "replicasets",
		},
		ObjectMeta: NewObjectMeta(),
		Spec: ReplicaSetSpec{
			Template: NewPodTemplate(),
		},
		Status: ReplicaSetStatus{},
	}
}
