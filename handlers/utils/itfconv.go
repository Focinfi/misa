package utils

import (
	"encoding/json"
	"errors"
	"reflect"
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

func AynTypeToSlice(data interface{}) ([]interface{}, error) {
	t := reflect.TypeOf(data).Kind()
	v := reflect.ValueOf(data)
	if t == reflect.Ptr {
		t = reflect.TypeOf(data).Elem().Kind()
		v = reflect.ValueOf(data).Elem()
	}
	if t != reflect.Slice {
		return nil, errors.New("request data must a slice")
	}

	items := make([]interface{}, v.Len())
	for i := 0; i < v.Len(); i++ {
		items[i] = v.Index(i).Interface()
	}
	return items, nil
}
