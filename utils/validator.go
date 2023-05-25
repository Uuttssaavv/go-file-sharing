package utils

import (
	"encoding/json"
	"github.com/go-playground/validator/v10"
	gpc "github.com/restuwahyu13/go-playground-converter"
)

type Error struct {
	Results struct {
		Errors []map[string]map[string]string `json:"errors"`
	} `json:"results"`
}

func GoValidator(s interface{}, config []gpc.ErrorMetaConfig) (interface{}, int) {
	var validate *validator.Validate
	validators := gpc.NewValidator(validate)
	bind := gpc.NewBindValidator(validators)

	errResponse, errCount := bind.BindValidator(s, config)

	var response Error
	data, _ := json.Marshal(errResponse)

	json.Unmarshal([]byte(data), &response)

	return response.Results.Errors, errCount
}
