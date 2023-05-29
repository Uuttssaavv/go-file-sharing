package utils

import (
	"encoding/json"
)

func ObjectToJson[T any](object interface{}, data *T) *T {

	jason, _ := json.Marshal(object)

	json.Unmarshal([]byte(jason), &data)

	return data
}
