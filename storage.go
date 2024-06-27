package inventory

type Storage struct {
	StorageClasses    StorageClasses    `json:"storageClasses"`
	PersistentVolumes PersistentVolumes `json:"persistentVolumes"`
}

func NewStorage() *Storage {
	return &Storage{
		StorageClasses:    StorageClasses{},
		PersistentVolumes: PersistentVolumes{},
	}
}

func (s *Storage) AddStorageClass(sc *StorageClass) {
	s.StorageClasses = s.StorageClasses.Add(sc)
}

func (s *Storage) DeleteStorageClass(sc *StorageClass) {
	s.StorageClasses = s.StorageClasses.Delete(sc)
}

func (s *Storage) AddPersistentVolume(pv *PersistentVolume) {
	s.PersistentVolumes = s.PersistentVolumes.Add(pv)
}

func (s *Storage) DeletePersistentVolume(pv *PersistentVolume) {
	s.PersistentVolumes = s.PersistentVolumes.Delete(pv)
}
