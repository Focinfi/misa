package pipelines

import (
	"testing"
)

func Test_initHandlers(t *testing.T) {
	if err := InitHandlers("../configs/conf.example.json"); err != nil {
		t.Fatal(err)
	}
	t.Log(PipelineMap)
}

func Test_pipelineMap_UpdatePipeline(t *testing.T) {
	type args struct {
		id       string
		confJSON string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "normal",
			args: args{
				id: "parse-json",
				confJSON: `
	[
      {
        "desc": "parse json string",
        "timeout": 1000,
        "required": true,
        "handler_builder_name": "parser-json"
      }
    ]`,
			},

			wantErr: false,
		},
		{
			name: "cycle deps error",
			args: args{
				id: "parse-json",
				confJSON: `
	[
      {
        "desc": "parse json string",
        "timeout": 1000,
        "required": true,
        "handler_builder_name": "parser-json"
      },
      {
        "desc": "notify-desktop",
        "timeout": 100,
        "required": true,
        "ref_handler_id": "notify-desktop"
      }
    ]`,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := InitHandlers("../configs/conf.example.json")
			if err != nil {
				t.Fatal(err)
			}
			err = PipelineMap.UpdatePipeline(tt.args.id, tt.args.confJSON)
			t.Log("err:", err)
			if (err != nil) != tt.wantErr {
				t.Errorf("pipelineMap.UpdatePipeline() error = %v, wantErr %v", err, tt.wantErr)
			}
			t.Log((*PipelineMap)[tt.args.id].Handler)
		})
	}
}
