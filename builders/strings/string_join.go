package strings

import (
	"context"
	"fmt"
	"strings"

	"github.com/Focinfi/go-pipeline"
	"github.com/Focinfi/misa/builders/utils"
)

type Join struct {
	SeparatorConf
}

func (str Join) Handle(ctx context.Context, reqRes *pipeline.HandleRes) (respRes *pipeline.HandleRes, err error) {
	if reqRes == nil || reqRes.Data == nil {
		return &pipeline.HandleRes{
			Status: pipeline.HandleStatusOK,
			Data:   "",
		}, nil
	}
	inRes, err := reqRes.Copy()
	if err != nil {
		return nil, err
	}
	inData, err := utils.AynTypeToSlice(reqRes.Data)
	if err != nil {
		return nil, fmt.Errorf("request data type wrong: %v", err)
	}
	items := make([]string, 0, len(inData))
	for _, data := range inData {
		items = append(items, fmt.Sprint(data))
	}

	inRes.Status = pipeline.HandleStatusOK
	inRes.Data = strings.Join(items, str.Separator)
	return inRes, nil
}
