package main

import (
	"fmt"
	"runtime"
)

func main(){
	runtime.GOMAXPROCS(2)
	go func() {
		i := 0
		for i = 0;i < 100;i++{
			fmt.Println(i)
		}
	}()

	go func() {
		k := 100
		k *= 10
		fmt.Println("hello",k)
	}()
	fmt.Println("start")
	for i := 1000;i < 1020;i++{
		fmt.Println(i)
	}
	select{}
}
