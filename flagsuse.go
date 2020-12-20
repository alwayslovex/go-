package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
)

// 解析命令行的参数的包使用：参考文章 https://www.jianshu.com/p/f9cf46a4de0e

func main() {
	ip := flag.String("ip", "localhost", "input remote host !")
	port := flag.Int64("port", 6789, "input remote port!")
	intliststr := flag.String("intlist", "", "1,2,3")
	flag.Parse()
	if len(os.Args) < 3 {
		flag.Usage()
		return
	}
	fmt.Println("input:", *ip, *port, strings.Split(*intliststr, ","))

	fmt.Println("\n")
}
