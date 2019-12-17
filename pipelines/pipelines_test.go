package pipelines

import (
	"testing"
)

func Test_initHandlers(t *testing.T) {
	pipelines, err := InitLines("../configs/conf.example.json")
	if err != nil {
		t.Fatal(err)
	}
	t.Log(pipelines)
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
		{
			name: "cycle deps self",
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
        "ref_handler_id": "parse-json"
      }
    ]`,
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			lines, err := InitLines("../configs/conf.example.json")
			if err != nil {
				t.Fatal(err)
			}
			err = lines.UpdateByConfJSON(tt.args.id, tt.args.confJSON)
			t.Log("err:", err)
			if (err != nil) != tt.wantErr {
				t.Errorf("lines.UpdateByConfJSON() error = %v, wantErr %v", err, tt.wantErr)
			}
			t.Log((*lines).LineMap[tt.args.id].Handler)
		})
	}
}

func Test_pipelines_Delete(t *testing.T) {
	type args struct {
		id string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "normal",
			args: args{
				id: "get-first-pipeline-id",
			},
			wantErr: false,
		},
		{
			name: "has deped",
			args: args{
				id: "parse-json",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p, err := InitLines("../configs/conf.example.json")
			if err != nil {
				t.Fatal(err)
			}
			err = p.Delete(tt.args.id)
			t.Log("err:", err)
			if (err != nil) != tt.wantErr {
				t.Errorf("lines.Delete() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
