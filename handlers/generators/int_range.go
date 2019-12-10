package generators

import (
	"context"
	"errors"

	"github.com/Focinfi/go-pipeline"
)

var DefaultIntRange = IntRange{}

type IntRange struct{}

func (ir IntRange) Handle(ctx context.Context, reqRes *pipeline.HandleRes) (respRes *pipeline.HandleRes, err error) {
	if reqRes != nil && reqRes.Data != nil {
		respRes, err = reqRes.Copy()
		if err != nil {
			return nil, err
		}

		conf := reqRes.Data.(map[string]int)
		start := conf["start"]
		end := conf["end"]
		step := conf["step"]
		rt := make([]int, 0)
		cur := start
		for cur <= end {
			rt = append(rt, cur)
			cur += step
		}

		respRes.Status = pipeline.HandleStatusOK
		respRes.Data = rt
		return respRes, nil
	}
	return nil, errors.New("empty request data")
}
