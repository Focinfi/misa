package strings

import (
	"context"
	"strings"

	"github.com/Focinfi/misa/builders/utils"

	"github.com/Focinfi/go-pipeline"
)

type Replacer struct {
	Old string `json:"old" validate:"required"`
	New string `json:"new" validate:"-"`
	N   int    `json:"n" validate:"gte=-1"`
}

func (replacer Replacer) Build() (pipeline.Handler, error) {
	return &Replacer{
		Old: replacer.Old,
		New: replacer.New,
		N:   replacer.N,
	}, nil
}

func (replacer *Replacer) Handle(ctx context.Context, reqRes *pipeline.HandleRes) (respRes *pipeline.HandleRes, err error) {
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

			respRes.Data = strings.Replace(data, replacer.Old, replacer.New, replacer.N)
		}
	}

	respRes.Status = pipeline.HandleStatusOK
	return respRes, nil
}
