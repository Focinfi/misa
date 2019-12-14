package utils

import (
	"reflect"
)

type FieldDefinition struct {
	Name string
	Type string
	Tags map[string]string
}

func GetTypeDefinitions(obj interface{}, tags []string) []FieldDefinition {
	t := reflect.TypeOf(obj)
	fields := make([]FieldDefinition, 0)
	for i := 0; i < t.NumField(); i++ {
		f := t.Field(i)

		t := make(map[string]string, len(tags))
		for _, tag := range tags {
			t[tag] = f.Tag.Get(tag)
		}

		fields = append(fields, FieldDefinition{
			Name: f.Name,
			Type: f.Type.String(),
			Tags: t,
		})
	}
	return fields
}
