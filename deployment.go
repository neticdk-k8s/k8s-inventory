package inventory

import metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

type Deployment struct {
	TypeMeta
	ObjectMeta

	Spec   DeploymentSpec   `json:"spec" db:"spec"`
	Status DeploymentStatus `json:"status" db:"status"`
}

type DeploymentSpec struct {
	Replicas *int32       `json:"replicas"`
	Strategy string       `json:"strategy"`
	Template *PodTemplate `json:"template"`
}

type DeploymentStatus struct {
	Replicas            int32 `json:"replicas"`
	ReadyReplicas       int32 `json:"ready_replicas"`
	UpdatedReplicas     int32 `json:"updated_replicas"`
	AvailableReplicas   int32 `json:"available_replicas"`
	UnavailableReplicas int32 `json:"unavailable_replicas"`
}

func NewDeployment() *Deployment {
	return &Deployment{
		TypeMeta: TypeMeta{
			Kind:         "Deployment",
			APIGroup:     "apps",
			APIVersion:   "v1",
			ResourceType: "deployments",
		},
		ObjectMeta: NewObjectMeta(metav1.ObjectMeta{}),
		Spec: DeploymentSpec{
			Template: NewPodTemplate(),
		},
		Status: DeploymentStatus{},
	}
}
