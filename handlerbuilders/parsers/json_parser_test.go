package parsers

import (
	"context"
	"reflect"
	"testing"

	"github.com/Focinfi/go-pipeline"
)

func TestParserJSON_Handle(t *testing.T) {
	type args struct {
		ctx    context.Context
		reqRes *pipeline.HandleRes
	}
	tests := []struct {
		name        string
		p           ParserJSON
		args        args
		wantRespRes *pipeline.HandleRes
		wantErr     bool
	}{
		{
			name: "object",
			p:    ParserJSON{},
			args: args{
				ctx: context.Background(),
				reqRes: &pipeline.HandleRes{
					Data: `{"name": "foo", "age": 1, "cars": ["bar"]}`,
				},
			},
			wantRespRes: &pipeline.HandleRes{
				Status: pipeline.HandleStatusOK,
				Data: map[string]interface{}{
					"name": "foo",
					"age":  float64(1),
					"cars": []interface{}{"bar"},
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := ParserJSON{}
			gotRespRes, err := p.Handle(tt.args.ctx, tt.args.reqRes)
			if (err != nil) != tt.wantErr {
				t.Errorf("ParserJSON.Handle() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotRespRes, tt.wantRespRes) {
				t.Errorf("ParserJSON.Handle() = %#v, want %#v", gotRespRes, tt.wantRespRes)
			}
		})
	}
}
