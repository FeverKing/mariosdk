package sdkreq

import (
	"encoding/json"
)

func ConvertInterfaceToJson(v interface{}) []byte {
	if v == nil {
		return nil
	}
	data, err := json.Marshal(v)
	if err != nil {
		return nil
	}
	return data
}
