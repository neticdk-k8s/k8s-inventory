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

// Set is a set like structure based on the EqualTo methods of the given objects
type Set[I Object] []I

func (s Set[I]) Add(obj Object) Set[I] {
	o := obj.(I)
	for n, e := range s {
		if obj.EqualTo(e) {
			s[n] = o
			return s
		}
	}
	return append(s, obj.(I))
}

func (s Set[I]) Delete(obj Object) Set[I] {
	for n, e := range s {
		if obj.EqualTo(e) {
			return append(s[:n], s[n+1:]...)
		}
	}
	return s
}
