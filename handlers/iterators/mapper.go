package iterators

import (
	"context"
	"errors"
	"fmt"

	"github.com/Focinfi/go-pipeline"
	"github.com/Focinfi/misa/handlers/interpreters"
)

const mapScriptTemplateTengo = `
transition := func(i, v) {
	return %s
}
mapped := []
for i, v in arr {
	mapped = append(mapped, transition(i, v))
}
`

type Mapper struct {
	InterpreterName string           `json:"interpreter_name"`
	MapScript       string           `json:"map_script"`
	interpreter     pipeline.Handler `json:"-"`
}

func NewMapper(interpreterName, mapScript string) (*Mapper, error) {
	builder, ok := interpreters.GetHandlerBuilderOK(interpreterName)
	if !ok {
		return nil, errors.New("unsupported interpreter")
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
		MapScript:       mapScript,
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
			inArr = reqRes.Data
		}
	}
	inRes.Data = map[string]interface{}{
		"arr": inArr,
	}
	return selector.interpreter.Handle(ctx, inRes)
}
