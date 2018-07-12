package json

import (
	"github.com/json-iterator/go"
)

var jsonIter = jsoniter.ConfigDefault

func Marshal(data interface{}) ([]byte, error) {
	if serialized, err := jsonIter.Marshal(data); err != nil {
		return nil, err
	} else {
		return serialized, nil
	}
}

func Unmarshal(data []byte, v interface{}) error {
	if err := jsonIter.Unmarshal(data, v); err != nil {
		return err
	}
	return nil
}
