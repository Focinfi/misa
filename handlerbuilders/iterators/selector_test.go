package iterators

import (
	"context"
	"reflect"
	"testing"

	"github.com/Focinfi/go-pipeline"
)

func TestNewSelector(t *testing.T) {
	type args struct {
		interpreterName string
		script          string
	}
	tests := []struct {
		name     string
		args     args
		reqRes   *pipeline.HandleRes
		wantResp *pipeline.HandleRes
		wantErr  bool
	}{
		{
			name: "odd number",
			args: args{
				interpreterName: "tengo",
				script:          "int(value) % 2 == 0",
			},
			reqRes: &pipeline.HandleRes{
				Data: []interface{}{
					1, "2", "3", 4,
				},
			},
			wantResp: &pipeline.HandleRes{
				Status: pipeline.HandleStatusOK,
				Data: []interface{}{
					"2", int64(4),
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewSelector(Conf{
				InterpreterName: tt.args.interpreterName,
				Script:          tt.args.script,
			})
			if err != nil {
				t.Fatal(err)
			}
			resp, err := got.Handle(context.Background(), tt.reqRes)
			if (err != nil) != tt.wantErr {
				t.Errorf("Selector.Handle() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(resp, tt.wantResp) {
				t.Errorf("Selector.Handle() = %#v, want %#v", resp, tt.wantResp)
			}
		})
	}
}
