package strings

import (
	"context"
	"fmt"
	"strings"

	"github.com/Focinfi/go-pipeline"
	"github.com/Focinfi/misa/handlers/utils"
)

type Join struct {
	Separator string `json:"separator"`
}

func BuildJoin(conf map[string]interface{}) pipeline.Handler {
	return Join{Separator: conf["separator"].(string)}
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
		return nil, err
	}
	items := make([]string, 0, len(inData))
	for _, data := range inData {
		items = append(items, fmt.Sprint(data))
	}

	return &pipeline.HandleRes{
		Status: pipeline.HandleStatusOK,
		Data:   strings.Join(items, str.Separator),
		Meta:   inRes.Meta,
	}, nil
}
