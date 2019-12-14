package interpreters

import (
	"fmt"

	"github.com/Focinfi/go-pipeline"
	"github.com/Focinfi/misa/handlerbuilders/confparam"
	"github.com/Focinfi/misa/handlerbuilders/utils"
)

var confParams = map[string]confparam.ConfParam{}

func init() {
	params, err := confparam.GetConfParams(Conf{})
	if err != nil {
		panic(err)
	}
	confParams = params
}

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

func (c Conf) ConfParams() map[string]confparam.ConfParam {
	return confParams
}

func (c *Conf) InitByConf(conf map[string]interface{}) error {
	if err := utils.JSONUnmarshalWithMap(conf, c); err != nil {
		return fmt.Errorf("init interpreters.Conf err: %v", err)
	}
	return nil
}
