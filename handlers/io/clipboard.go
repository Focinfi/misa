package io

import (
	"context"
	"fmt"

	"github.com/Focinfi/go-pipeline"
	"github.com/go-vgo/robotgo/clipboard"
)

type ClipboardReader struct{}

func (cb ClipboardReader) Handle(ctx context.Context, reqRes *pipeline.HandleRes) (respRes *pipeline.HandleRes, err error) {
	cbData, err := clipboard.ReadAll()
	if err != nil {
		return nil, err
	}

	respRes = &pipeline.HandleRes{}
	if reqRes != nil {
		respRes, err = reqRes.Copy()
		if err != nil {
			return nil, err
		}
	}

	respRes.Status = pipeline.HandleStatusOK
	respRes.Data = cbData
	return respRes, nil
}

type ClipboardWriter struct{}

func (cb ClipboardWriter) Handle(ctx context.Context, reqRes *pipeline.HandleRes) (respRes *pipeline.HandleRes, err error) {
	if reqRes == nil {
		return nil, nil
	}

	respRes, err = reqRes.Copy()
	if err != nil {
		return nil, err
	}
	if err := clipboard.WriteAll(fmt.Sprint(reqRes.Data)); err != nil {
		return nil, err
	}
	respRes.Status = pipeline.HandleStatusOK
	return respRes, nil
}
