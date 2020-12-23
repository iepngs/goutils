package httpclient

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

type HttpClient struct {
	Method  string
	Link    string
	Headers map[string]string
	Body    string
}

func (hc HttpClient) Request() (rawResponse []byte, err error) {
	method := strings.ToUpper(hc.Method)
	if method == "FORM"{
		method = "POST"
	}
	if hc.Headers == nil {
		hc.Headers = make(map[string]string, 0)
	}
	if _,ok := hc.Headers["Content-Type"]; !ok {
		if method == "JSON" {
			hc.Headers["Content-Type"] = "application/json"
		}else{
			hc.Headers["Content-Type"] = "application/x-www-form-urlencoded"
		}
	}
	hc.Method = method
	client := &http.Client{
		Timeout: 60 * time.Second,
	}
	req, err := http.NewRequest(hc.Method, hc.Link, strings.NewReader(hc.Body))
	if err != nil {
		errMsg := fmt.Sprintf("request %s error: %s", hc.Link, err.Error())
		err = errors.New(errMsg)
		return
	}
	for key, val := range hc.Headers {
		req.Header.Set(key, val)
	}
	resp, err := client.Do(req)
	defer func() {
		_ = resp.Body.Close()
	}()
	rawResponse, err = ioutil.ReadAll(resp.Body)
	if resp.StatusCode != 200 {
		errMsg := fmt.Sprintf("request %s response code: %d", hc.Link, resp.StatusCode)
		err = errors.New(errMsg)
		return
	}
	return
}
