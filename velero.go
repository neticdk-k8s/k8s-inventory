package inventory

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type Velero struct {
	Schedules []*VeleroSchedule `json:"schedules"`
	Backups   []*VeleroBackup   `json:"backups"`
}

func NewVelero() *Velero {
	return &Velero{
		Schedules: make([]*VeleroSchedule, 0),
		Backups:   make([]*VeleroBackup, 0),
	}
}

type VeleroSchedule struct {
	TypeMeta
	ObjectMeta

	Spec   VeleroScheduleSpec   `json:"spec"`
	Status VeleroScheduleStatus `json:"status"`
}

type VeleroScheduleSpec struct {
	Schedule           string          `json:"schedule"`
	ExcludedNamespaces []string        `json:"excludedNamespaces"`
	SnapshotVolumes    *bool           `json:"snapshotVolumes"`
	TTL                metav1.Duration `json:"ttl"`
}

type VeleroScheduleStatus struct {
	LastBackup *metav1.Time `json:"lastBackup"`
	Phase      string       `json:"phase"`
}

func NewVeleroSchedule() *VeleroSchedule {
	return &VeleroSchedule{
		TypeMeta: TypeMeta{
			Kind:         "Schedule",
			APIGroup:     "velero.io",
			APIVersion:   "v1",
			ResourceType: "schedules",
		},
		ObjectMeta: NewObjectMeta(metav1.ObjectMeta{}),
		Spec:       VeleroScheduleSpec{},
		Status:     VeleroScheduleStatus{},
	}
}

type VeleroBackup struct {
	TypeMeta
	ObjectMeta

	Spec   VeleroBackupSpec   `json:"spec"`
	Status VeleroBackupStatus `json:"status"`
}

type VeleroBackupSpec struct {
	ScheduleName       string          `json:"scheduleName"`
	ExcludedNamespaces []string        `json:"excludedNamespaces"`
	StorageLocation    string          `json:"storageLocation"`
	SnapshotVolumes    *bool           `json:"snapshotVolumes"`
	TTL                metav1.Duration `json:"ttl"`
}

type VeleroBackupStatus struct {
	StartTimestamp      *metav1.Time `json:"startTimestamp"`
	CompletionTimestamp *metav1.Time `json:"completionTimestamp"`
	Expiration          *metav1.Time `json:"expiration"`
	Phase               string       `json:"phase"`
	ItemsBackedUp       int          `json:"itemsBackedUp"`
	TotalItems          int          `json:"totalItems"`
	Warnings            int          `json:"warnings"`
	Errors              int          `json:"errors"`
	Version             int          `json:"version"`
}

func NewVeleroBackup() *VeleroBackup {
	return &VeleroBackup{
		TypeMeta: TypeMeta{
			Kind:         "Backup",
			APIGroup:     "velero.io",
			APIVersion:   "v1",
			ResourceType: "backups",
		},
		ObjectMeta: NewObjectMeta(metav1.ObjectMeta{}),
		Spec:       VeleroBackupSpec{},
		Status:     VeleroBackupStatus{},
	}
}
