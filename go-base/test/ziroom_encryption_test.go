package test

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/goinggo/mapstructure"
	"math/rand"
	"net/http"
	"strconv"
	"testing"
	"time"
)

type BaseResult struct {
	Success      bool        `json:"success"`
	ErrorMessage string      `json:"errorMessage"`
	Message      string      `json:"message"`
	Data         interface{} `json:"data"`
	Code         int         `json:"code"`
}

type Protection struct {
	ProtectionCode         string `json:"protectionCode"`
	ProtectionCodeSecurity string `json:"protectionCodeSecurity"`
}

const baseUrl = "http://127.0.0.1:7001/tass/"

func TestMybatisPlus(t *testing.T) {

	for i := 0; i < 600; i++ {
		fmt.Println("===================================================================================================")
		fmt.Println("第[", i, "]次请求")
		if e, i, err := Encryption(); err == nil {
			_ = Decryption(e, i)
		} else {
			fmt.Println(err)
		}
	}
}

func Encryption() (string, int, error) {
	var r = rand.New(rand.NewSource(time.Now().Unix()))
	i := r.Intn(1000)
	i = 0
	if i&1 == 0 {
		fmt.Println("AES软加密")
	} else {
		fmt.Println("TASS硬加密")
	}
	encryptionUrl := baseUrl + "encryption?keyMode=" + strconv.Itoa(i)
	resp, err := http.Get(encryptionUrl)
	if err != nil {
		fmt.Println(err)
		return "", 0, err
	}
	if resp.StatusCode != 200 {
		fmt.Println("Http请求的状态码为: ", resp.StatusCode)
		return "", 0, errors.New("请求失败")
	}

	br := BaseResult{}

	_ = json.NewDecoder(resp.Body).Decode(&br)
	return br.Message, i, nil
}

func Decryption(code string, i int) error {

	if i&1 == 1 {
		fmt.Println("AES软解密")
	} else {
		fmt.Println("TASS硬解密")
	}

	decryptionUrl := baseUrl + "decryption?uuid=" + code + "&keyMode=" + strconv.Itoa(i)
	resp, _ := http.Get(decryptionUrl)
	if resp.StatusCode != 200 {
		fmt.Println("Http请求的状态码为: ", resp.StatusCode)
		return errors.New("请求失败")
	}

	br := BaseResult{}
	_ = json.NewDecoder(resp.Body).Decode(&br)

	p := Protection{}
	_ = mapstructure.Decode(br.Data.(map[string]interface{}), &p)

	if p.ProtectionCode == code {
		fmt.Println("校验通过")
	} else {
		fmt.Println("校验失败")
	}

	return nil
}
