package main

import (
	"fmt"
	"runtime"
)

//捕获panic异常,会打印出异常处的函数名和错误信息
func Catch() {
	if err := recover(); err != nil {
		if v, ok := err.(error); ok {
			pc, _, _, _ := runtime.Caller(3)
			funName := runtime.FuncForPC(pc)
			fmt.Println(funName.Name(),v)
		}
	}
}

func Panic(){
	var a map[int]int
	defer Catch()

	a[1] = 1
}
func main(){
	Panic()
}