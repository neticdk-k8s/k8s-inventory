package inventory

type StorageClass struct {
	TypeMeta
	ObjectMeta

	Provisioner string `json:"provisioner" db:"provisioner"`
}

func NewStorageClass() *StorageClass {
	return &StorageClass{
		TypeMeta: TypeMeta{
			Kind:         "StorageClass",
			APIGroup:     "storage.k8s.io",
			APIVersion:   "v1",
			ResourceType: "storageclasses",
		},
		ObjectMeta: NewObjectMeta(),
	}
}
