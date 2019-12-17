package sql

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/Focinfi/go-pipeline"
	"github.com/Focinfi/misa/handlerbuilders/utils"
	_ "github.com/go-sql-driver/mysql"
)

type MySQL struct {
	Dsn               string `json:"dsn" desc:"https://github.com/go-sql-driver/mysql#dsn-data-source-name" validate:"required"`
	MaxOpenConns      int    `json:"max_open_conns" desc:"max count of opening connections" validate:"gte=0"`
	MaxIdleConns      int    `json:"max_idle_conns" desc:"max count of idle connections" validate:"gte=0"`
	ConnMaxLifeSecond int    `json:"conn_max_life_second" desc:"max life time for connection" validate:"gte=0"`
	db                *sql.DB
}

func (m MySQL) Build() (pipeline.Handler, error) {
	db, err := sql.Open("mysql", m.Dsn)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}
	db.SetMaxOpenConns(m.MaxOpenConns)
	db.SetMaxIdleConns(m.MaxIdleConns)
	db.SetConnMaxLifetime(time.Second * time.Duration(m.ConnMaxLifeSecond))

	return &MySQL{
		Dsn:               m.Dsn,
		MaxOpenConns:      m.MaxOpenConns,
		MaxIdleConns:      m.MaxIdleConns,
		ConnMaxLifeSecond: m.ConnMaxLifeSecond,
		db:                db,
	}, nil
}

func (m MySQL) Handle(ctx context.Context, reqRes *pipeline.HandleRes) (respRes *pipeline.HandleRes, err error) {
	respRes = &pipeline.HandleRes{}
	if reqRes != nil {
		respRes, err = reqRes.Copy()
		if err != nil {
			return nil, err
		}

		if reqRes.Data == nil {
			return
		}
		data := reqRes.Data.(map[string]interface{})
		queryType := fmt.Sprint(data["query_type"])
		query := fmt.Sprint(data["query"])
		args, err := utils.AynTypeToSlice(data["args"])
		if err != nil {
			return nil, fmt.Errorf("param args in not a slice err: %v", err)
		}
		fmt.Printf("args: %#v", args)

		var respData interface{}
		switch queryType {
		case "exec":
			respData, err = m.exec(ctx, query, args...)
		case "query_rows":
			respData, err = m.queryRows(ctx, query, args...)
		default:
			return nil, fmt.Errorf("unsupported query type: %v", queryType)
		}
		if err != nil {
			return nil, err
		}
		respRes.Data = respData
	}

	respRes.Status = pipeline.HandleStatusOK
	return respRes, nil
}

func (m MySQL) exec(ctx context.Context, query string, args ...interface{}) (map[string]interface{}, error) {
	res, err := m.db.ExecContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return nil, fmt.Errorf("get affected rows err: %v", err)
	}
	return map[string]interface{}{
		"rows_affected": rowsAffected,
	}, nil
}

func (m MySQL) queryRows(ctx context.Context, query string, args ...interface{}) ([]map[string]string, error) {
	rows, err := m.db.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}

	columns, err := rows.Columns()
	if err != nil {
		return nil, err
	}

	data := make([]map[string]string, 0)
	values := make([]sql.RawBytes, len(columns))
	scanArgs := make([]interface{}, len(values))
	for i := range values {
		scanArgs[i] = &values[i]
	}

	for rows.Next() {
		err = rows.Scan(scanArgs...)
		if err != nil {
			panic(err.Error())
		}

		item := make(map[string]string, len(columns))
		for i, col := range values {
			val := make(sql.RawBytes, len(col))
			copy(val, col)
			item[columns[i]] = string(val)
		}
		data = append(data, item)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return data, nil
}
