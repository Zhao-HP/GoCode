package util

import (
	"bytes"
	"encoding/json"
	"fmt"
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
