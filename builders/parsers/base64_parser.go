package parsers

import (
	"context"
	"encoding/base64"

	"github.com/Focinfi/go-pipeline"
	"github.com/Focinfi/misa/builders/utils"
)

var DefaultBase64Parser = Base64Parser{}

type Base64Parser struct{}

func (Base64Parser) Handle(ctx context.Context, reqRes *pipeline.HandleRes) (respRes *pipeline.HandleRes, err error) {
	respRes = &pipeline.HandleRes{}
	if reqRes != nil {
		respRes, err = reqRes.Copy()
		if err != nil {
			return nil, err
		}

		if reqRes.Data != nil {
			data, err := utils.AnyTypeToString(reqRes.Data)
			if err != nil {
				return nil, err
			}

			b, err := base64.StdEncoding.DecodeString(data)
			if err != nil {
				return nil, err
			}

			respRes.Data = b
		}
	}
	respRes.Status = pipeline.HandleStatusOK
	return respRes, nil
}
