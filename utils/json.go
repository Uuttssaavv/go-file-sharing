package utils

import (
	"encoding/json"
)

func ObjectToJson[T any](object interface{}, data *T) {

	jason, _ := json.Marshal(object)

	json.Unmarshal([]byte(jason), &data)
}

func StringToEntity[T any](value any, object *T) error {
	data, _ := json.Marshal(value)
	errors := json.Unmarshal([]byte(data), &object)
	return errors
}
