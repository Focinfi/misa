package io

import (
	"context"
	"io/ioutil"
	"os"

	"github.com/Focinfi/go-pipeline"
	"github.com/Focinfi/misa/handlerbuilders/utils"
)

type WriterFile struct {
	Path string `json:"path"`
}

func BuildFile(conf map[string]interface{}) pipeline.Handler {
	return &WriterFile{
		Path: conf["path"].(string),
	}
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
			if err := ioutil.WriteFile(file.Path, []byte(content), os.ModePerm); err != nil {
				return nil, err
			}
		}
	}
	return respRes, nil
}
