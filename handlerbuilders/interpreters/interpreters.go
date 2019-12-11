package interpreters

import "github.com/Focinfi/go-pipeline"

var all = map[string]pipeline.HandlerBuilder{
	"tengo": pipeline.HandlerBuilderFunc(func(conf map[string]interface{}) pipeline.Handler {
		return &Tengo{
			Meta: Meta{
				Script:     conf["script"].(string),
				InitVarMap: conf["init_var_map"].(map[string]interface{}),
				RtVarName:  conf["rt_var_name"].(string),
			},
		}
	}),
}

func GetHandlerBuilderOK(name string) (pipeline.HandlerBuilder, bool) {
	hb, ok := all[name]
	return hb, ok
}
