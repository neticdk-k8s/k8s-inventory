package inventory

type StorageClass struct {
	TypeMeta
	ObjectMeta

	Provisioner string `json:"provisioner" db:"provisioner"`
}

func NewStorageClass() *StorageClass {
	return &StorageClass{
		ObjectMeta: NewObjectMeta(),
	}
}
