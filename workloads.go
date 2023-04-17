package inventory

type Workload struct {
	Deployments  []*Deployment  `json:"deployments"`
	StatefulSets []*StatefulSet `json:"stateful_sets"`
	DaemonSets   []*DaemonSet   `json:"daemon_sets"`
	ReplicaSets  []*ReplicaSet  `json:"replica_sets"`
	CronJobs     []*CronJob     `json:"cron_jobs"`
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
