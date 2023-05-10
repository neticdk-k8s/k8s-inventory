package inventory

type Storage struct {
	StorageClasses    []*StorageClass     `json:"storageclasses"`
	PersistentVolumes []*PersistentVolume `json:"persistentvolumes"`
}

func NewStorage() *Storage {
	return &Storage{
		StorageClasses:    make([]*StorageClass, 0),
		PersistentVolumes: make([]*PersistentVolume, 0),
	}
}
