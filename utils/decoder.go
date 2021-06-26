package utils

import (
	"encoding/json"
	"fmt"

	"github.com/mitchellh/mapstructure"
)

func PrintJSON(in interface{}) {
	data, _ := json.MarshalIndent(in, "", "    ")
	fmt.Println(string(data))
}

func DecodeTypedWeakly(in interface{}, out interface{}) error {
	decoder, err := mapstructure.NewDecoder(&mapstructure.DecoderConfig{
		TagName:          "json",
		Result:           out,
		WeaklyTypedInput: true,
	})

	if err != nil {
		return err
	}

	return decoder.Decode(in)
}
