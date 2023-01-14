package utils

import "encoding/json"

func StrToJson(input string, output interface{}) error {
	err2 := json.Unmarshal([]byte(input), &output)
	if err2 != nil {
		return err2
	}
	return nil
}

func JsonToStr(input interface{}) (string, error) {
	marshal, err := json.Marshal(input)
	if err != nil {
		return "", err
	}
	return string(marshal), nil
}
