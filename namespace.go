package inventory

import (
	"database/sql/driver"
	"encoding/json"
	"errors"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

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

type Namespaces []*Namespace

type Namespace struct {
	Name              string                `json:"name" db:"name"`
	CreationTimestamp metav1.Time           `json:"creation_timestamp" db:"creation_timestamp"`
	Labels            KubernetesLabels      `json:"labels" db:"labels"`
	Annotations       KubernetesAnnotations `json:"annotations" db:"annotations"`
}

func NewNamespace() *Namespace {
	return &Namespace{
		Labels:      make(KubernetesLabels, 0),
		Annotations: make(KubernetesAnnotations, 0),
	}
}
