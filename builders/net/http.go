package net

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/Focinfi/go-pipeline"
)

var DefaultHTTP = HTTP{}

type HTTP struct{}

func (HTTP) Handle(ctx context.Context, reqRes *pipeline.HandleRes) (respRes *pipeline.HandleRes, err error) {
	respRes = &pipeline.HandleRes{}

	if reqRes == nil {
		return
	}
	respRes, err = reqRes.Copy()
	if err != nil {
		return nil, err
	}
	if reqRes.Data == nil {
		return
	}

	data := respRes.Data.(map[string]interface{})
	method := fmt.Sprint(data["method"])
	url := fmt.Sprint(data["url"])
	header := make(map[string]interface{})
	if h, ok := data["header"]; ok && h != nil {
		header = h.(map[string]interface{})
	}
	body := fmt.Sprint(data["body"])

	httpReq, err := http.NewRequest(method, url, strings.NewReader(body))
	if err != nil {
		return nil, err
	}
	for k, v := range header {
		httpReq.Header.Set(k, fmt.Sprint(v))
	}
	cli := http.Client{}
	resp, err := cli.Do(httpReq)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	rtData := make(map[string]interface{})
	rtData["status_code"] = resp.StatusCode
	rtData["header"] = map[string][]string(resp.Header)
	rtData["body"] = string(respBody)
	respRes.Data = rtData

	return respRes, nil
}
