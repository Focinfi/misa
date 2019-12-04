package iterators

import (
	"errors"

	"github.com/Focinfi/go-pipeline"
)

func NewIterator(t, interpreterName, script string) (pipeline.Handler, error) {
	switch t {
	case "select":
		return NewSelector(interpreterName, script)
	case "map":
		return NewMapper(interpreterName, script)
	case "reduce":
		return NewReducer(interpreterName, script)
	default:
		return nil, errors.New("unsupported iterator type")
	}
}
