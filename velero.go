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
	Name               string          `json:"name"`
	Namespace          string          `json:"namespace"`
	Schedule           string          `json:"schedule"`
	ExcludedNamespaces []string        `json:"excluded_namespaces"`
	SnapshotVolumes    *bool           `json:"snapshot_volumes"`
	TTL                metav1.Duration `json:"ttl"`
	LastBackup         *metav1.Time    `json:"last_backup"`
	Phase              string          `json:"phase"`
}

type VeleroBackup struct {
	Name                string          `json:"name"`
	Namespace           string          `json:"namespace"`
	ScheduleName        string          `json:"schedule_name"`
	ExcludedNamespaces  []string        `json:"excluded_namespaces"`
	StorageLocation     string          `json:"storage_location"`
	SnapshotVolumes     *bool           `json:"snapshot_volumes"`
	TTL                 metav1.Duration `json:"ttl"`
	StartTimestamp      *metav1.Time    `json:"start_timestamp"`
	CompletionTimestamp *metav1.Time    `json:"completion_timestamp"`
	Expiration          *metav1.Time    `json:"expiration"`
	Phase               string          `json:"phase"`
	ItemsBackedUp       int             `json:"items_backed_up"`
	TotalItems          int             `json:"total_items"`
	Warnings            int             `json:"warnings"`
	Errors              int             `json:"errors"`
	Version             int             `json:"version"`
}
