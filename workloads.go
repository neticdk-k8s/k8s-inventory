package inventory

type Workload struct {
	Deployments  []*Deployment  `json:"deployments"`
	StatefulSets []*StatefulSet `json:"statefulsets"`
	DaemonSets   []*DaemonSet   `json:"daemonsets"`
	ReplicaSets  []*ReplicaSet  `json:"replicasets"`
	CronJobs     []*CronJob     `json:"cronjobs"`
}

func NewWorkload() *Workload {
	return &Workload{
		Deployments:  make([]*Deployment, 0),
		StatefulSets: make([]*StatefulSet, 0),
		DaemonSets:   make([]*DaemonSet, 0),
		ReplicaSets:  make([]*ReplicaSet, 0),
		CronJobs:     make([]*CronJob, 0),
	}
}
