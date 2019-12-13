package iterators

import (
	"context"
	"fmt"

	"github.com/Focinfi/go-pipeline"
	"github.com/Focinfi/misa/handlerbuilders/interpreters"
)

const mapScriptTemplateTengo = `
mapped := import("enum").map(arr, func(key, value){ return %s })
`

type Mapper struct {
	InterpreterName string           `json:"interpreter_name"`
	Script          string           `json:"script"`
	interpreter     pipeline.Handler `json:"-"`
}

func NewMapper(interpreterName, mapScript string) (*Mapper, error) {
	builder, err := interpreters.GetHandlerBuilder(interpreterName)
	if err != nil {
		return nil, err
	}
	meta := interpreters.Meta{
		Script: buildMapScript(interpreterName, mapScript),
		InitVarMap: map[string]interface{}{
			"arr": []interface{}{},
		},
		RtVarName: "mapped",
	}
	interpreter := builder.Build(meta.ToMap())
	return &Mapper{
		InterpreterName: interpreterName,
		Script:          mapScript,
		interpreter:     interpreter,
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
