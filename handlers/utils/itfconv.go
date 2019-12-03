package utils

import (
	"encoding/json"
)

func AnyTypeToString(src interface{}) (string, error) {
	switch data := src.(type) {
	case string:
		return data, nil
	case []byte:
		return string(data), nil
	default:
		b, err := json.Marshal(data)
		if err != nil {
			return "", err
		}
		return string(b), nil
	}
}
