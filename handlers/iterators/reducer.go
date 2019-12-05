package iterators

import (
	"context"
	"errors"
	"fmt"

	"github.com/Focinfi/go-pipeline"
	"github.com/Focinfi/misa/handlers/interpreters"
	"github.com/Focinfi/misa/handlers/utils"
)

const reduceScriptTemplateTengo = `
reduced := ""
import("enum").each(arr, func(key, value) { reduced = %s })
`

type Reducer struct {
	InterpreterName string           `json:"interpreter_name"`
	MapScript       string           `json:"map_script"`
	interpreter     pipeline.Handler `json:"-"`
}

func NewReducer(interpreterName, reduceScript string) (*Reducer, error) {
	builder, ok := interpreters.GetHandlerBuilderOK(interpreterName)
	if !ok {
		return nil, errors.New("unsupported interpreter")
	}
	meta := interpreters.Meta{
		Script: buildReduceScript(interpreterName, reduceScript),
		InitVarMap: map[string]interface{}{
			"arr": []interface{}{},
		},
		RtVarName: "reduced",
	}
	interpreter := builder.Build(meta.ToMap())
	return &Reducer{
		InterpreterName: interpreterName,
		MapScript:       reduceScript,
		interpreter:     interpreter,
	}, nil
}

func buildReduceScript(interpreterName, reduceScript string) string {
	switch interpreterName {
	case "tengo":
		return fmt.Sprintf(reduceScriptTemplateTengo, reduceScript)
	}
	return ""
}

func (selector Reducer) Handle(ctx context.Context, reqRes *pipeline.HandleRes) (respRes *pipeline.HandleRes, err error) {
	inRes := &pipeline.HandleRes{}
	var inArr interface{}
	if reqRes != nil {
		inRes, err = reqRes.Copy()
		if err != nil {
			return nil, err
		}

		if reqRes.Data != nil {
			items, err := utils.AynTypeToSlice(reqRes.Data)
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
