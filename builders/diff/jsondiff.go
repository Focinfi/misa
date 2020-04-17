package diff

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"strings"

	"github.com/Focinfi/go-pipeline"
	"github.com/Focinfi/misa/builders/utils"
	jd "github.com/josephburnett/jd/lib"
)

var DefaultJSON = JSON{}

type JSON struct {
}

func (J JSON) Handle(ctx context.Context, reqRes *pipeline.HandleRes) (respRes *pipeline.HandleRes, err error) {
	if reqRes != nil && reqRes.Data != nil {
		respRes, err = reqRes.Copy()
		if err != nil {
			return nil, err
		}

		params := reqRes.Data.(map[string]interface{})
		pathA, err := utils.AnyTypeToString(params["path_a"])
		if err != nil {
			return nil, fmt.Errorf("param path_a invalid: %v", err.Error())
		}
		pathB, err := utils.AnyTypeToString(params["path_b"])
		if err != nil {
			return nil, fmt.Errorf("param path_b invalid: %v", err.Error())
		}
		jsonA, err := jd.ReadJsonFile(pathA)
		if err != nil {
			return nil, fmt.Errorf("read file from path_a invalid: %v", err.Error())
		}
		jsonB, err := jd.ReadJsonFile(pathB)
		if err != nil {
			return nil, fmt.Errorf("read file from path_b invalid: %v", err.Error())
		}
		diff := jsonA.Diff(jsonB)
		extraList := make([]string, 0)
		missingList := make([]string, 0)
		inconsistentList := make([]string, 0)
		for _, d := range diff {
			path, err := json.Marshal(d.Path)
			if err != nil {
				return nil, err
			}

			if len(d.OldValues) == 0 && len(d.NewValues) > 0 { // missing
				value, err := json.Marshal(d.OldValues)
				if err != nil {
					return nil, err
				}

				missingList = append(missingList, string(path)+": "+string(value))
			} else if len(d.OldValues) > 0 && len(d.NewValues) == 0 { // extra
				value, err := json.Marshal(d.OldValues)
				if err != nil {
					return nil, err
				}

				extraList = append(extraList, string(path)+": "+string(value))
			} else { // inconsistent
				oldValue, err := json.Marshal(d.OldValues)
				if err != nil {
					return nil, err
				}

				newValue, err := json.Marshal(d.NewValues)
				if err != nil {
					return nil, err
				}

				inconsistentList = append(inconsistentList, string(path)+": "+string(oldValue)+" => "+string(newValue))
			}
		}

		var buf strings.Builder
		if len(extraList) > 0 {
			buf.WriteString("Extra:\n")
			buf.WriteString(strings.Join(extraList, "\n"))
		}
		if len(missingList) > 0 {
			buf.WriteString("\nMissing:\n")
			buf.WriteString(strings.Join(missingList, "\n"))
		}
		if len(inconsistentList) > 0 {
			buf.WriteString("\nInconsistent:\n")
			buf.WriteString(strings.Join(inconsistentList, "\n"))
		}
		respRes.Status = pipeline.HandleStatusOK
		respRes.Data = buf.String()
		return respRes, nil
	}
	return nil, errors.New("need params path_a and path_b")
}
