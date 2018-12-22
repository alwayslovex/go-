package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int, 10)
	defer close(ch)
	go produce(ch, 2)
	go produce(ch, 1)
	consumer(ch)
	time.Sleep(4 * time.Second)

}

func produce(ch chan int, i int) {
	for {
		ch <- 1
		fmt.Printf("insert %d\n", i)
		time.Sleep(1 * time.Second)
	}
}
func consumer(ch chan int) {
	for {
		var a int
		a = <-ch
		fmt.Println(a)
	}

}
