package iterators

import (
	"context"
	"fmt"

	"github.com/Focinfi/go-pipeline"
	"github.com/Focinfi/misa/builders/interpreters"
)

const chunkScriptTemplateTengo = `
chunked := import("enum").chunk(arr, %s)
`

type Chunker struct {
	Conf
	interpreter pipeline.Handler `json:"-"`
}

func NewChunker(conf Conf) (*Chunker, error) {
	iteratorConf := interpreters.Conf{
		Type:   conf.InterpreterName,
		Script: buildChunkScript(conf.InterpreterName, conf.Script),
		InitVarMap: map[string]interface{}{
			"arr": []interface{}{},
		},
		RtVarName: "chunked",
	}
	interpreter, err := iteratorConf.Build()
	if err != nil {
		return nil, err
	}
	return &Chunker{
		Conf:        conf,
		interpreter: interpreter,
	}, nil
}

func buildChunkScript(interpreterName, chunkScript string) string {
	switch interpreterName {
	case "tengo":
		return fmt.Sprintf(chunkScriptTemplateTengo, chunkScript)
	}
	return ""
}

func (selector Chunker) Handle(ctx context.Context, reqRes *pipeline.HandleRes) (respRes *pipeline.HandleRes, err error) {
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
