package confparam

import (
	"reflect"
	"testing"
)

type TestObj struct {
	Foo string                 `json:"foo_key" desc:"foo foo foo" validate:"required"`
	Bar map[string]interface{} `json:"bar" desc:"bar map" validate:"required"`
	prv int                    `json:"prv"`
}

func TestGetConfParams(t *testing.T) {
	type args struct {
		obj interface{}
	}
	tests := []struct {
		name    string
		args    args
		want    map[string]ConfParam
		wantErr bool
	}{
		{
			name: "normal",
			args: args{
				obj: TestObj{},
			},
			want: map[string]ConfParam{
				"foo_key": {
					Name:       "Foo",
					Type:       "string",
					Desc:       "foo foo foo",
					Validation: "required",
				},
				"bar": {
					Name:       "Bar",
					Type:       "map[string]interface {}",
					Desc:       "bar map",
					Validation: "required",
				},
			},
		},
		{
			name: "duplicated json tag",
			args: args{
				obj: struct {
					Foo string `json:"foo"`
					Bar int    `json:"foo"`
				}{},
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetConfParams(tt.args.obj)
			if err != nil {
				t.Log("err:", err)
			}
			if (err != nil) != tt.wantErr {
				t.Errorf("GetConfParams() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetConfParams() = %v, want %v", got, tt.want)
			}
		})
	}
}
