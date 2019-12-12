package strings

import (
	"context"
	"fmt"
	"strings"

	"github.com/Focinfi/go-pipeline"
	"github.com/Focinfi/misa/handlerbuilders/interpreters"
	"github.com/Focinfi/misa/handlerbuilders/utils"
)

const (
	tengoScriptTmplObject = `found := import("json").decode(data).%s`
	tengoScriptTmplArray  = `found := import("json").decode(data)%s`
)

type FinderJSONAttr struct {
	AttrPath         string `json:"attr_path"`
	interpreterTengo pipeline.Handler
}

func NewFinderJSONAttr(attrPath string) (*FinderJSONAttr, error) {
	interpreterBuilder, ok := interpreters.GetHandlerBuilderOK("tengo")
	if !ok {
		return nil, fmt.Errorf("interpreter tengo is unsupported")
	}

	scriptTmpl := tengoScriptTmplObject
	if strings.HasPrefix(attrPath, "[") {
		scriptTmpl = tengoScriptTmplArray
	}
	meta := interpreters.Meta{
		Script: fmt.Sprintf(scriptTmpl, attrPath),
		InitVarMap: map[string]interface{}{
			"data": "",
		},
		RtVarName: "found",
	}

	interpreter := interpreterBuilder.Build(meta.ToMap())
	return &FinderJSONAttr{
		AttrPath:         attrPath,
		interpreterTengo: interpreter,
	}, nil
}

func BuildFinderJSONAttr(conf map[string]interface{}) pipeline.Handler {
	finder, err := NewFinderJSONAttr(fmt.Sprint(conf["attr_path"]))
	if err != nil {
		panic(err)
	}
	return finder
}

func (finder FinderJSONAttr) Handle(ctx context.Context, reqRes *pipeline.HandleRes) (respRes *pipeline.HandleRes, err error) {
	inRes := &pipeline.HandleRes{}
	var inData string
	if reqRes != nil {
		inRes, err = reqRes.Copy()
		if err != nil {
			return nil, err
		}
		if reqRes.Data != nil {
			inData, err = utils.AnyTypeToString(reqRes.Data)
			if err != nil {
				return nil, err
			}
		}
	}

	inRes.Data = map[string]interface{}{
		"data": inData,
	}
	return finder.interpreterTengo.Handle(ctx, inRes)
}
