package inventory

import (
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type PersistentVolume struct {
	TypeMeta
	ObjectMeta

	Spec   PersistentVolumeSpec   `json:"spec" db:"spec"`
	Status PersistentVolumeStatus `json:"status" db:"status"`

	Source   string `json:"source" db:"source"`
	Driver   string `json:"driver" db:"driver"`
	Path     string `json:"path" db:"path"`
	FSType   string `json:"fs_type" db:"fs_type"`
	VolumeID string `json:"volume_id" db:"volume_id"`
}

type PersistentVolumeSpec struct {
	Capacity         int64  `json:"capacity"`
	AccessModes      string `json:"access_modes"`
	Claim            string `json:"claim"`
	StorageClassName string `json:"storage_class"`
	VolumeMode       string `json:"volume_mode"`
}

type PersistentVolumeStatus struct {
	Phase   string `json:"phase"`
	Message string `json:"message"`
}

func NewPersistentVolume() *PersistentVolume {
	return &PersistentVolume{
		TypeMeta: TypeMeta{
			Kind:         "PersistentVolume",
			APIGroup:     "core",
			APIVersion:   "v1",
			ResourceType: "persistentvolumes",
		},
		ObjectMeta: NewObjectMeta(metav1.ObjectMeta{}),
		Spec:       PersistentVolumeSpec{},
		Status:     PersistentVolumeStatus{},
	}
}

func (p *PersistentVolume) SetPersistentVolumeSource(o v1.PersistentVolume) {
	switch {
	case o.Spec.PersistentVolumeSource.GCEPersistentDisk != nil:
		p.Source = "GCEPersistentDisk"
	case o.Spec.PersistentVolumeSource.AWSElasticBlockStore != nil:
		p.Source = "AWSElasticBlockStore"
		p.FSType = o.Spec.PersistentVolumeSource.AWSElasticBlockStore.FSType
		p.VolumeID = o.Spec.PersistentVolumeSource.AWSElasticBlockStore.VolumeID
	case o.Spec.PersistentVolumeSource.HostPath != nil:
		p.Source = "HostPath"
		p.Path = o.Spec.PersistentVolumeSource.HostPath.Path
	case o.Spec.PersistentVolumeSource.Glusterfs != nil:
		p.Source = "Glusterfs"
	case o.Spec.PersistentVolumeSource.NFS != nil:
		p.Source = "NFS"
	case o.Spec.PersistentVolumeSource.RBD != nil:
		p.Source = "RBD"
	case o.Spec.PersistentVolumeSource.ISCSI != nil:
		p.Source = "ISCSI"
	case o.Spec.PersistentVolumeSource.Cinder != nil:
		p.Source = "Cinder"
	case o.Spec.PersistentVolumeSource.CephFS != nil:
		p.Source = "CephFS"
	case o.Spec.PersistentVolumeSource.FC != nil:
		p.Source = "FC"
	case o.Spec.PersistentVolumeSource.Flocker != nil:
		p.Source = "Flocker"
	case o.Spec.PersistentVolumeSource.FlexVolume != nil:
		p.Source = "FlexVolume"
	case o.Spec.PersistentVolumeSource.AzureFile != nil:
		p.Source = "AzureFile"
	case o.Spec.PersistentVolumeSource.VsphereVolume != nil:
		p.Source = "VsphereVolume"
	case o.Spec.PersistentVolumeSource.Quobyte != nil:
		p.Source = "Quobyte"
	case o.Spec.PersistentVolumeSource.AzureDisk != nil:
		p.Source = "AzureDisk"
		p.FSType = *o.Spec.PersistentVolumeSource.AzureDisk.FSType
	case o.Spec.PersistentVolumeSource.PhotonPersistentDisk != nil:
		p.Source = "PhotonPersistentDisk"
	case o.Spec.PersistentVolumeSource.PortworxVolume != nil:
		p.Source = "PortworxVolume"
	case o.Spec.PersistentVolumeSource.ScaleIO != nil:
		p.Source = "ScaleIO"
	case o.Spec.PersistentVolumeSource.Local != nil:
		p.Source = "Local"
		p.FSType = *o.Spec.PersistentVolumeSource.Local.FSType
		p.Path = o.Spec.PersistentVolumeSource.Local.Path
	case o.Spec.PersistentVolumeSource.StorageOS != nil:
		p.Source = "StorageOS"
	case o.Spec.PersistentVolumeSource.CSI != nil:
		p.Source = "CSI"
		p.FSType = o.Spec.PersistentVolumeSource.CSI.FSType
		p.Driver = o.Spec.PersistentVolumeSource.CSI.Driver
	default:
		p.Source = "Unknown"
	}
}
