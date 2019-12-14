package iterators

import (
	"context"
	"reflect"
	"testing"

	"github.com/Focinfi/go-pipeline"
)

func TestNewReducer(t *testing.T) {
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
			name: "sum number",
			args: args{
				interpreterName: "tengo",
				script:          "int(reduced) ? int(reduced) + int(value) : int(value)",
			},
			reqRes: &pipeline.HandleRes{
				Data: []interface{}{
					1, "2", "3", 4,
				},
			},
			wantResp: &pipeline.HandleRes{
				Status: pipeline.HandleStatusOK,
				Data:   int64(10),
			},
			wantErr: false,
		},
		{
			name: "join",
			args: args{
				interpreterName: "tengo",
				script:          `string(reduced) ? string(reduced) + "," + string(value) : string(value)`,
			},
			reqRes: &pipeline.HandleRes{
				Data: []interface{}{
					1, "2", "3", 4,
				},
			},
			wantResp: &pipeline.HandleRes{
				Status: pipeline.HandleStatusOK,
				Data:   "1,2,3,4",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewReducer(Conf{
				InterpreterName: tt.args.interpreterName,
				Script:          tt.args.script,
			})
			if err != nil {
				t.Fatal(err)
			}
			resp, err := got.Handle(context.Background(), tt.reqRes)
			if (err != nil) != tt.wantErr {
				t.Errorf("Reducer.Handle() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(resp, tt.wantResp) {
				t.Errorf("Reducer.Handle() = %v, want %v", resp, tt.wantResp)
			}
		})
	}
}
