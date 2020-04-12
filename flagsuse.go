package main

import (
	"flag"
	"fmt"
	"os"
)

//解析命令行的参数的包使用：参考文章https://www.jianshu.com/p/f9cf46a4de0e

func main() {
	ip := flag.String("ip", "localhost", "input remote host !")
	port := flag.Int64("port", 6789, "input remote port!")

	flag.Parse()
	if len(os.Args) < 3 {
		flag.Usage()
		return
	}
	fmt.Println("input:", *ip, *port)

	fmt.Println("\n")
}
