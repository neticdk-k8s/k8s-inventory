package inventory

import (
	v1 "k8s.io/api/core/v1"
)

type PV struct {
	TypeMeta
	ObjectMeta

	Spec   PVSpec   `json:"spec" db:"spec"`
	Status PVStatus `json:"status" db:"status"`

	Source   string `json:"source" db:"source"`
	Driver   string `json:"driver" db:"driver"`
	Path     string `json:"path" db:"path"`
	FSType   string `json:"fs_type" db:"fs_type"`
	VolumeID string `json:"volume_id" db:"volume_id"`
}

type PVSpec struct {
	Capacity         int64    `json:"capacity"`
	AccessModes      []string `json:"access_modes"`
	Claim            string   `json:"claim"`
	StorageClassName string   `json:"storage_class"`
	VolumeMode       string   `json:"volume_mode"`
}

type PVStatus struct {
	Phase   string `json:"phase"`
	Message string `json:"message"`
}

func NewPV() *PV {
	return &PV{
		ObjectMeta: NewObjectMeta(),
		Spec:       PVSpec{},
		Status:     PVStatus{},
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
