package main

import (
	"net/http"
	"fmt"
	"io/ioutil"
)

func httpGet(url string) string{
	res,err := http.Get(url)
	if err != nil{
		fmt.Printf("error = %s",err.Error())
		//nErr := fmt.Errorf("unknow  error ")//nErr := errors.New("unknow err")
		//fmt.Println(nErr.Error())
		return ""
	}
	resu,_ := ioutil.ReadAll(res.Body)
	return string(resu)
}

func main(){
	fmt.Println(httpGet("http://www.baidu.com"))
}
