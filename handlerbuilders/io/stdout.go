package io

import (
	"context"

	"github.com/Focinfi/go-pipeline"
)

var DefaultWriterStdOut = WriterStdOut{}

type WriterStdOut struct {
	//FormatName string `json:"formatName"`
}

func (stdout WriterStdOut) Handle(ctx context.Context, reqRes *pipeline.HandleRes) (respRes *pipeline.HandleRes, err error) {
	respRes = &pipeline.HandleRes{}
	if reqRes != nil {
		respRes, err = reqRes.Copy()
		if err != nil {
			return nil, err
		}
	}
	respRes.Status = pipeline.HandleStatusOK
	return
}
