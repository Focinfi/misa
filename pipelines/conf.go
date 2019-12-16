package pipelines

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type pipeConf struct {
	ID   string          `json:"id"`
	Conf json.RawMessage `json:"conf"`
}

func parseConfJSON(confPath string) ([]pipeConf, error) {
	pipeConfs := make([]pipeConf, 0)

	b, err := ioutil.ReadFile(confPath)
	if err != nil {
		return nil, fmt.Errorf("read conf file err: %v", err)
	}

	if err := json.Unmarshal(b, &pipeConfs); err != nil {
		return nil, fmt.Errorf("unmarshal conf file in json format err: %v", err)
	}
	return pipeConfs, nil
}
