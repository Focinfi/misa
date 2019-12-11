package strings

import (
	"context"

	"github.com/Focinfi/go-pipeline"
	"github.com/Focinfi/misa/handlerbuilders/iterators"
)

type StringIterator struct {
	Separator     string             `json:"separator"`
	IteratorConfs []iterators.Conf   `json:"iterators"`
	handlers      []pipeline.Handler `json:"-"`
}

func NewStringIterator(separator string, iteratorConfs []iterators.Conf) (*StringIterator, error) {
	handlers := make([]pipeline.Handler, 0)
	splitter := Splitter{Separator: separator}
	iterators, err := iterators.NewIterators(iteratorConfs)
	if err != nil {
		return nil, err
	}

	handlers = append(handlers, splitter)
	handlers = append(handlers, iterators)
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
