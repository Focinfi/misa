package io

import (
	"context"
	"fmt"
	"io/ioutil"

	"github.com/Focinfi/misa/handlerbuilders/confparam"

	"github.com/Focinfi/go-pipeline"
	"github.com/Focinfi/misa/handlerbuilders/utils"
)

var writerFileConfParams = map[string]confparam.ConfParam{}

func init() {
	params, err := confparam.GetConfParams(WriterFile{})
	if err != nil {
		panic(err)
	}
	writerFileConfParams = params
}

type WriterFile struct {
	Path string `json:"path"`
}

func (file WriterFile) Build() (pipeline.Handler, error) {
	return WriterFile{Path: file.Path}, nil
}

func (file WriterFile) ConfParams() map[string]confparam.ConfParam {
	return writerFileConfParams
}

func (file *WriterFile) InitByConf(conf map[string]interface{}) error {
	return utils.JSONUnmarshalWithMap(conf, file)
}

func (file WriterFile) Handle(ctx context.Context, reqRes *pipeline.HandleRes) (respRes *pipeline.HandleRes, err error) {
	respRes = &pipeline.HandleRes{}
	if reqRes != nil {
		respRes, err = reqRes.Copy()
		if err != nil {
			return nil, err
		}

		if reqRes.Data != nil {
			content, err := utils.AnyTypeToString(reqRes.Data)
			if err != nil {
				return nil, err
			}
			if err := ioutil.WriteFile(file.Path, []byte(content), 0644); err != nil {
				return nil, fmt.Errorf("write into file err: %v", err)
			}
		}
	}
	return respRes, nil
}
