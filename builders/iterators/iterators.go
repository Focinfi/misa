package iterators

import (
	"fmt"

	"github.com/Focinfi/go-pipeline"
)

type Conf struct {
	Type            string `json:"type" desc:"enum: map|reduce|select" validate:"required"`
	InterpreterName string `json:"interpreter_name" desc:"interpreter name" validate:"required"`
	Script          string `json:"script" desc:"iterator script" validate:"required"`
}

func (itr Conf) Build() (pipeline.Handler, error) {
	return NewIterator(itr)
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
