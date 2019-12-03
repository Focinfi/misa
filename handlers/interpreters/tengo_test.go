package interpreters

import (
	"context"
	"reflect"
	"testing"

	"github.com/Focinfi/go-pipeline"
	"github.com/d5/tengo/script"
)

func TestTengo_Handle(t *testing.T) {
	type fields struct {
		Script     string
		InitVarMap map[string]interface{}
		ReturnVar  string
		compiled   *script.Compiled
	}
	type args struct {
		ctx    context.Context
		reqRes *pipeline.HandleRes
	}
	tests := []struct {
		name        string
		fields      fields
		args        args
		wantRespRes *pipeline.HandleRes
		wantErr     bool
	}{
		{
			name: "normal",
			fields: fields{
				Script: "a := b + 1",
				InitVarMap: map[string]interface{}{
					"b": 0,
				},
				ReturnVar: "a",
			},
			args: args{
				ctx: context.Background(),
				reqRes: &pipeline.HandleRes{
					Data: map[string]interface{}{
						"b": 2,
					},
				},
			},
			wantRespRes: &pipeline.HandleRes{
				Status: pipeline.HandleStatusOK,
				Data:   int64(3),
			},
			wantErr: false,
		},
		{
			name: "func",
			fields: fields{
				Script: `
     					 text := import("text")
						 prefixed_by := func(s, p) {
							return text.has_prefix(s, p)		
						 }
 						 rt := prefixed_by(str, prefix)`,
				InitVarMap: map[string]interface{}{
					"str":    "foo",
					"prefix": "bar",
				},
				ReturnVar: "rt",
			},
			args: args{
				ctx: context.Background(),
				reqRes: &pipeline.HandleRes{
					Data: map[string]interface{}{
						"str":    "foobar",
						"prefix": "foo",
					},
				},
			},
			wantRespRes: &pipeline.HandleRes{
				Status: pipeline.HandleStatusOK,
				Data:   true,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tengo := &Tengo{
				Script:     tt.fields.Script,
				RtVarName:  tt.fields.ReturnVar,
				InitVarMap: tt.fields.InitVarMap,
				compiled:   tt.fields.compiled,
			}
			gotRespRes, err := tengo.Handle(tt.args.ctx, tt.args.reqRes)
			if (err != nil) != tt.wantErr {
				t.Errorf("Tengo.Handle() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotRespRes, tt.wantRespRes) {
				t.Errorf("Tengo.Handle() = %v, want %v", gotRespRes, tt.wantRespRes)
			}
		})
	}
}
