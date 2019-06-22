package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
)

//golang 中文件的操作，可以有不同的方式
//其中第一个,readfile 和第四个readall比较快。

func main() {
	body, _ := ioutil.ReadFile("abc")
	fmt.Println(string(body))

	f, _ := os.Open("abc")
	defer f.Close()
	buff := make([]byte, 100)
	n, _ := f.Read(buff)
	fmt.Println(string(buff[0:n]))
	f.Seek(0, 0)
	nf := bufio.NewReader(f)
	n, _ = nf.Read(buff)
	fmt.Println(string(buff[0:n]))
	f.Seek(0, 0)
	body, _ = ioutil.ReadAll(f)
	fmt.Println(string(body))
}
