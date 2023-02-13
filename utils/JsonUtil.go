package util

import (
	"bytes"
	"encoding/json"
	"fmt"
	jsoniter "github.com/json-iterator/go"
	"github.com/pkg/errors"
)

// PrintMapJsonStr 将Map按照Json格式输出
func PrintMapJsonStr(i map[string]interface{}) {
	s, err := MapToJsonFormatString(i)
	if err != nil {
		return
	}
	fmt.Println(s)
}

// MapToJsonFormatString Map转换成Json格式的字符串
func MapToJsonFormatString(i map[string]interface{}) (string, error) {
	bs, err := json.Marshal(i)
	if err != nil {
		return "", err
	}

	var out bytes.Buffer
	err = json.Indent(&out, bs, "", "")
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%s", out.String()), nil
}

// MapToJsonStr 将Map转换成Json字符串，没有换行等格式化
func MapToJsonStr(i map[string]interface{}) string {
	b, _ := json.Marshal(i)
	return string(b)
}

func interfaceToStr[T any](i T) string {
	b, _ := json.Marshal(i)
	return string(b)
}

func StrToStruct[T any](str string) (T, error) {
	result := new(T)
	if len(str) == 0 {
		return *result, errors.Errorf("JSON字符串为空")
	}

	err := jsoniter.UnmarshalFromString(str, result)

	return *result, err
}

func StrToStructArr[T any](str string) ([]T, error) {
	if len(str) == 0 {
		return nil, errors.Errorf("JSON字符串为空")
	}

	result := new([]T)
	err := jsoniter.UnmarshalFromString(str, result)

	return *result, err
}
