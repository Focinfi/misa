package parsers

import (
	"context"
	"reflect"
	"testing"
	"time"

	"github.com/Focinfi/go-pipeline"
)

func TestUnixParser_Handle(t *testing.T) {

	type args struct {
		ctx    context.Context
		reqRes *pipeline.HandleRes
	}
	tests := []struct {
		name        string
		u           UnixParser
		args        args
		wantRespRes *pipeline.HandleRes
		wantErr     bool
	}{
		{
			name: "normal int64",
			u:    UnixParser{},
			args: args{
				ctx: context.Background(),
				reqRes: &pipeline.HandleRes{
					Data: int64(0),
				},
			},
			wantRespRes: &pipeline.HandleRes{
				Status: pipeline.HandleStatusOK,
				Data:   time.Unix(0, 0),
			},
			wantErr: false,
		},
		{
			name: "normal string",
			u:    UnixParser{},
			args: args{
				ctx: context.Background(),
				reqRes: &pipeline.HandleRes{
					Data: "0",
				},
			},
			wantRespRes: &pipeline.HandleRes{
				Status: pipeline.HandleStatusOK,
				Data:   time.Unix(0, 0),
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := UnixParser{}
			gotRespRes, err := u.Handle(tt.args.ctx, tt.args.reqRes)
			if (err != nil) != tt.wantErr {
				t.Errorf("UnixParser.Handle() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotRespRes, tt.wantRespRes) {
				t.Errorf("UnixParser.Handle() = %v, want %v", gotRespRes, tt.wantRespRes)
			}
		})
	}
}
