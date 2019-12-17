package io

import (
	"context"
	"fmt"

	"github.com/Focinfi/go-pipeline"
	"github.com/go-vgo/robotgo/clipboard"
)

var (
	DefaultReaderClipboard = ReaderClipboard{}
	DefaultWriterClipboard = WriterClipboard{}
)

type ReaderClipboard struct{}

func (cb ReaderClipboard) Handle(ctx context.Context, reqRes *pipeline.HandleRes) (respRes *pipeline.HandleRes, err error) {
	cbData, err := clipboard.ReadAll()
	if err != nil {
		return nil, fmt.Errorf("clipboard read err: %v", err)
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

type WriterClipboard struct{}

func (cb WriterClipboard) Handle(ctx context.Context, reqRes *pipeline.HandleRes) (respRes *pipeline.HandleRes, err error) {
	if reqRes == nil {
		return nil, nil
	}

	respRes, err = reqRes.Copy()
	if err != nil {
		return nil, err
	}
	if err := clipboard.WriteAll(fmt.Sprint(reqRes.Data)); err != nil {
		return nil, fmt.Errorf("clipboard write err: %v", err)
	}
	respRes.Status = pipeline.HandleStatusOK
	return respRes, nil
}
