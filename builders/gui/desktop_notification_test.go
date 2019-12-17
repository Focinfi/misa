package gui

import (
	"context"
	"os"
	"reflect"
	"testing"

	"github.com/Focinfi/go-pipeline"
)

func TestDesktopNotificator_Handle(t *testing.T) {
	if v := os.Getenv("TEST_DESKTOP"); v != "TRUE" {
		t.Skip()
	}

	type fields struct {
		AppName         string
		DefaultIconPath string
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
				AppName: "Misa Test",
			},
			args: args{
				ctx: context.Background(),
				reqRes: &pipeline.HandleRes{
					Data: map[string]interface{}{
						"title":   "Hello Misa",
						"text":    "Testing",
						"urgency": "critical",
					},
				},
			},
			wantRespRes: &pipeline.HandleRes{
				Status: pipeline.HandleStatusOK,
				Data: map[string]interface{}{
					"title":   "Hello Misa",
					"text":    "Testing",
					"urgency": "critical",
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			n := NewDesktopNotificator(tt.fields.AppName, tt.fields.DefaultIconPath)
			gotRespRes, err := n.Handle(tt.args.ctx, tt.args.reqRes)
			if (err != nil) != tt.wantErr {
				t.Errorf("DesktopNotificator.Handle() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotRespRes, tt.wantRespRes) {
				t.Errorf("DesktopNotificator.Handle() = %v, want %v", gotRespRes, tt.wantRespRes)
			}
		})
	}
}
