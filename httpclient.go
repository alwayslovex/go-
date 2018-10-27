package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
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

func main() {
	fmt.Println(httpGet("http://www.baidu.com"))
	mp := make(map[string]string, 0)
	mp["name"] = "ita"
	mp["age"] = "24"
	mp["basic_info"] = "no girl friend"
	fmt.Println(httpPostJson("http://localhost:7890/postjson", mp))
}
