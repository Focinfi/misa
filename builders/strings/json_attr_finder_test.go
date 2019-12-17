package strings

import (
	"context"
	"reflect"
	"testing"

	"github.com/Focinfi/go-pipeline"
)

func TestFinderJSONAttr_Handle(t *testing.T) {
	type fields struct {
		AttrPath string
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
			name: "from json object string",
			fields: fields{
				AttrPath: "a.b[0].c",
			},
			args: args{
				ctx: context.Background(),
				reqRes: &pipeline.HandleRes{
					Data: `{"a": {"b": [{"c": 2}]}}`,
				},
			},
			wantRespRes: &pipeline.HandleRes{
				Status: pipeline.HandleStatusOK,
				Data:   float64(2),
			},
			wantErr: false,
		},
		{
			name: "from json array string",
			fields: fields{
				AttrPath: "[0].a.b[0].c",
			},
			args: args{
				ctx: context.Background(),
				reqRes: &pipeline.HandleRes{
					Data: `[{"a": {"b": [{"c": 2}]}}]`,
				},
			},
			wantRespRes: &pipeline.HandleRes{
				Status: pipeline.HandleStatusOK,
				Data:   float64(2),
			},
			wantErr: false,
		},
		{
			name: "not found from json object string",
			fields: fields{
				AttrPath: "a.b[0].d",
			},
			args: args{
				ctx: context.Background(),
				reqRes: &pipeline.HandleRes{
					Data: `{"a": {"b": [{"c": 2}]}}`,
				},
			},
			wantRespRes: &pipeline.HandleRes{
				Status: pipeline.HandleStatusOK,
				Data:   nil,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			finder, err := NewFinderJSONAttr(tt.fields.AttrPath)
			if err != nil {
				t.Fatal(err)
			}
			gotInRes, err := finder.Handle(tt.args.ctx, tt.args.reqRes)
			if (err != nil) != tt.wantErr {
				t.Errorf("FinderJSONAttr.Handle() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotInRes, tt.wantRespRes) {
				t.Errorf("FinderJSONAttr.Handle() = %v, want %v", gotInRes, tt.wantRespRes)
			}
		})
	}
}
