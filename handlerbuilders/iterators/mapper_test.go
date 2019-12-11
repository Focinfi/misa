package iterators

import (
	"context"
	"reflect"
	"testing"

	"github.com/Focinfi/go-pipeline"
)

func TestNewMapper(t *testing.T) {
	type args struct {
		interpreterName string
		mapScript       string
	}
	tests := []struct {
		name     string
		args     args
		reqRes   *pipeline.HandleRes
		wantResp *pipeline.HandleRes
		wantErr  bool
	}{
		{
			name: "square string number",
			args: args{
				interpreterName: "tengo",
				mapScript:       "int(value) * int(value)",
			},
			reqRes: &pipeline.HandleRes{
				Data: []string{
					"1", "2", "3", "4",
				},
			},
			wantResp: &pipeline.HandleRes{
				Status: pipeline.HandleStatusOK,
				Data: []interface{}{
					int64(1), int64(4), int64(9), int64(16),
				},
			},
			wantErr: false,
		},
		{
			name: "square mix number",
			args: args{
				interpreterName: "tengo",
				mapScript:       "int(value) * int(value)",
			},
			reqRes: &pipeline.HandleRes{
				Data: []interface{}{
					1, "2", "3", 4,
				},
			},
			wantResp: &pipeline.HandleRes{
				Status: pipeline.HandleStatusOK,
				Data: []interface{}{
					int64(1), int64(4), int64(9), int64(16),
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewMapper(tt.args.interpreterName, tt.args.mapScript)
			if err != nil {
				t.Fatal(err)
			}
			resp, err := got.Handle(context.Background(), tt.reqRes)
			if (err != nil) != tt.wantErr {
				t.Errorf("Mapper.Handle() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(resp, tt.wantResp) {
				t.Errorf("Mapper.Handle() = %v, want %v", resp, tt.wantResp)
			}
		})
	}
}
