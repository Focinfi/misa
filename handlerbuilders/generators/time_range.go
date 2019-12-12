package generators

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/araddon/dateparse"

	"github.com/Focinfi/go-pipeline"
)

var DefaultTimeRange = TimeRange{}

type TimeRange struct{}

func (tr TimeRange) Handle(ctx context.Context, reqRes *pipeline.HandleRes) (respRes *pipeline.HandleRes, err error) {
	if reqRes != nil && reqRes.Data != nil {
		respRes, err = reqRes.Copy()
		if err != nil {
			return nil, err
		}

		conf := reqRes.Data.(map[string]interface{})
		start, err := dateparse.ParseAny(fmt.Sprint(conf["start"]))
		if err != nil {
			return nil, fmt.Errorf("param start is not a time, err: %v", err)
		}
		end, err := dateparse.ParseAny(fmt.Sprint(conf["end"]))
		if err != nil {
			return nil, fmt.Errorf("param end is not a time, err: %v", err)
		}
		step, err := time.ParseDuration(fmt.Sprint(conf["step"]))
		if err != nil {
			return nil, fmt.Errorf("param start is not a time, err: %v", err)
		}
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
