package formatters

import (
	"bytes"
	"context"
	"encoding/xml"
	"io"

	"github.com/Focinfi/go-pipeline"
	"github.com/Focinfi/misa/builders/utils"
)

var DefaultXML = &XML{}

type XML struct{}

func (x *XML) Handle(ctx context.Context, reqRes *pipeline.HandleRes) (respRes *pipeline.HandleRes, err error) {
	respRes = &pipeline.HandleRes{}
	if reqRes != nil {
		respRes, err = reqRes.Copy()
		if err != nil {
			return nil, err
		}

		if reqRes.Data != nil {
			data, err := utils.AnyTypeToString(reqRes.Data)
			if err != nil {
				return nil, err
			}

			b, err := formatXML([]byte(data))
			if err != nil {
				return nil, err
			}
			respRes.Data = string(b)
		}
	}

	respRes.Status = pipeline.HandleStatusOK
	return respRes, nil
}

func formatXML(data []byte) ([]byte, error) {
	b := &bytes.Buffer{}
	decoder := xml.NewDecoder(bytes.NewReader(data))
	encoder := xml.NewEncoder(b)
	encoder.Indent("", "  ")
	for {
		token, err := decoder.Token()
		if err == io.EOF {
			if err := encoder.Flush(); err != nil {
				return nil, err
			}
			return b.Bytes(), nil
		}
		if err != nil {
			return nil, err
		}
		err = encoder.EncodeToken(token)
		if err != nil {
			return nil, err
		}
	}
}
