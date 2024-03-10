package inventory

type IntOrString struct {
	Type   int    `json:"type"`
	IntVal int32  `json:"intVal"`
	StrVal string `json:"strVal"`
}
