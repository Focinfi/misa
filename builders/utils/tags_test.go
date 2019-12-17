package utils

import (
	"reflect"
	"testing"
)

func TestGetTypeDefinition(t *testing.T) {
	type args struct {
		obj  interface{}
		tags []string
	}
	tests := []struct {
		name string
		args args
		want []FieldDefinition
	}{
		{
			name: "normal",
			args: args{
				obj: struct {
					A string `json:"a" desc:"a field" validate:"-"`
					B int    `json:"b"`
					c map[string]interface{}
				}{},
				tags: []string{"json", "desc", "validate"},
			},
			want: []FieldDefinition{
				{
					Name: "A",
					Type: "string",
					Tags: map[string]string{
						"json":     "a",
						"desc":     "a field",
						"validate": "-",
					},
				},
				{
					Name: "B",
					Type: "int",
					Tags: map[string]string{
						"json":     "b",
						"desc":     "",
						"validate": "",
					},
				},
				{
					Name: "c",
					Type: "map[string]interface {}",
					Tags: map[string]string{
						"json":     "",
						"desc":     "",
						"validate": "",
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetTypeDefinitions(tt.args.obj, tt.args.tags); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetTypeDefinitions() = %v, want %v", got, tt.want)
			}
		})
	}
}
