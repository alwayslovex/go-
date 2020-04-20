package main

import "C"
import "fmt"

//本文件是为了 给c语言进行go代码的调用，用来制作静态库，动态库的示范文件
//export Echo
func Echo(content string) string {
	fmt.Println(content)
	return content
}

func main() {
}

//生成动态库 go build -o echo.so -buildmode=c-shared echo.go
//生成静态库 go build -buildmode=c-archive -o echo.a echo.go
