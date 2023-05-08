package inventory

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
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
