package generators

import (
	"context"
	"reflect"
	"testing"
	"time"

	"github.com/Focinfi/go-pipeline"
)

func TestTimeRange_Handle(t *testing.T) {
	type args struct {
		ctx    context.Context
		reqRes *pipeline.HandleRes
	}
	tests := []struct {
		name        string
		tr          TimeRange
		args        args
		wantRespRes *pipeline.HandleRes
		wantErr     bool
	}{
		{
			name: "normal",
			tr:   TimeRange{},
			args: args{
				ctx: context.Background(),
				reqRes: &pipeline.HandleRes{
					Data: map[string]interface{}{
						"start": time.Date(2019, 12, 1, 0, 0, 0, 0, time.Local),
						"end":   time.Date(2019, 12, 2, 0, 0, 0, 0, time.Local),
						"step":  "7h",
					},
				},
			},
			wantRespRes: &pipeline.HandleRes{
				Status: pipeline.HandleStatusOK,
				Data: []time.Time{
					time.Date(2019, 12, 1, 0, 0, 0, 0, time.Local),
					time.Date(2019, 12, 1, 7, 0, 0, 0, time.Local),
					time.Date(2019, 12, 1, 14, 0, 0, 0, time.Local),
					time.Date(2019, 12, 1, 21, 0, 0, 0, time.Local),
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tr := TimeRange{}
			gotRespRes, err := tr.Handle(tt.args.ctx, tt.args.reqRes)
			if (err != nil) != tt.wantErr {
				t.Errorf("TimeRange.Handle() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotRespRes, tt.wantRespRes) {
				t.Errorf("TimeRange.Handle() = %v, want %v", gotRespRes, tt.wantRespRes)
			}
		})
	}
}
