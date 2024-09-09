package sdkreq

import (
	"encoding/json"
)

func ConvertInterfaceToJson(v interface{}) []byte {
	data, err := json.Marshal(v)
	if err != nil {
		return nil
	}
	return data
}
