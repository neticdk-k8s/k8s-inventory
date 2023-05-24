package inventory

import (
	"database/sql/driver"
	"encoding/json"
	"errors"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type DeploymentSpec struct {
	Replicas *int32       `json:"replicas"`
	Strategy string       `json:"strategy"`
	Template *PodTemplate `json:"template"`
}

func (ds *DeploymentSpec) Value() (driver.Value, error) {
	bytes, err := json.Marshal(ds)
	return bytes, err
}

func (ds *DeploymentSpec) Scan(val interface{}) error {
	b, ok := val.([]byte)
	if !ok {
		return errors.New("type assertion to []byte failed")
	}

	return json.Unmarshal(b, &ds)
}

type DeploymentStatus struct {
	Replicas            int32 `json:"replicas"`
	ReadyReplicas       int32 `json:"ready_replicas"`
	UpdatedReplicas     int32 `json:"updated_replicas"`
	AvailableReplicas   int32 `json:"available_replicas"`
	UnavailableReplicas int32 `json:"unavailable_replicas"`
}

func (ds *DeploymentStatus) Value() (driver.Value, error) {
	bytes, err := json.Marshal(ds)
	return bytes, err
}

func (ds *DeploymentStatus) Scan(val interface{}) error {
	b, ok := val.([]byte)
	if !ok {
		return errors.New("type assertion to []byte failed")
	}

	return json.Unmarshal(b, &ds)
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
