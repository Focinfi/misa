package generators

import (
	"context"
	"reflect"
	"testing"

	"github.com/Focinfi/go-pipeline"
)

func TestIntRange_Handle(t *testing.T) {
	type args struct {
		ctx    context.Context
		reqRes *pipeline.HandleRes
	}
	tests := []struct {
		name        string
		ir          IntRange
		args        args
		wantRespRes *pipeline.HandleRes
		wantErr     bool
	}{
		{
			name: "normal",
			ir:   IntRange{},
			args: args{
				ctx: context.Background(),
				reqRes: &pipeline.HandleRes{
					Data: map[string]interface{}{
						"start": 1,
						"end":   10,
						"step":  2,
					},
				},
			},
			wantRespRes: &pipeline.HandleRes{
				Status: pipeline.HandleStatusOK,
				Data: []int64{
					1, 3, 5, 7, 9,
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ir := IntRange{}
			gotRespRes, err := ir.Handle(tt.args.ctx, tt.args.reqRes)
			if (err != nil) != tt.wantErr {
				t.Errorf("IntRange.Handle() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotRespRes, tt.wantRespRes) {
				t.Errorf("IntRange.Handle() = %v, want %v", gotRespRes, tt.wantRespRes)
			}
		})
	}
}
