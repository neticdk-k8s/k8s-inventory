package inventory

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type IntOrString struct {
	Type   int    `json:"type"`
	IntVal int32  `json:"intVal"`
	StrVal string `json:"strVal"`
}

type Reference struct {
	Kind       string `json:"kind"`
	APIGroup   string `json:"apiGroup"`
	APIVersion string `json:"apiVersion"`
	Name       string `json:"name"`
	Namespace  string `json:"namespace,omitempty"`
}

type RootOwner Reference

// Object represents an object with Kubernetes metadata associated
type Object interface {
	GetKind() string
	GetAPIGroup() string
	GetResourceType() string
	GetAPIVersion() string
	GetName() string
	GetNamespace() string
	SetFromObjectMeta(om metav1.ObjectMeta)
	EqualTo(o Object) bool
}

type PartialObject struct {
	TypeMeta
	ObjectMeta
}

func (p *PartialObject) EqualTo(o Object) bool {
	return p.APIGroup == o.GetAPIGroup() && p.APIVersion == o.GetAPIVersion() && p.Kind == o.GetKind() &&
		p.Name == o.GetName() && p.Namespace == o.GetNamespace()
}
