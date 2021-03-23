package httpclient

import (
    "bufio"
    "bytes"
    "crypto/tls"
    "errors"
    "fmt"
    "io"
    "io/ioutil"
    "log"
    "net/http"
    "net/url"
    "os"
    "strings"
    "time"
)

type HttpClient struct {
    Method  string // 支持 get form<post-form>  json<post-json>
    Link    string
    Headers map[string]string
    Body    string
    Proxy   string // socks5://127.0.0.1:10808
}

// HTTP远程请求
func (hc HttpClient) Request() ([]byte, error) {
    m := strings.ToUpper(hc.Method)
    if hc.Headers == nil {
        hc.Headers = make(map[string]string, 0)
    }
    if _, ok := hc.Headers["Content-Type"]; !ok {
        switch m {
        case "FORM", "POST":
            hc.Headers["Content-Type"] = "application/x-www-form-urlencoded"
        case "JSON":
            hc.Headers["Content-Type"] = "application/json"
        }
    }
    if m == "JSON" || m == "FORM" {
        hc.Method = "POST"
    }
    var proxy func(*http.Request) (*url.URL, error) = nil
    if hc.Proxy != "" {
        proxy = func(_ *http.Request) (*url.URL, error) {
            return url.Parse(hc.Proxy)
        }
    }
    tr := &http.Transport{
        TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
        Proxy:           proxy,
    }
    client := &http.Client{
        Timeout:   60 * time.Second,
        Transport: tr,
    }
    hc.Method = strings.ToUpper(hc.Method)
    req, err := http.NewRequest(hc.Method, hc.Link, strings.NewReader(hc.Body))
    if err != nil {
        errMsg := fmt.Sprintf("request %s error: %s", hc.Link, err.Error())
        err = errors.New(errMsg)
        return nil, err
    }
    for key, val := range hc.Headers {
        req.Header.Set(key, val)
    }
    resp, err := client.Do(req)
    defer func() {
        err = resp.Body.Close()
        if err != nil {
            log.Println(err)
        }
    }()
    rawResponse, err := ioutil.ReadAll(resp.Body)
    if resp.StatusCode != 200 {
        errMsg := fmt.Sprintf("request %s response code: %d\n%s\n", hc.Link, resp.StatusCode, string(rawResponse))
        err = errors.New(errMsg)
    }
    return rawResponse, err
}

// 远程下载
func Download(resource string, dest string) error {
    hc := HttpClient{
        Method: "get",
        Link:   resource,
    }
    rawResp, err := hc.Request()
    if err != nil {
        return err
    }
    // 获得get请求响应的reader对象
    reader := bufio.NewReaderSize(bytes.NewReader(rawResp), 32*1024)
    out, err := os.Create(dest)
    if err != nil {
        return err
    }
    defer out.Close()
    // 获得文件的writer对象
    writer := bufio.NewWriter(out)
    if _, err = io.Copy(writer, reader); err != nil {
        return err
    }
    return nil
}

// QQ消息推送
func PushMsg(msg string) ([]byte, error) {
    hc := HttpClient{
        Method: "get",
        Link:   "https://qmsg.zendee.cn/send/49165175f256124c3da594b51cda2274?msg=" + url.QueryEscape(msg),
    }
    return hc.Request()
}
