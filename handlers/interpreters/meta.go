package interpreters

type Meta struct {
	Script     string                 `json:"script"`
	InitVarMap map[string]interface{} `json:"int_var_map"`
	RtVarName  string                 `json:"rt_var_name"`
}

func (meta Meta) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"script":       meta.Script,
		"init_var_map": meta.InitVarMap,
		"rt_var_name":  meta.RtVarName,
	}
}
