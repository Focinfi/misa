package interpreters

import (
	"context"
	"errors"

	"github.com/Focinfi/go-pipeline"
	"github.com/d5/tengo/script"
	"github.com/d5/tengo/stdlib"
)

type Tengo struct {
	Meta
	compiled *script.Compiled `json:"-"`
}

func (tengo *Tengo) Handle(ctx context.Context, reqRes *pipeline.HandleRes) (respRes *pipeline.HandleRes, err error) {
	compiled, err := tengo.compile()
	if err != nil {
		return nil, err
	}
	compiled = compiled.Clone()
	if reqRes != nil && reqRes.Data != nil {
		if err := tengo.setVar(compiled, reqRes.Data); err != nil {
			return nil, err
		}
	}
	if err := compiled.RunContext(ctx); err != nil {
		return nil, err
	}

	var rtData interface{}
	if tengo.RtVarName != "" {
		if rtVal := compiled.Get(tengo.RtVarName); rtVal != nil {
			rtData = rtVal.Value()
		}
	}
	respRes = &pipeline.HandleRes{}
	if reqRes != nil {
		respRes, err = reqRes.Copy()
		if err != nil {
			return nil, err
		}
	}
	respRes.Status = pipeline.HandleStatusOK
	respRes.Data = rtData
	return respRes, nil
}

func (tengo *Tengo) setVar(compiled *script.Compiled, args interface{}) error {
	argsMap, ok := args.(map[string]interface{})
	if !ok {
		return errors.New("args is not a map[string]interface{}")
	}
	for name, val := range argsMap {
		if err := compiled.Set(name, val); err != nil {
			return err
		}
	}
	return nil
}

func (tengo *Tengo) compile() (*script.Compiled, error) {
	if tengo.compiled != nil {
		return tengo.compiled, nil
	}
	s := script.New([]byte(tengo.Script))
	s.SetImports(stdlib.GetModuleMap(stdlib.AllModuleNames()...))
	for varName, val := range tengo.InitVarMap {
		if err := s.Add(varName, val); err != nil {
			return nil, err
		}
	}
	c, err := s.Compile()
	if err != nil {
		return nil, err
	}
	tengo.compiled = c
	return c, nil
}
