package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"sync"
	"time"
)

func httpGet(url string) string {
	res, err := http.Get(url)
	if err != nil {
		fmt.Printf("error = %s", err.Error())
		//nErr := fmt.Errorf("unknow  error ")//nErr := errors.New("unknow err")
		//fmt.Println(nErr.Error())
		return ""
	}
	resu, _ := ioutil.ReadAll(res.Body)
	return string(resu)
}
func httpPostJson(url string, kv map[string]string) string {
	body, err := json.Marshal(kv)
	if err != nil {
		return ""
	}
	rw := bytes.NewReader(body)
	resp, err := http.Post(url, "Application/json", rw)
	if err != nil {
		fmt.Printf("error = %s", err.Error())
		return ""
	}
	b, _ := ioutil.ReadAll(resp.Body)
	return string(b)
}

//
//以下是自定义，普通的http客户端，也就是不实用默认的，使用了once来实现了单例模式。
var httpClient *http.Client
var once sync.Once

func GetHttpClient() *http.Client {
	once.Do(func() {
		if httpClient == nil {
			httpClient = new(http.Client)
			httpClient.Timeout = time.Minute * 2
		}
	})
	return httpClient
}
func SendJsonHttpPost(url string, body io.Reader) (resp *http.Response, err error) {
	resp, err = GetHttpClient().Post(url, "application/json", body)
	return
}

//用于禁止跳转的http客户端
var authHttpClient *http.Client
var authOnce sync.Once

func GetAuthHttpClient() *http.Client {
	authOnce.Do(func() {
		if authHttpClient == nil {
			authHttpClient = new(http.Client)
			authHttpClient.Timeout = time.Minute
			authHttpClient.CheckRedirect = func(req *http.Request, via []*http.Request) error {
				return http.ErrUseLastResponse
			}
		}
	})
	return authHttpClient
}

func main() {
	fmt.Println(httpGet("http://www.baidu.com"))
	mp := make(map[string]string, 0)
	mp["name"] = "ita"
	mp["age"] = "24"
	mp["basic_info"] = "no girl friend"
	fmt.Println(httpPostJson("http://localhost:7890/postjson", mp))
}
