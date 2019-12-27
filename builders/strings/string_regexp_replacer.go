package strings

import (
	"context"
	"regexp"

	"github.com/Focinfi/go-pipeline"
	"github.com/Focinfi/misa/builders/utils"
)

type RegexpAllReplacer struct {
	Expr        string `json:"expr" validate:"required"`
	Replacement string `json:"replacement" validate:"required"`
	regexp      *regexp.Regexp
}

func (r RegexpAllReplacer) Build() (pipeline.Handler, error) {
	rp, err := regexp.Compile(r.Expr)
	if err != nil {
		return nil, err
	}

	return &RegexpAllReplacer{
		Expr:        r.Expr,
		Replacement: r.Replacement,
		regexp:      rp,
	}, nil
}

func (r *RegexpAllReplacer) Handle(ctx context.Context, reqRes *pipeline.HandleRes) (respRes *pipeline.HandleRes, err error) {
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

			respRes.Data = r.regexp.ReplaceAllString(data, r.Replacement)
		}
	}

	respRes.Status = pipeline.HandleStatusOK
	return respRes, nil
}
