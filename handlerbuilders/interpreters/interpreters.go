package interpreters

import (
	"fmt"

	"github.com/Focinfi/go-pipeline"
)

var all = map[string]pipeline.HandlerBuilder{
	"tengo": pipeline.HandlerBuilderFunc(func(conf map[string]interface{}) (pipeline.Handler, error) {
		return &Tengo{
			Meta: Meta{
				Script:     conf["script"].(string),
				InitVarMap: conf["init_var_map"].(map[string]interface{}),
				RtVarName:  conf["rt_var_name"].(string),
			},
		}, nil
	}),
}

func GetHandlerBuilderOK(name string) (pipeline.HandlerBuilder, bool) {
	hb, ok := all[name]
	return hb, ok
}

func GetHandlerBuilder(name string) (pipeline.HandlerBuilder, error) {
	hb, ok := GetHandlerBuilderOK(name)
	if !ok {
		return nil, fmt.Errorf(`interpreter("%v") not found`, name)
	}
	return hb, nil
}
