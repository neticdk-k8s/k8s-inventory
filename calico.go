package inventory

import metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

type CalicoClusterInformation struct {
	PartialObject

	Spec   CalicoClusterInformationSpec   `json:"spec"`
	Status CalicoClusterInformationStatus `json:"status"`
}

type CalicoClusterInformationSpec struct {
	Version string `json:"version"`
}

type CalicoClusterInformationStatus struct{}

func NewCalicoClusterInformation() *CalicoClusterInformation {
	return &CalicoClusterInformation{
		PartialObject: PartialObject{
			TypeMeta: TypeMeta{
				Kind:         "ClusterInformation",
				APIGroup:     "crd.projectcalico.org",
				APIVersion:   "v1",
				ResourceType: "clusterinformations",
			},
			ObjectMeta: NewObjectMeta(metav1.ObjectMeta{}),
		},
		Spec:   CalicoClusterInformationSpec{},
		Status: CalicoClusterInformationStatus{},
	}
}
