package confparam

import (
	"fmt"
	"reflect"
	"regexp"

	"github.com/Focinfi/misa/handlerbuilders/utils"
)

var publicFieldRegexp = regexp.MustCompile("^[A-Z]")

type ConfParam struct {
	Name       string `json:"name"`
	Type       string `json:"type"`
	Desc       string `json:"desc"`
	Validation string `json:"validation"`
}

func GetConfParams(obj interface{}) (map[string]ConfParam, error) {
	params := make(map[string]ConfParam)
	fields := utils.GetTypeDefinitions(obj, []string{"json", "desc", "validate"})
	for _, field := range fields {
		name := field.Tags["json"]
		if name == "" || name == "-" || !publicFieldRegexp.MatchString(field.Name) {
			continue
		}
		f, ok := params[name]
		if ok {
			objType := reflect.TypeOf(obj)
			return nil, fmt.Errorf("duplicated json tag(%#v) in %#v and %#v of type %#v",
				name, f.Name, field.Name, objType.PkgPath()+objType.Name())
		}
		params[name] = ConfParam{
			Name:       field.Name,
			Type:       field.Type,
			Desc:       field.Tags["desc"],
			Validation: field.Tags["validate"],
		}
	}
	return params, nil
}
