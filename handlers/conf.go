package handlers

import (
	"encoding/json"
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
		return nil, err
	}

	if err := json.Unmarshal(b, &pipeConfs); err != nil {
		return nil, err
	}

	return pipeConfs, nil
}
