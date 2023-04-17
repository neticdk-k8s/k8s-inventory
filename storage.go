package inventory

type Storage struct {
	StorageClasses    []*StorageClass `json:"storage_classes"`
	PersistentVolumes []*PV           `json:"persistent_volumes"`
}

func NewStorage() *Storage {
	return &Storage{
		StorageClasses:    make([]*StorageClass, 0),
		PersistentVolumes: make([]*PV, 0),
	}
}
