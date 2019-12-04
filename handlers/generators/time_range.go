package generators

import (
	"context"
	"errors"
	"time"

	"github.com/Focinfi/go-pipeline"
)

type TimeRange struct{}

func (tr TimeRange) Handle(ctx context.Context, reqRes *pipeline.HandleRes) (respRes *pipeline.HandleRes, err error) {
	if reqRes != nil && reqRes.Data != nil {
		respRes, err = reqRes.Copy()
		if err != nil {
			return nil, err
		}

		conf := reqRes.Data.(map[string]interface{})
		start := conf["start"].(time.Time)
		end := conf["end"].(time.Time)
		step := conf["step"].(time.Duration)
		rt := make([]time.Time, 0)
		cur := start
		for cur.Before(end) || cur.Equal(end) {
			rt = append(rt, cur)
			cur = cur.Add(step)
		}

		respRes.Status = pipeline.HandleStatusOK
		respRes.Data = rt
		return respRes, nil
	}
	return nil, errors.New("empty request data")
}
