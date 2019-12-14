package utils

import (
	"encoding/json"
	"fmt"
)

func JSONUnmarshalWithMap(data map[string]interface{}, dst interface{}) error {
	b, err := json.Marshal(data)
	if err != nil {
		return fmt.Errorf("marshal into JSON err: %v", err)
	}
	if err := json.Unmarshal(b, dst); err != nil {
		return fmt.Errorf("unmarshal err: %v", err)
	}
	if err := Validator.Struct(dst); err != nil {
		return fmt.Errorf("validate err: %v", err)
	}
	return nil
}
