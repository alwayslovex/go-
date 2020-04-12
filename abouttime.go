package main

import (
	"fmt"
	"time"
)

func main() {
	var t = time.Now()
	fmt.Printf("now = %s\n", time.Now())
	fmt.Println(t.Unix())
	fmt.Println(t.Second())
	fmt.Println(t.Nanosecond())
	fmt.Println(t.UnixNano())
	//unix timestamp --> date
	var unixtimestamp = t.Unix()
	tm2 := time.Unix(unixtimestamp, 0)
	fmt.Println(tm2.Format("2006-01-02 03:04:05"))
	fmt.Println(time.Now().Format("20060102"))

	//date --> unix timestamp

	t, _ = time.Parse("2006-01-02-03", "2018-08-22-09")
	fmt.Println(t)
}
