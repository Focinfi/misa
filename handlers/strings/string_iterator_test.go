package strings

import (
	"context"
	"reflect"
	"testing"

	"github.com/Focinfi/misa/handlers/iterators"

	"github.com/Focinfi/go-pipeline"
)

func TestNewStringIterator(t *testing.T) {
	type args struct {
		separator     string
		iteratorConfs []iterators.Conf
	}
	tests := []struct {
		name     string
		args     args
		reqRes   *pipeline.HandleRes
		wantResp *pipeline.HandleRes
		wantErr  bool
	}{
		{
			name: "normal",
			args: args{
				separator: ",",
				iteratorConfs: []iterators.Conf{
					{
						Type:            "select",
						InterpreterName: "tengo",
						Script:          "int(v) % 2 == 0",
					},
					{
						Type:            "map",
						InterpreterName: "tengo",
						Script:          "int(v) * int(v)",
					},
					{
						Type:            "reduce",
						InterpreterName: "tengo",
						Script:          "int(r) ? int(r) + int(v) : int(v)",
					},
				},
			},
			reqRes: &pipeline.HandleRes{
				Data: `1,2,3,4`,
			},
			wantResp: &pipeline.HandleRes{
				Status: pipeline.HandleStatusOK,
				Data:   int64(20),
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewStringIterator(tt.args.separator, tt.args.iteratorConfs)
			if err != nil {
				t.Fatal(err)
			}

			resp, err := got.Handle(context.Background(), tt.reqRes)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewStringIterator() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(resp, tt.wantResp) {
				t.Errorf("NewStringIterator() = %v, want %v", resp, tt.wantResp)
			}
		})
	}
}
