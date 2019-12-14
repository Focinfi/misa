package interpreters

import (
	"fmt"

	"github.com/Focinfi/go-pipeline"
)

type Conf struct {
	Type       string                 `json:"type" desc:"enum: tengo" validate:"required"`
	Script     string                 `json:"script" validate:"required"`
	InitVarMap map[string]interface{} `json:"int_var_map" desc:"init variables in script" validate:"-"`
	RtVarName  string                 `json:"rt_var_name" desc:"returned variable name" validate:"-"`
}

func (c Conf) Build() (pipeline.Handler, error) {
	switch c.Type {
	case "tengo":
		return NewTengo(c)
	default:
		return nil, fmt.Errorf("unsuportted interpreter type(%#v)", c.Type)
	}
}
