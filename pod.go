package inventory

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type Pod struct {
	TypeMeta
	ObjectMeta

	Spec   PodSpec   `json:"spec,omitempty"`
	Status PodStatus `json:"status,omitempty"`

	RootOwner *RootOwner `json:"rootOwner"`
}

type RootOwner struct {
	Kind       string `json:"kind"`
	APIGroup   string `json:"apiGroup"`
	APIVersion string `json:"apiVersion"`
	Name       string `json:"name"`
	Namespace  string `json:"namespace,omitempty"`
}

func NewPod() *Pod {
	return &Pod{
		TypeMeta: TypeMeta{
			Kind:         "Pod",
			APIGroup:     "core",
			APIVersion:   "v1",
			ResourceType: "pods",
		},
		ObjectMeta: NewObjectMeta(metav1.ObjectMeta{}),
		Spec:       PodSpec{},
		Status:     PodStatus{},
	}
}

func NewPods() []*Pod {
	return make([]*Pod, 0)
}

type PodSpec struct {
	Volumes            []Volume            `json:"volumes,omitempty"`
	InitContainers     []Container         `json:"initContainers,omitempty"`
	Containers         []Container         `json:"containers"`
	RestartPolicy      string              `json:"restartPolicy,omitempty"`
	ServiceAccountName string              `json:"serviceAccountName,omitempty"`
	NodeName           string              `json:"nodeName,omitempty"`
	HostNetwork        bool                `json:"hostNetwork,omitempty"`
	SecurityContext    *PodSecurityContext `json:"securityContext,omitempty"`
	PriorityClassName  string              `json:"priorityClassName,omitempty"`
	Priority           *int32              `json:"priority,omitempty"`
	PreemptionPolicy   *string             `json:"preemptionPolicy,omitempty"`
}

type Volume struct {
	Name   string `json:"name"`
	Source string `json:"source"`
}

type PodSecurityContext struct {
	RunAsUser          *int64  `json:"runAsUser,omitempty"`
	RunAsGroup         *int64  `json:"runAsGroup,omitempty"`
	RunAsNonRoot       *bool   `json:"runAsNonRoot,omitempty"`
	SupplementalGroups []int64 `json:"supplementalGroups,omitempty"`
}

type PodStatus struct {
	Phase                 string            `json:"phase,omitempty"`
	Conditions            []PodCondition    `json:"conditions,omitempty"`
	PodIP                 string            `json:"podIP,omitempty"`
	StartTime             *metav1.Time      `json:"startTime,omitempty"`
	InitContainerStatuses []ContainerStatus `json:"initContainerStatuses,omitempty"`
	ContainerStatuses     []ContainerStatus `json:"containerStatuses,omitempty"`
	QOSClass              string            `json:"qosClass,omitempty"`
}

type PodCondition struct {
	Type    string `json:"type"`
	Status  string `json:"status"`
	Message string `json:"message,omitempty"`
}

type PodTemplate struct {
	Containers     []*PodTemplateContainer `json:"containers"`
	InitContainers []*PodTemplateContainer `json:"init_containers"`
}

func NewPodTemplate() *PodTemplate {
	return &PodTemplate{
		Containers:     make([]*PodTemplateContainer, 0),
		InitContainers: make([]*PodTemplateContainer, 0),
	}
}
