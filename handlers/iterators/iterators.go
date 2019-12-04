package iterators

import (
	"context"
	"errors"

	"github.com/Focinfi/misa/handlers/combination"

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

type Conf struct {
	Type            string `json:"type"`
	InterpreterName string `json:"interpreter_name"`
	Script          string `json:"script"`
}

type Iterators struct {
	Confs   []Conf                  `json:"confs"`
	handler combination.HandlerList `json:"-"`
}

func NewIterators(iteratorConfs []Conf) (*Iterators, error) {
	handlers := make([]pipeline.Handler, 0, len(iteratorConfs))
	for _, conf := range iteratorConfs {
		iterator, err := NewIterator(conf.Type, conf.InterpreterName, conf.Script)
		if err != nil {
			return nil, err
		}
		handlers = append(handlers, iterator)
	}
	return &Iterators{
		Confs:   iteratorConfs,
		handler: combination.HandlerList{Handlers: handlers},
	}, nil
}

func (it Iterators) Handle(ctx context.Context, reqRes *pipeline.HandleRes) (respRes *pipeline.HandleRes, err error) {
	return it.handler.Handle(ctx, reqRes)
}
