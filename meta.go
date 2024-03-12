package inventory

import (
	"database/sql/driver"
	"encoding/json"
	"errors"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type ObjectMeta struct {
	Name              string                `json:"name"`
	Namespace         string                `json:"namespace,omitempty"`
	CreationTimestamp metav1.Time           `json:"creationTimestamp"`
	Labels            KubernetesLabels      `json:"labels"`
	Annotations       KubernetesAnnotations `json:"annotations"`
	OwnerReferences   OwnerReferences       `json:"ownerReferences"`
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
	APIVersion string `json:"apiVersion"`
	Name       string `json:"name"`
	Controller *bool  `json:"controller,omitempty"`
}

type TypeMeta struct {
	Kind         string `json:"kind"`
	APIGroup     string `json:"apiGroup"`
	ResourceType string `json:"resourceType"`
	APIVersion   string `json:"apiVersion"`
}

type KubernetesLabels map[string]string

type KubernetesAnnotations map[string]string

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
