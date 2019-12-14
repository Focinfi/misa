package strings

import (
	"fmt"

	"github.com/Focinfi/go-pipeline"
)

type SeparatorConf struct {
	Type      string `json:"type" desc:"enum: join|split" validate:"required"`
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
