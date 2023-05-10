package inventory

type CalicoClusterInformation struct {
	TypeMeta
	ObjectMeta

	Spec   CalicoClusterInformationSpec   `json:"spec" db:"spec"`
	Status CalicoClusterInformationStatus `json:"status" db:"status"`
}

type CalicoClusterInformationSpec struct {
	Version string `json:"version"`
}

type CalicoClusterInformationStatus struct{}

func NewCalicoClusterInformation() *CalicoClusterInformation {
	return &CalicoClusterInformation{
		TypeMeta: TypeMeta{
			Kind:         "ClusterInformation",
			APIGroup:     "crd.projectcalico.org",
			APIVersion:   "v1",
			ResourceType: "clusterinformations",
		},
		ObjectMeta: NewObjectMeta(),
		Spec:       CalicoClusterInformationSpec{},
		Status:     CalicoClusterInformationStatus{},
	}
}
