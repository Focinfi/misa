package utils

import (
	"reflect"
	"testing"
)

func TestAynTypeToSlice(t *testing.T) {
	type args struct {
		data interface{}
	}
	tests := []struct {
		name    string
		args    args
		want    []interface{}
		wantErr bool
	}{
		{
			name: "not a slice",
			args: args{
				data: 1,
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "a string slice",
			args: args{
				data: []string{"a", "c"},
			},
			want:    []interface{}{"a", "c"},
			wantErr: false,
		},
		{
			name: "a string slice pointer",
			args: args{
				data: &[]string{"a", "c"},
			},
			want:    []interface{}{"a", "c"},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := AynTypeToSlice(tt.args.data)
			if (err != nil) != tt.wantErr {
				t.Errorf("AynTypeToSlice() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AynTypeToSlice() = %v, want %v", got, tt.want)
			}
		})
	}
}
