package strings

import (
	"context"
	"fmt"
	"strings"

	"github.com/Focinfi/misa/handlerbuilders/confparam"

	"github.com/Focinfi/go-pipeline"
	"github.com/Focinfi/misa/handlerbuilders/interpreters"
	"github.com/Focinfi/misa/handlerbuilders/utils"
)

const (
	tengoScriptTmplObject = `found := import("json").decode(data).%s`
	tengoScriptTmplArray  = `found := import("json").decode(data)%s`
)

var finderJSONAttrParams = make(map[string]confparam.ConfParam)

func init() {
	params, err := confparam.GetConfParams(FinderJSONAttr{})
	if err != nil {
		panic(err)
	}
	finderJSONAttrParams = params
}

type FinderJSONAttr struct {
	AttrPath         string `json:"attr_path" desc:"json value path, e.g. a.b[0].c" validate:"required"`
	interpreterTengo pipeline.Handler
}

func (finder *FinderJSONAttr) Build() (pipeline.Handler, error) {
	return NewFinderJSONAttr(finder.AttrPath)
}

func (finder *FinderJSONAttr) ConfParams() map[string]confparam.ConfParam {
	return finderJSONAttrParams
}

func (finder *FinderJSONAttr) InitByConf(conf map[string]interface{}) error {
	return utils.JSONUnmarshalWithMap(conf, finder)
}

func NewFinderJSONAttr(attrPath string) (*FinderJSONAttr, error) {
	scriptTmpl := tengoScriptTmplObject
	if strings.HasPrefix(attrPath, "[") {
		scriptTmpl = tengoScriptTmplArray
	}
	tengoConf := interpreters.Conf{
		Type:   "tengo",
		Script: fmt.Sprintf(scriptTmpl, attrPath),
		InitVarMap: map[string]interface{}{
			"data": "",
		},
		RtVarName: "found",
	}

	interpreter, err := tengoConf.Build()
	if err != nil {
		return nil, err
	}
	return &FinderJSONAttr{
		AttrPath:         attrPath,
		interpreterTengo: interpreter,
	}, nil
}

func BuildFinderJSONAttr(conf map[string]interface{}) (pipeline.Handler, error) {
	return NewFinderJSONAttr(fmt.Sprint(conf["attr_path"]))
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
