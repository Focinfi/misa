package parsers

import (
	"context"
	"errors"
	"fmt"
	"strconv"
	"time"

	"github.com/Focinfi/go-pipeline"
)

var DefaultUnixParser = UnixParser{}

type UnixParser struct{}

func (UnixParser) Handle(ctx context.Context, reqRes *pipeline.HandleRes) (respRes *pipeline.HandleRes, err error) {
	if reqRes != nil && reqRes.Data != nil {
		respRes, err = reqRes.Copy()
		if err != nil {
			return nil, err
		}

		i, err := strconv.ParseInt(fmt.Sprint(reqRes.Data), 10, 64)
		if err != nil {
			return nil, fmt.Errorf("parse request data into int64 failed, err: %v", err)
		}

		respRes.Status = pipeline.HandleStatusOK
		respRes.Data = time.Unix(i, 0)
		return respRes, nil
	}

	return nil, errors.New("empty request")
}
