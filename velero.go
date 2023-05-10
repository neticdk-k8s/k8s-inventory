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

	Schedule           string          `json:"schedule" db:"schedule"`
	ExcludedNamespaces []string        `json:"excluded_namespaces"`
	SnapshotVolumes    *bool           `json:"snapshot_volumes" db:"snapshot_volumes"`
	TTL                metav1.Duration `json:"ttl" db:"ttl"`
	LastBackup         *metav1.Time    `json:"last_backup" db:"last_backup"`
	Phase              string          `json:"phase" db:"phase"`
}

type VeleroBackup struct {
	TypeMeta
	ObjectMeta

	ScheduleName        string          `json:"schedule_name" db:"schedule_name"`
	ExcludedNamespaces  []string        `json:"excluded_namespaces"`
	StorageLocation     string          `json:"storage_location" db:"storage_location"`
	SnapshotVolumes     *bool           `json:"snapshot_volumes" db:"snapshot_volumes"`
	TTL                 metav1.Duration `json:"ttl" db:"ttl"`
	StartTimestamp      *metav1.Time    `json:"start_timestamp" db:"start_timestamp"`
	CompletionTimestamp *metav1.Time    `json:"completion_timestamp" db:"completion_timestamp"`
	Expiration          *metav1.Time    `json:"expiration" db:"expiration"`
	Phase               string          `json:"phase" db:"phase"`
	ItemsBackedUp       int             `json:"items_backed_up" db:"items_backed_up"`
	TotalItems          int             `json:"total_items" db:"total_items"`
	Warnings            int             `json:"warnings" db:"warnings"`
	Errors              int             `json:"errors" db:"errors"`
	Version             int             `json:"version" db:"version"`
}
