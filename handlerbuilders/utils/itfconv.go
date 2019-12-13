package utils

import (
	"encoding/json"
	"errors"
	"fmt"
	"reflect"
	"strconv"
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
		return nil, errors.New("must a slice")
	}

	items := make([]interface{}, v.Len())
	for i := 0; i < v.Len(); i++ {
		items[i] = v.Index(i).Interface()
	}
	return items, nil
}

func AnyTypeToInt64(data interface{}) (int64, error) {
	return strconv.ParseInt(fmt.Sprint(data), 10, 64)
}
