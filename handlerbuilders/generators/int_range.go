package generators

import (
	"context"
	"errors"
	"fmt"

	"github.com/Focinfi/misa/handlerbuilders/utils"

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

		conf := reqRes.Data.(map[string]interface{})
		start, err := utils.AnyTypeToInt64(conf["start"])
		if err != nil {
			return nil, fmt.Errorf("param start is not int, err: %v", err)
		}
		end, err := utils.AnyTypeToInt64(conf["end"])
		if err != nil {
			return nil, fmt.Errorf("param end is not int, err: %v", err)
		}
		step, err := utils.AnyTypeToInt64(conf["step"])
		if err != nil {
			return nil, fmt.Errorf("param step is not int, err: %v", err)
		}
		rt := make([]int64, 0)
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
