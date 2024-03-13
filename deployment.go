package inventory

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type DeploymentSpec struct {
	Replicas *int32       `json:"replicas"`
	Strategy string       `json:"strategy"`
	Template *PodTemplate `json:"template"`
}

type DeploymentStatus struct {
	Replicas            int32 `json:"replicas"`
	ReadyReplicas       int32 `json:"readyReplicas"`
	UpdatedReplicas     int32 `json:"updatedReplicas"`
	AvailableReplicas   int32 `json:"availableReplicas"`
	UnavailableReplicas int32 `json:"unavailableReplicas"`
}

func NewDeployment() *Workload {
	return &Workload{
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
