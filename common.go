package inventory

type IntOrString struct {
	Type   int    `json:"type"`
	IntVal int32  `json:"int_val"`
	StrVal string `json:"str_val"`
}
