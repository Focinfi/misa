package combination

import (
	"context"

	"github.com/Focinfi/go-pipeline"
)

type HandlerList struct {
	Handlers []pipeline.Handler `json:"builders"`
}

func (hl HandlerList) Handle(ctx context.Context, reqRes *pipeline.HandleRes) (respRes *pipeline.HandleRes, err error) {
	for _, handler := range hl.Handlers {
		respRes, err = handler.Handle(ctx, reqRes)
		if err != nil {
			return nil, err
		}
		reqRes = respRes
	}
	return
}
