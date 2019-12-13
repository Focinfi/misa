package iterators

import (
	"context"
	"fmt"

	"github.com/Focinfi/go-pipeline"
	"github.com/Focinfi/misa/handlerbuilders/interpreters"
)

const detectionScriptTemplateTengo = `
selected := import("enum").filter(arr, func(key, value) { return %s })
`

type Selector struct {
	InterpreterName string           `json:"interpreter_name"`
	Script          string           `json:"script"`
	interpreter     pipeline.Handler `json:"-"`
}

func NewSelector(interpreterName, detectionScript string) (*Selector, error) {
	builder, err := interpreters.GetHandlerBuilder(interpreterName)
	if err != nil {
		return nil, err
	}
	meta := interpreters.Meta{
		Script: buildDetectionScript(interpreterName, detectionScript),
		InitVarMap: map[string]interface{}{
			"arr": []interface{}{},
		},
		RtVarName: "selected",
	}
	interpreter := builder.Build(meta.ToMap())
	return &Selector{
		InterpreterName: interpreterName,
		Script:          detectionScript,
		interpreter:     interpreter,
	}, nil
}

func buildDetectionScript(interpreterName, detectionScript string) string {
	switch interpreterName {
	case "tengo":
		return fmt.Sprintf(detectionScriptTemplateTengo, detectionScript)
	}
	return ""
}

func (selector Selector) Handle(ctx context.Context, reqRes *pipeline.HandleRes) (respRes *pipeline.HandleRes, err error) {
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
