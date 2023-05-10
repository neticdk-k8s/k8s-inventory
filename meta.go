package inventory

import (
	"database/sql/driver"
	"encoding/json"
	"errors"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type ObjectMeta struct {
	Name              string                `json:"name" db:"name"`
	Namespace         string                `json:"namespace" db:"namespace"`
	CreationTimestamp metav1.Time           `json:"creation_timestamp" db:"creation_timestamp"`
	Labels            KubernetesLabels      `json:"labels" db:"labels"`
	Annotations       KubernetesAnnotations `json:"annotations" db:"annotations"`
	OwnerReferences   []OwnerReference      `json:"owner_references" db:"owner_references"`
}

func NewObjectMeta() ObjectMeta {
	return ObjectMeta{
		Labels:          make(KubernetesLabels, 0),
		Annotations:     make(KubernetesAnnotations, 0),
		OwnerReferences: make([]OwnerReference, 0),
	}
}

type TypeMeta struct {
	Kind         string `json:"kind" db:"kind"`
	APIGroup     string `json:"api_group" db:"api_group"`
	ResourceType string `json:"resource_type" db:"resource_type"`
	APIVersion   string `json:"api_version" db:"api_version"`
}

type OwnerReference struct {
	Kind       string `json:"kind" db:"kind"`
	APIVersion string `json:"api_version" db:"api_version"`
	Name       string `json:"name" db:"name"`
	Controller *bool  `json:"controller" db:"controller"`
}

type KubernetesLabels map[string]string

func (kl *KubernetesLabels) Scan(val interface{}) error {
	b, ok := val.([]byte)
	if !ok {
		return errors.New("type assertion to []byte failed")
	}

	return json.Unmarshal(b, &kl)
}

func (kl *KubernetesLabels) Value() (driver.Value, error) {
	bytes, err := json.Marshal(kl)
	return bytes, err
}

type KubernetesAnnotations map[string]string

func (ka *KubernetesAnnotations) Value() (driver.Value, error) {
	bytes, err := json.Marshal(ka)
	return bytes, err
}

func (ka *KubernetesAnnotations) Scan(val interface{}) error {
	b, ok := val.([]byte)
	if !ok {
		return errors.New("type assertion to []byte failed")
	}

	return json.Unmarshal(b, &ka)
}
