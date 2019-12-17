package parsers

import (
	"context"
	"encoding/json"

	"github.com/Focinfi/go-pipeline"
	"github.com/Focinfi/misa/builders/utils"
)

var DefaultParserJSON = ParserJSON{}

type ParserJSON struct{}

func (ParserJSON) Handle(ctx context.Context, reqRes *pipeline.HandleRes) (respRes *pipeline.HandleRes, err error) {
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

			var j interface{}
			if err := json.Unmarshal([]byte(data), &j); err != nil {
				return nil, err
			}
			respRes.Data = j
		}
	}
	respRes.Status = pipeline.HandleStatusOK
	return respRes, nil
}
