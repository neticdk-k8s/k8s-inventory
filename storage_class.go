package inventory

import metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

type StorageClass struct {
	PartialObject

	Provisioner string `json:"provisioner"`
}

type StorageClasses = Set[*StorageClass]

func NewStorageClass() *StorageClass {
	return &StorageClass{
		PartialObject: PartialObject{
			TypeMeta: TypeMeta{
				Kind:         "StorageClass",
				APIGroup:     "storage.k8s.io",
				APIVersion:   "v1",
				ResourceType: "storageclasses",
			},
			ObjectMeta: NewObjectMeta(metav1.ObjectMeta{}),
		},
	}
}
