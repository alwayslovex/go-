package main

import (
	"fmt"
	"log"
	"net"
	"strconv"
	"time"
)

func main() {
	lister, err := net.Listen("tcp", "localhost:8090")
	if err != nil {
		log.Fatal("error : ", err)
	}
	//go clientRequest()
	for {
		conn, err := lister.Accept()
		if err != nil {
			log.Fatal(err)
			continue
		}
		go handleconn(conn)
	}
}

func handleconn(conn net.Conn) {
	defer conn.Close()
	for {
		body := make([]byte, 110) //注意这里，如果是body := make([]byte,0)那么在read的时候它不会自动增长，而是一直都是0。也就读不到东西。
		n, err := conn.Read(body)
		if err != nil {
			fmt.Print("conn read error : ", err)
			return
		}
		if n == 0 {
			continue
		}
		ss := string(body[:n])
		fmt.Printf("body %s: ,num : %d\n", ss, n)
		conn.Write(body[0:n])
	}
}

func clientRequest() {
	time.Sleep(3 * time.Second)
	for i := 0; i < 100; i++ {
		go doReq(i)
	}
}
func doReq(i int) {
	conn, err := net.Dial("tcp", "localhost:8090")
	if err != nil {
		fmt.Println(err)
	}
	conn.Write([]byte("hello" + strconv.Itoa(i)))
	var body [10]byte
	n, _ := conn.Read(body[:10])
	conn.Close()
	fmt.Println(string(body[:n]))
}
