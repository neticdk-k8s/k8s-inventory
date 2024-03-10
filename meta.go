package inventory

import (
	"database/sql/driver"
	"encoding/json"
	"errors"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type ObjectMeta struct {
	Name              string                `json:"name" db:"name"`
	Namespace         string                `json:"namespace,omitempty" db:"namespace"`
	CreationTimestamp metav1.Time           `json:"creation_timestamp" db:"creation_timestamp"`
	Labels            KubernetesLabels      `json:"labels" db:"labels"`
	Annotations       KubernetesAnnotations `json:"annotations" db:"annotations"`
	OwnerReferences   OwnerReferences       `json:"owner_references" db:"owner_references"`
}

func NewObjectMeta(o metav1.ObjectMeta) ObjectMeta {
	return ObjectMeta{
		Name:              o.Name,
		Namespace:         o.Namespace,
		Labels:            getLabels(o),
		Annotations:       filterAnnotations(o),
		CreationTimestamp: o.CreationTimestamp,
		OwnerReferences:   getOwnerReferences(o),
	}
}

type OwnerReferences []OwnerReference

type OwnerReference struct {
	Kind       string `json:"kind"`
	APIVersion string `json:"api_version"`
	Name       string `json:"name"`
	Controller *bool  `json:"controller,omitempty"`
}

func (o *OwnerReferences) Scan(val interface{}) error {
	b, ok := val.([]byte)
	if !ok {
		return errors.New("type assertion to []byte failed")
	}

	return json.Unmarshal(b, &o)
}

func (o *OwnerReferences) Value() (driver.Value, error) {
	bytes, err := json.Marshal(o)
	return bytes, err
}

type TypeMeta struct {
	Kind         string `json:"kind" db:"kind"`
	APIGroup     string `json:"api_group" db:"api_group"`
	ResourceType string `json:"resource_type" db:"resource_type"`
	APIVersion   string `json:"api_version" db:"api_version"`
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

type LabelSelectorRequirement struct {
	Key      string   `json:"key"`
	Operator string   `json:"operator"`
	Values   []string `json:"values,omitempty"`
}

type LabelSelector struct {
	MatchLabels      map[string]string          `json:"matchLabels,omitempty"`
	MatchExpressions []LabelSelectorRequirement `json:"matchExpressions,omitempty"`
}

func filterAnnotations(o metav1.ObjectMeta) (a map[string]string) {
	a = o.GetAnnotations()
	if len(a) > 0 {
		delete(a, "kubectl.kubernetes.io/last-applied-configuration")
	} else {
		a = make(map[string]string)
	}
	return a
}

func getOwnerReferences(o metav1.ObjectMeta) []OwnerReference {
	refs := make([]OwnerReference, 0)
	for _, v := range o.OwnerReferences {
		ref := OwnerReference{
			Kind:       v.Kind,
			APIVersion: v.APIVersion,
			Name:       v.Name,
			Controller: v.Controller,
		}
		refs = append(refs, ref)
	}
	return refs
}

func getLabels(o metav1.ObjectMeta) map[string]string {
	labels := o.GetLabels()
	if len(labels) == 0 {
		labels = make(map[string]string, 0)
	}
	return labels
}
