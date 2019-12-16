package pipelines

import (
	"testing"
)

func Test_parseConfJSON(t *testing.T) {
	confs, err := parseConfJSON("../configs/conf.example.json")
	if err != nil {
		t.Fatal(err)
	}
	t.Log(len(confs))
}
