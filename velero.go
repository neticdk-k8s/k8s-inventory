package inventory

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type Velero struct {
	Schedules VeleroSchedules `json:"schedules"`
	Backups   VeleroBackups   `json:"backups"`
}

func NewVelero() *Velero {
	return &Velero{
		Schedules: VeleroSchedules{},
		Backups:   VeleroBackups{},
	}
}

func (v *Velero) AddSchedule(s *VeleroSchedule) {
	v.Schedules = v.Schedules.Add(s)
}

func (v *Velero) DeleteSchedule(s *VeleroSchedule) {
	v.Schedules = v.Schedules.Delete(s)
}

func (v *Velero) AddBackup(b *VeleroBackup) {
	v.Backups = v.Backups.Add(b)
}

func (v *Velero) DeleteBackup(b *VeleroBackup) {
	v.Backups = v.Backups.Delete(b)
}

type VeleroSchedule struct {
	PartialObject

	Spec   VeleroScheduleSpec   `json:"spec"`
	Status VeleroScheduleStatus `json:"status"`
}

type VeleroSchedules = Set[*VeleroSchedule]

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
		PartialObject: PartialObject{
			TypeMeta: TypeMeta{
				Kind:         "Schedule",
				APIGroup:     "velero.io",
				APIVersion:   "v1",
				ResourceType: "schedules",
			},
			ObjectMeta: NewObjectMeta(metav1.ObjectMeta{}),
		},
		Spec:   VeleroScheduleSpec{},
		Status: VeleroScheduleStatus{},
	}
}

type VeleroBackup struct {
	PartialObject

	Spec   VeleroBackupSpec   `json:"spec"`
	Status VeleroBackupStatus `json:"status"`
}

type VeleroBackups = Set[*VeleroBackup]

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
		PartialObject: PartialObject{
			TypeMeta: TypeMeta{
				Kind:         "Backup",
				APIGroup:     "velero.io",
				APIVersion:   "v1",
				ResourceType: "backups",
			},
			ObjectMeta: NewObjectMeta(metav1.ObjectMeta{}),
		},
		Spec:   VeleroBackupSpec{},
		Status: VeleroBackupStatus{},
	}
}
