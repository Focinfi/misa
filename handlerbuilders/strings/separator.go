package strings

import (
	"fmt"

	"github.com/Focinfi/go-pipeline"
	"github.com/Focinfi/misa/handlerbuilders/confparam"
	"github.com/Focinfi/misa/handlerbuilders/utils"
)

var separatorConfParams = make(map[string]confparam.ConfParam)

func init() {
	params, err := confparam.GetConfParams(SeparatorConf{})
	if err != nil {
		panic(err)
	}
	separatorConfParams = params
}

type SeparatorConf struct {
	Type      string `json:"type" validate:"required"`
	Separator string `json:"separator" validate:"required"`
}

func (separator SeparatorConf) Build() (pipeline.Handler, error) {
	switch separator.Type {
	case "join":
		return Join{SeparatorConf: separator}, nil
	case "split":
		return Split{SeparatorConf: separator}, nil
	}
	return nil, fmt.Errorf("unsupportted separator type: %v", separator.Type)
}

func (*SeparatorConf) ConfParams() map[string]confparam.ConfParam {
	return separatorConfParams
}

func (separator *SeparatorConf) InitByConf(conf map[string]interface{}) error {
	return utils.JSONUnmarshalWithMap(conf, separator)
}
