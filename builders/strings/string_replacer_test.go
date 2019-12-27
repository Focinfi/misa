package strings

import (
	"context"
	"reflect"
	"testing"

	"github.com/Focinfi/go-pipeline"
)

func TestReplacer_Handle(t *testing.T) {
	type fields struct {
		Old string
		New string
		N   int
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
				Old: "\u003c",
				New: ">",
				N:   -1,
			},
			args: args{
				reqRes: &pipeline.HandleRes{
					Data: "uabc\u003c",
				},
			},
			wantRespRes: &pipeline.HandleRes{
				Status: pipeline.HandleStatusOK,
				Data:   "uabc>",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			replacer := &Replacer{
				Old: tt.fields.Old,
				New: tt.fields.New,
				N:   tt.fields.N,
			}
			gotRespRes, err := replacer.Handle(tt.args.ctx, tt.args.reqRes)
			if (err != nil) != tt.wantErr {
				t.Errorf("Replacer.Handle() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotRespRes, tt.wantRespRes) {
				t.Errorf("Replacer.Handle() = %v, want %v", gotRespRes, tt.wantRespRes)
			}
		})
	}
}
