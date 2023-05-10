package inventory

type StatefulSet struct {
	TypeMeta
	ObjectMeta

	Spec   StatefulSetSpec   `json:"spec" db:"spec"`
	Status StatefulSetStatus `json:"status" db:"status"`
}

type StatefulSetSpec struct {
	Replicas       *int32       `json:"replicas"`
	Template       *PodTemplate `json:"template"`
	ServiceName    string       `json:"service_name"`
	UpdateStrategy string       `json:"update_strategy"`
}

type StatefulSetStatus struct {
	Replicas          int32 `json:"replicas"`
	ReadyReplicas     int32 `json:"ready_replicas"`
	CurrentReplicas   int32 `json:"current_replicas"`
	UpdatedReplicas   int32 `json:"updated_replicas"`
	AvailableReplicas int32 `json:"available_replicas"`
}

func NewStatefulSet() *StatefulSet {
	return &StatefulSet{
		ObjectMeta: NewObjectMeta(),
		Spec: StatefulSetSpec{
			Template: NewPodTemplate(),
		},
	}
}
