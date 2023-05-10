package inventory

type Deployment struct {
	TypeMeta
	ObjectMeta

	Spec   DeploymenSpec    `json:"spec" db:"spec"`
	Status DeploymentStatus `json:"status" db:"status"`
}

type DeploymenSpec struct {
	Replicas *int32       `json:"replicas"`
	Strategy string       `json:"strategy"`
	Template *PodTemplate `json:"template"`
}

func NewDeploymentSpec() DeploymenSpec {
	return DeploymenSpec{
		Template: NewPodTemplate(),
	}
}

type DeploymentStatus struct {
	Replicas            *int32 `json:"replicas"`
	ReadyReplicas       *int32 `json:"ready_replicas"`
	UpdatedReplicas     *int32 `json:"updated_replicas"`
	AvailableReplicas   *int32 `json:"available_replicas"`
	UnavailableReplicas *int32 `json:"unavailable_replicas"`
}

func NewDeployment() *Deployment {
	return &Deployment{
		ObjectMeta: NewObjectMeta(),
		Spec:       NewDeploymentSpec(),
		Status:     DeploymentStatus{},
	}
}
