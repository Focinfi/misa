package iterators

import (
	"context"
	"fmt"

	"github.com/Focinfi/go-pipeline"
	"github.com/Focinfi/misa/builders/interpreters"
)

const mapScriptTemplateTengo = `
mapped := import("enum").map(arr, func(key, value){ return %s })
`

type Mapper struct {
	Conf
	interpreter pipeline.Handler `json:"-"`
}

func NewMapper(conf Conf) (*Mapper, error) {
	iteratorConf := interpreters.Conf{
		Type:   conf.InterpreterName,
		Script: buildMapScript(conf.InterpreterName, conf.Script),
		InitVarMap: map[string]interface{}{
			"arr": []interface{}{},
		},
		RtVarName: "mapped",
	}
	interpreter, err := iteratorConf.Build()
	if err != nil {
		return nil, err
	}
	return &Mapper{
		Conf:        conf,
		interpreter: interpreter,
	}, nil
}

func buildMapScript(interpreterName, mapScript string) string {
	switch interpreterName {
	case "tengo":
		return fmt.Sprintf(mapScriptTemplateTengo, mapScript)
	}
	return ""
}

func (selector Mapper) Handle(ctx context.Context, reqRes *pipeline.HandleRes) (respRes *pipeline.HandleRes, err error) {
	inRes := &pipeline.HandleRes{}
	var inArr interface{}
	if reqRes != nil {
		inRes, err = reqRes.Copy()
		if err != nil {
			return nil, err
		}

		if reqRes.Data != nil {
			items, err := InterfaceToSlice(reqRes.Data)
			if err != nil {
				return nil, err
			}
			inArr = items
		}
	}

	inRes.Data = map[string]interface{}{
		"arr": inArr,
	}
	return selector.interpreter.Handle(ctx, inRes)
}
