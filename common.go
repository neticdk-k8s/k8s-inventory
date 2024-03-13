package inventory

type IntOrString struct {
	Type   int    `json:"type"`
	IntVal int32  `json:"intVal"`
	StrVal string `json:"strVal"`
}

type RootOwner struct {
	Kind       string `json:"kind"`
	APIGroup   string `json:"apiGroup"`
	APIVersion string `json:"apiVersion"`
	Name       string `json:"name"`
	Namespace  string `json:"namespace,omitempty"`
}
