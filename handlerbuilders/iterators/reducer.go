package iterators

import (
	"context"
	"fmt"

	"github.com/Focinfi/go-pipeline"
	"github.com/Focinfi/misa/handlerbuilders/interpreters"
)

const reduceScriptTemplateTengo = `
reduced := ""
import("enum").each(arr, func(key, value) { reduced = %s })
`

type Reducer struct {
	Conf
	interpreter pipeline.Handler `json:"-"`
}

func NewReducer(conf Conf) (*Reducer, error) {
	iteratorConf := interpreters.Conf{
		Type:   conf.InterpreterName,
		Script: buildReduceScript(conf.InterpreterName, conf.Script),
		InitVarMap: map[string]interface{}{
			"arr": []interface{}{},
		},
		RtVarName: "reduced",
	}
	interpreter, err := iteratorConf.Build()
	if err != nil {
		return nil, err
	}
	return &Reducer{
		Conf:        conf,
		interpreter: interpreter,
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
