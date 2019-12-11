package io

import (
	"context"
	"fmt"

	"github.com/Focinfi/go-pipeline"
)

var DefaultStdOut = StdOut{}

type StdOut struct {
	//FormatName string `json:"formatName"`
}

func (stdout StdOut) Handle(ctx context.Context, reqRes *pipeline.HandleRes) (respRes *pipeline.HandleRes, err error) {
	respRes = &pipeline.HandleRes{}
	if reqRes != nil {
		fmt.Println(reqRes.Data)
		respRes, err = reqRes.Copy()
		if err != nil {
			return nil, err
		}
	}
	respRes.Status = pipeline.HandleStatusOK
	return
}
