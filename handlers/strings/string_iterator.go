package strings

import (
	"context"

	"github.com/Focinfi/go-pipeline"
	"github.com/Focinfi/misa/handlers/iterators"
)

type IteratorConf struct {
	Type            string `json:"type"`
	InterpreterName string `json:"interpreter_name"`
	Script          string `json:"script"`
}

type StringIterator struct {
	Separator     string             `json:"separator"`
	IteratorConfs []IteratorConf     `json:"iterators"`
	handlers      []pipeline.Handler `json:"-"`
}

func NewStringIterator(separator string, iteratorConfs []IteratorConf) (*StringIterator, error) {
	handlers := make([]pipeline.Handler, 0, 1+len(iteratorConfs))
	splitter := String{Separator: separator}
	handlers = append(handlers, splitter)

	for _, conf := range iteratorConfs {
		iterator, err := iterators.NewIterator(conf.Type, conf.InterpreterName, conf.Script)
		if err != nil {
			return nil, err
		}
		handlers = append(handlers, iterator)
	}
	return &StringIterator{
		Separator:     separator,
		IteratorConfs: iteratorConfs,
		handlers:      handlers,
	}, nil
}

func (si StringIterator) Handle(ctx context.Context, reqRes *pipeline.HandleRes) (respRes *pipeline.HandleRes, err error) {
	for _, handler := range si.handlers {
		respRes, err = handler.Handle(ctx, reqRes)
		if err != nil {
			return nil, err
		}
		reqRes = respRes
	}
	return
}
