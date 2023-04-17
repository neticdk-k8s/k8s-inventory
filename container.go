package inventory

type Container struct {
	Image          string `json:"image"`
	LimitsCPU      int64  `json:"limits_cpu"`
	LimitsMemory   int64  `json:"limits_memory"`
	RequestsCPU    int64  `json:"requests_cpu"`
	RequestsMemory int64  `json:"requests_memory"`
}

func NewContainer() *Container {
	return &Container{}
}
