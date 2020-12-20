package main

import (
	"fmt"
	"math"
)

func main() {
	var f = 2.55
	fmt.Println(f * 100)
	s := fmt.Sprintf("%.2f", f*100)
	fmt.Println(s, int64(math.Ceil(f*100)))

}
