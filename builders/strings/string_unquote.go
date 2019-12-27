package strings

import (
	"context"
	"strconv"

	"github.com/Focinfi/go-pipeline"
	"github.com/Focinfi/misa/builders/utils"
)

var DefaultUnquote = Unquote{}

type Unquote struct{}

func (Unquote) Handle(ctx context.Context, reqRes *pipeline.HandleRes) (respRes *pipeline.HandleRes, err error) {
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
			quoted := ""
			for i := 0; i < len(data); i++ {
				s := string(data[i])
				if data[i] == '\\' && i <= len(data)-6 && data[i+1] == 'u' {
					s, err = strconv.Unquote(`"` + data[i:i+6] + `"`)
					if err != nil {
						return nil, err
					}
					i += 5
				}
				quoted += s
			}
			respRes.Data = quoted
		}
	}

	respRes.Status = pipeline.HandleStatusOK
	return respRes, nil
}
