package util

import (
	"fmt"
	"testing"
)

func TestGet(t *testing.T) {
	url := "http://127.0.0.1:9006/user/findUserByUid"
	params := map[string]string{
		"uid": "b7a7a1f6-a546-11ed-8a12-c6157a6d027e",
	}
	b, err := GetParamMap(url, params, nil)
	if err != nil {
		fmt.Println("异常")
		return
	}
	fmt.Println(string(b))
}

func TestPost(t *testing.T) {

}
