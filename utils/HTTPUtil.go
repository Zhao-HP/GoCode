package util

import (
	"encoding/json"
	"github.com/valyala/fasthttp"
	"strings"
	"time"
)

func newCustomClient(writeTimeout, readTimeout time.Duration) *fasthttp.Client {
	client := &fasthttp.Client{
		MaxIdleConnDuration:           time.Duration(29) * time.Second,
		DisableHeaderNamesNormalizing: true,
		MaxIdemponentCallAttempts:     3,
		//DisablePathNormalizing:        true,
		ReadTimeout:         time.Millisecond * readTimeout,
		WriteTimeout:        time.Millisecond * writeTimeout,
		MaxConnsPerHost:     1000,
		MaxResponseBodySize: 1024 * 1024 * 10,
	}
	return client
}

func Post(url string, params map[string]string, header map[string]string) ([]byte, error) {
	req := fasthttp.AcquireRequest()
	defer fasthttp.ReleaseRequest(req)
	resp := fasthttp.AcquireResponse()
	defer fasthttp.ReleaseResponse(resp)

	// 设置请求头
	req.Header.SetMethod("POST")
	if len(header) > 0 {
		for k, v := range header {
			req.Header.Add(k, v)
		}
	}

	// 设置请求体
	if len(params) > 0 {
		b, err := json.Marshal(params)
		if err == nil {
			req.SetBodyRaw(b)
		}
	}

	req.SetRequestURI(url)

	// 请求接口，超时2秒
	err := newCustomClient(time.Second*2, time.Second*2).Do(req, resp)
	if err != nil {
		return nil, err
	}

	b := resp.Body()
	return b, nil
}

func SimpleGet(url string, header map[string]string) ([]byte, error) {

	req := fasthttp.AcquireRequest()
	defer fasthttp.ReleaseRequest(req)
	resp := fasthttp.AcquireResponse()
	defer fasthttp.ReleaseResponse(resp)

	req.Header.SetMethod("GET")

	if len(header) > 0 {
		for k, v := range header {
			req.Header.Add(k, v)
		}
	}
	req.SetRequestURI(url)
	// 请求接口，超时2秒
	err := newCustomClient(time.Second*2, time.Second*2).Do(req, resp)
	if err != nil {
		return nil, err
	}

	b := resp.Body()
	return b, nil
}

func GetParamMap(url string, params map[string]string, header map[string]string) ([]byte, error) {

	builder := strings.Builder{}
	builder.WriteString(url)
	builder.WriteString("?")

	// 拼接参数
	if len(params) > 0 {
		for k, v := range params {
			builder.WriteString(k)
			builder.WriteString("=")
			builder.WriteString(v)
			builder.WriteString("&")
		}
	}

	b, err := SimpleGet(builder.String(), header)

	return b, err
}
