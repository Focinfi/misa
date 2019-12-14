package iterators

import (
	"fmt"

	"github.com/Focinfi/misa/handlerbuilders/confparam"

	"github.com/Focinfi/misa/handlerbuilders/utils"

	"github.com/Focinfi/go-pipeline"
)

var iteratorConfParams = make(map[string]confparam.ConfParam)

func init() {
	params, err := confparam.GetConfParams(Conf{})
	if err != nil {
		panic(err)
	}
	iteratorConfParams = params
}

type Conf struct {
	Type            string `json:"type" desc:"enum: map|reduce|select" validate:"required"`
	InterpreterName string `json:"interpreter_name" desc:"interpreter name" validate:"required"`
	Script          string `json:"script" desc:"iterator script" validate:"required"`
}

func (itr Conf) Build() (pipeline.Handler, error) {
	return NewIterator(itr)
}

func (itr *Conf) ConfParams() map[string]confparam.ConfParam {
	return iteratorConfParams
}

func (itr *Conf) InitByConf(conf map[string]interface{}) error {
	return utils.JSONUnmarshalWithMap(conf, itr)
}

func NewIterator(conf Conf) (pipeline.Handler, error) {
	switch conf.Type {
	case "select":
		return NewSelector(conf)
	case "map":
		return NewMapper(conf)
	case "reduce":
		return NewReducer(conf)
	default:
		return nil, fmt.Errorf("unsupported iterator type: %v", conf.Type)
	}
}
