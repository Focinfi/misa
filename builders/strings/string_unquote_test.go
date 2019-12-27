package strings

import (
	"context"
	"reflect"
	"testing"

	"github.com/Focinfi/go-pipeline"
)

func TestParserUtf8_Handle(t *testing.T) {
	type args struct {
		ctx    context.Context
		reqRes *pipeline.HandleRes
	}
	tests := []struct {
		name        string
		p           Unquote
		args        args
		wantRespRes *pipeline.HandleRes
		wantErr     bool
	}{
		{
			name: "normal",
			p:    Unquote{},
			args: args{
				reqRes: &pipeline.HandleRes{
					Data: `"\u4e2d\u56fd"`,
				},
			},
			wantRespRes: &pipeline.HandleRes{
				Status: pipeline.HandleStatusOK,
				Data:   `"中国"`,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotRespRes, err := tt.p.Handle(tt.args.ctx, tt.args.reqRes)
			if (err != nil) != tt.wantErr {
				t.Errorf("Unquote.Handle() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotRespRes, tt.wantRespRes) {
				t.Errorf("Unquote.Handle() = %v, want %v", gotRespRes, tt.wantRespRes)
			}
		})
	}
}
