package inventory

import metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

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
		ObjectMeta: NewObjectMeta(metav1.ObjectMeta{}),
	}
}
