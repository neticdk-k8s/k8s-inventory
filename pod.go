package inventory

import (
	"strings"

	v1 "k8s.io/api/core/v1"
)

type Pod struct {
	PodName         string `json:"pod_name"`
	PodNameFull     string `json:"pod_name_full"`
	AppName         string `json:"app_name"`
	Owner           string `json:"owner"`
	FullVersion     string `json:"full_version"`
	SemanticVersion string `json:"semantic_version"`
}

func NewPod(pod v1.Pod) (pi *Pod) {
	pi = &Pod{
		PodNameFull: pod.Name,
		PodName:     pod.Name,
	}
	if pod.OwnerReferences != nil && len(pod.OwnerReferences) > 0 {
		pi.Owner = pod.OwnerReferences[0].Kind + "/" + pod.OwnerReferences[0].Name
		pi.PodName = getPodNameFromKind(pod.OwnerReferences[0].Kind, pod.Name)
	}
	return
}

func getPodNameFromKind(kind string, fullName string) string {
	parts := strings.Split(fullName, "-")
	if kind == "ReplicaSet" {
		return strings.Join(parts[:len(parts)-2], "-")
	} else if kind == "DaemonSet" || kind == "StatefulSet" || kind == "Job" || kind == "Node" {
		return strings.Join(parts[:len(parts)-1], "-")
	}
	return fullName
}

type PodTemplate struct {
	Containers     []*Container `json:"containers"`
	InitContainers []*Container `json:"init_containers"`
}

func NewPodTemplate() *PodTemplate {
	return &PodTemplate{
		Containers:     make([]*Container, 0),
		InitContainers: make([]*Container, 0),
	}
}
