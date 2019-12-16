package pipelines

import (
	"reflect"
	"testing"
)

func TestDep_HasDepID(t *testing.T) {
	type fields struct {
		ID    string
		Deps  []*Dep
		Deped []*Dep
	}
	type args struct {
		id string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   []string
		want1  bool
	}{
		{
			name: "has dep id",
			fields: fields{
				ID: "a",
				Deps: []*Dep{
					{
						ID: "b",
					},
					{
						ID: "c",
						Deps: []*Dep{
							{
								ID: "d",
							},
						},
					},
				},
			},
			args: args{
				id: "d",
			},
			want:  []string{"a", "c", "d"},
			want1: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			dep := Dep{
				ID:    tt.fields.ID,
				Deps:  tt.fields.Deps,
				Deped: tt.fields.Deped,
			}
			got, got1 := dep.HasDepID(tt.args.id)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Dep.HasDepID() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("Dep.HasDepID() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
