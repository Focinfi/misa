package strings

import (
	"context"
	"strings"

	"github.com/Focinfi/go-pipeline"
	"github.com/Focinfi/misa/handlers/utils"
)

type Splitter struct {
	Separator string `json:"separator"`
}

func BuildSplitter(conf map[string]interface{}) pipeline.Handler {
	return Splitter{Separator: conf["separator"].(string)}
}

func (str Splitter) Handle(ctx context.Context, reqRes *pipeline.HandleRes) (respRes *pipeline.HandleRes, err error) {
	if reqRes == nil || reqRes.Data == nil {
		return nil, nil
	}
	inRes, err := reqRes.Copy()
	if err != nil {
		return nil, err
	}
	inData, err := utils.AnyTypeToString(inRes.Data)
	if err != nil {
		return nil, err
	}

	return &pipeline.HandleRes{
		Status: pipeline.HandleStatusOK,
		Data:   strings.Split(inData, str.Separator),
		Meta:   inRes.Meta,
	}, nil
}
