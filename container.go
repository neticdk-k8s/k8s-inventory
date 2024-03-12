package inventory

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type PodTemplateContainer struct {
	Image          string `json:"image"`
	LimitsCPU      int64  `json:"limits_cpu"`
	LimitsMemory   int64  `json:"limits_memory"`
	RequestsCPU    int64  `json:"requests_cpu"`
	RequestsMemory int64  `json:"requests_memory"`
}

func NewPodTemplateContainer() *PodTemplateContainer {
	return &PodTemplateContainer{}
}

type ContainerStatus struct {
	Name    string         `json:"name"`
	State   ContainerState `json:"state,omitempty"`
	Ready   bool           `json:"ready"`
	Image   string         `json:"image"`
	ImageID string         `json:"imageID"`
}

type ContainerState struct {
	Waiting    *ContainerStateWaiting    `json:"waiting,omitempty"`
	Running    *ContainerStateRunning    `json:"running,omitempty"`
	Terminated *ContainerStateTerminated `json:"terminated,omitempty"`
}

type ContainerStateWaiting struct {
	Reason  string `json:"reason,omitempty"`
	Message string `json:"message,omitempty"`
}

type ContainerStateRunning struct {
	StartedAt metav1.Time `json:"startedAt,omitempty"`
}

type ContainerStateTerminated struct {
	ExitCode    int32       `json:"exitCode"`
	Signal      int32       `json:"signal,omitempty"`
	Reason      string      `json:"reason,omitempty"`
	Message     string      `json:"message,omitempty"`
	StartedAt   metav1.Time `json:"startedAt,omitempty"`
	FinishedAt  metav1.Time `json:"finishedAt,omitempty"`
	ContainerID string      `json:"containerID,omitempty"`
}

type Container struct {
	Name            string               `json:"name"`
	Image           string               `json:"image,omitempty"`
	Command         []string             `json:"command,omitempty"`
	Args            []string             `json:"args,omitempty"`
	WorkingDir      string               `json:"workingDir,omitempty"`
	Ports           []ContainerPort      `json:"ports,omitempty"`
	Resources       ResourceRequirements `json:"resources,omitempty"`
	RestartPolicy   *string              `json:"restartPolicy,omitempty"`
	VolumeMounts    []VolumeMount        `json:"volumeMounts,omitempty"`
	ImagePullPolicy string               `json:"imagePullPolicy,omitempty"`
	SecurityContext *SecurityContext     `json:"securityContext,omitempty"`
}

type ContainerPort struct {
	Name          string `json:"name,omitempty"`
	HostPort      int32  `json:"hostPort,omitempty"`
	ContainerPort int32  `json:"containerPort"`
	Protocol      string `json:"protocol,omitempty"`
	HostIP        string `json:"hostIP,omitempty"`
}

type ResourceRequirements struct {
	LimitsCPU                int64 `json:"limitsCPU"`
	LimitsMemory             int64 `json:"limitsMemory"`
	RequestsCPU              int64 `json:"requestsCPU"`
	RequestsMemory           int64 `json:"requestsMemory"`
	RequestsStorage          int64 `json:"requestsStorage"`
	LimitsStorage            int64 `json:"limitsStorage"`
	RequestsStorageEphemeral int64 `json:"requestsStorageEphemeral"`
	LimitsStorageEphemeral   int64 `json:"limitsStorageEphemeral"`
}

type VolumeMount struct {
	Name        string `json:"name"`
	ReadOnly    bool   `json:"readOnly,omitempty"`
	MountPath   string `json:"mountPath"`
	SubPath     string `json:"subPath,omitempty"`
	SubPathExpr string `json:"subPathExpr,omitempty"`
}

type SecurityContext struct {
	Capabilities             *Capabilities `json:"capabilities,omitempty"`
	Privileged               *bool         `json:"privileged,omitempty"`
	RunAsUser                *int64        `json:"runAsUser,omitempty"`
	RunAsGroup               *int64        `json:"runAsGroup,omitempty"`
	RunAsNonRoot             *bool         `json:"runAsNonRoot,omitempty"`
	ReadOnlyRootFilesystem   *bool         `json:"readOnlyRootFilesystem,omitempty"`
	AllowPrivilegeEscalation *bool         `json:"allowPrivilegeEscalation,omitempty"`
}

type Capabilities struct {
	Add  []string `json:"add,omitempty"`
	Drop []string `json:"drop,omitempty"`
}
