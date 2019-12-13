package iterators

import (
	"fmt"

	"github.com/Focinfi/misa/handlerbuilders/utils"
)

func InterfaceToSlice(data interface{}) ([]interface{}, error) {
	items, err := utils.AynTypeToSlice(data)
	if err != nil {
		return nil, fmt.Errorf("request data type wrong: %v", err)
	}
	return items, nil
}
