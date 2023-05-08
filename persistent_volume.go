package inventory

import (
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type PV struct {
	Name              string            `json:"name" db:"name"`
	CreationTimestamp metav1.Time       `json:"creation_timestamp" db:"creation_timestamp"`
	Labels            map[string]string `json:"labels" db:"labels"`
	Annotations       map[string]string `json:"annotations" db:"annotations"`
	StorageClass      string            `json:"storage_class" db:"storage_class"`
	Claim             string            `json:"claim" db:"claim"`
	Status            string            `json:"status" db:"status"`
	AccessModes       string            `json:"access_modes" db:"access_modes"`
	VolumeMode        string            `json:"volume_mode" db:"volume_mode"`
	Capacity          int64             `json:"capacity" db:"capacity"`
	Source            string            `json:"source" db:"source"`
	Driver            string            `json:"driver" db:"driver"`
	Path              string            `json:"path" db:"path"`
	FSType            string            `json:"fs_type" db:"fs_type"`
	VolumeID          string            `json:"volume_id" db:"volume_id"`
}

func NewPV() *PV {
	return &PV{
		Labels:      make(map[string]string, 0),
		Annotations: make(map[string]string, 0),
	}
}

func (i *PV) SetPersistentVolumeSource(o v1.PersistentVolume) {
	switch {
	case o.Spec.PersistentVolumeSource.GCEPersistentDisk != nil:
		i.Source = "GCEPersistentDisk"
	case o.Spec.PersistentVolumeSource.AWSElasticBlockStore != nil:
		i.Source = "AWSElasticBlockStore"
		i.FSType = o.Spec.PersistentVolumeSource.AWSElasticBlockStore.FSType
		i.VolumeID = o.Spec.PersistentVolumeSource.AWSElasticBlockStore.VolumeID
	case o.Spec.PersistentVolumeSource.HostPath != nil:
		i.Source = "HostPath"
		i.Path = o.Spec.PersistentVolumeSource.HostPath.Path
	case o.Spec.PersistentVolumeSource.Glusterfs != nil:
		i.Source = "Glusterfs"
	case o.Spec.PersistentVolumeSource.NFS != nil:
		i.Source = "NFS"
	case o.Spec.PersistentVolumeSource.RBD != nil:
		i.Source = "RBD"
	case o.Spec.PersistentVolumeSource.ISCSI != nil:
		i.Source = "ISCSI"
	case o.Spec.PersistentVolumeSource.Cinder != nil:
		i.Source = "Cinder"
	case o.Spec.PersistentVolumeSource.CephFS != nil:
		i.Source = "CephFS"
	case o.Spec.PersistentVolumeSource.FC != nil:
		i.Source = "FC"
	case o.Spec.PersistentVolumeSource.Flocker != nil:
		i.Source = "Flocker"
	case o.Spec.PersistentVolumeSource.FlexVolume != nil:
		i.Source = "FlexVolume"
	case o.Spec.PersistentVolumeSource.AzureFile != nil:
		i.Source = "AzureFile"
	case o.Spec.PersistentVolumeSource.VsphereVolume != nil:
		i.Source = "VsphereVolume"
	case o.Spec.PersistentVolumeSource.Quobyte != nil:
		i.Source = "Quobyte"
	case o.Spec.PersistentVolumeSource.AzureDisk != nil:
		i.Source = "AzureDisk"
		i.FSType = *o.Spec.PersistentVolumeSource.AzureDisk.FSType
	case o.Spec.PersistentVolumeSource.PhotonPersistentDisk != nil:
		i.Source = "PhotonPersistentDisk"
	case o.Spec.PersistentVolumeSource.PortworxVolume != nil:
		i.Source = "PortworxVolume"
	case o.Spec.PersistentVolumeSource.ScaleIO != nil:
		i.Source = "ScaleIO"
	case o.Spec.PersistentVolumeSource.Local != nil:
		i.Source = "Local"
		i.FSType = *o.Spec.PersistentVolumeSource.Local.FSType
		i.Path = o.Spec.PersistentVolumeSource.Local.Path
	case o.Spec.PersistentVolumeSource.StorageOS != nil:
		i.Source = "StorageOS"
	case o.Spec.PersistentVolumeSource.CSI != nil:
		i.Source = "CSI"
		i.FSType = o.Spec.PersistentVolumeSource.CSI.FSType
		i.Driver = o.Spec.PersistentVolumeSource.CSI.Driver
	default:
		i.Source = "Unknown"
	}
}
