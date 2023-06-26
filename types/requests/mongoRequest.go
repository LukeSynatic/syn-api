package requests

import (
	"encoding/json"
	"reflect"
)

type MongoReq struct {
	DataSource string `json:"dataSource"`
	Database   string `json:"database"`
	Collection string `json:"collection"`
}

func Opts(reqBody []byte, optionsType reflect.Type) (interface{}, error) {
	options := reflect.New(optionsType).Interface()

	if err := json.Unmarshal(reqBody, &options); err != nil {
		return nil, err
	}

	return options, nil
}

func Req(reqBody []byte, optionsType reflect.Type) (interface{}, error) {
	req := reflect.New(optionsType).Interface()

	if err := json.Unmarshal(reqBody, &req); err != nil {
		return nil, err
	}

	return req, nil
}
