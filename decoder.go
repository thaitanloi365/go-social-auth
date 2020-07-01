package auth

import "github.com/mitchellh/mapstructure"

// DecodeTypedWeakly deacode typed weakly
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
