package inventory

type Storage struct {
	StorageClasses    []*StorageClass     `json:"storageClasses"`
	PersistentVolumes []*PersistentVolume `json:"persistentVolumes"`
}

func NewStorage() *Storage {
	return &Storage{
		StorageClasses:    make([]*StorageClass, 0),
		PersistentVolumes: make([]*PersistentVolume, 0),
	}
}
