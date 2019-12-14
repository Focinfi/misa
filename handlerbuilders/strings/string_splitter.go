package strings

import (
	"context"
	"strings"

	"github.com/Focinfi/go-pipeline"
	"github.com/Focinfi/misa/handlerbuilders/utils"
)

type Split struct {
	SeparatorConf
}

func (str Split) Handle(ctx context.Context, reqRes *pipeline.HandleRes) (respRes *pipeline.HandleRes, err error) {
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

	inRes.Status = pipeline.HandleStatusOK
	inRes.Data = strings.Split(inData, str.Separator)
	return inRes, nil
}
