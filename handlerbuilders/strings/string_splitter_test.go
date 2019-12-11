package strings

import (
	"context"
	"reflect"
	"testing"

	"github.com/Focinfi/go-pipeline"
)

func TestString_Handle(t *testing.T) {
	type fields struct {
		Separator string
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
			name: "string input",
			fields: fields{
				Separator: ",",
			},
			args: args{
				ctx: context.Background(),
				reqRes: &pipeline.HandleRes{
					Data: "one,two,three",
				},
			},
			wantRespRes: &pipeline.HandleRes{
				Status: pipeline.HandleStatusOK,
				Data:   []string{"one", "two", "three"},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			str := Splitter{
				Separator: tt.fields.Separator,
			}
			gotRespRes, err := str.Handle(tt.args.ctx, tt.args.reqRes)
			if (err != nil) != tt.wantErr {
				t.Errorf("String.Handle() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotRespRes, tt.wantRespRes) {
				t.Errorf("String.Handle() = %v, want %v", gotRespRes, tt.wantRespRes)
			}
		})
	}
}
