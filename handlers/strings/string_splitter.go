package strings

import (
	"context"
	"strings"

	"github.com/Focinfi/go-pipeline"
	"github.com/Focinfi/misa/handlers/utils"
)

type String struct {
	Separator string `json:"separator"`
}

func (str String) Handle(ctx context.Context, reqRes *pipeline.HandleRes) (respRes *pipeline.HandleRes, err error) {
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
