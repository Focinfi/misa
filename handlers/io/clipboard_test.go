package io

import (
	"context"
	"reflect"
	"testing"

	"github.com/Focinfi/go-pipeline"
	"github.com/go-vgo/robotgo/clipboard"
)

func TestClipboardReader_Handle(t *testing.T) {
	tests := []struct {
		name        string
		toCopy      string
		wantRespRes *pipeline.HandleRes
		wantErr     bool
	}{
		{
			name:   "normal",
			toCopy: "to copy line",
			wantRespRes: &pipeline.HandleRes{
				Status: pipeline.HandleStatusOK,
				Data:   "to copy line",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := clipboard.WriteAll(tt.toCopy); err != nil {
				t.Fatal(err)
			}
			cb := ClipboardReader{}
			gotRespRes, err := cb.Handle(context.Background(), nil)
			if (err != nil) != tt.wantErr {
				t.Errorf("ClipboardReader.Handle() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotRespRes, tt.wantRespRes) {
				t.Errorf("ClipboardReader.Handle() = %v, want %v", gotRespRes, tt.wantRespRes)
			}
		})
	}
}

func TestClipboardWriter_Handle(t *testing.T) {
	tests := []struct {
		name        string
		reqRes      *pipeline.HandleRes
		wantRespRes *pipeline.HandleRes
		copied      string
		wantErr     bool
	}{
		{
			name: "normal",
			reqRes: &pipeline.HandleRes{
				Data: "clipboard line",
			},
			wantRespRes: &pipeline.HandleRes{
				Status: pipeline.HandleStatusOK,
				Data:   "clipboard line",
			},
			copied:  "clipboard line",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cb := ClipboardWriter{}
			gotRespRes, err := cb.Handle(context.Background(), tt.reqRes)
			if (err != nil) != tt.wantErr {
				t.Errorf("ClipboardWriter.Handle() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotRespRes, tt.wantRespRes) {
				t.Errorf("ClipboardWriter.Handle() = %v, want %v", gotRespRes, tt.wantRespRes)
			}
			copied, err := clipboard.ReadAll()
			if err != nil {
				t.Fatal(err)
			}
			if copied != tt.copied {
				t.Errorf("clipboard: want=%v, got=%v", tt.copied, copied)
			}
		})
	}
}
