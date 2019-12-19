package parsers

import (
	"context"
	"fmt"

	"github.com/360EntSecGroup-Skylar/excelize"
	"github.com/Focinfi/go-pipeline"
)

var DefaultXLSXParser = XLSXParser{}

type XLSXParser struct{}

func (p XLSXParser) Handle(ctx context.Context, reqRes *pipeline.HandleRes) (respRes *pipeline.HandleRes, err error) {
	respRes = &pipeline.HandleRes{}

	if reqRes != nil {
		respRes, err = reqRes.Copy()
		if err != nil {
			return nil, err
		}

		if reqRes.Data != nil {
			conf := reqRes.Data.(map[string]interface{})
			filePath := fmt.Sprint(conf["path"])
			sheet := fmt.Sprint(conf["sheet"])
			useFirstRowAsHead := fmt.Sprint(conf["first_row_as_head"]) == "true"
			f, err := excelize.OpenFile(filePath)
			if err != nil {
				return nil, err
			}
			rows, err := f.Rows(fmt.Sprint(sheet))
			if err != nil {
				return nil, err
			}
			data := make([]map[string]interface{}, 0)

			head := make([]string, 0)
			if useFirstRowAsHead {
				if rows.Next() {
					head = rows.Columns()
				} else {
					return respRes, nil
				}
			}

			for rows.Next() {
				data = append(data)
				columns := rows.Columns()
				row := make(map[string]interface{}, len(columns))
				for i, col := range columns {
					colName := fmt.Sprint(i)
					if len(head) >= i+1 {
						colName = head[i]
					}
					row[colName] = col
				}
				data = append(data, row)
			}
			respRes.Data = data
		}
	}

	respRes.Status = pipeline.HandleStatusOK
	return respRes, nil
}
