package main

import (
	"fmt"
	"log"
	"net"
)

func main() {
	lister, err := net.Listen("tcp", "localhost:8090")
	if err != nil {
		log.Fatal("error : ", err)
	}
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
		body := make([]byte, 110)

		n, err := conn.Read(body)
		if err != nil {
			fmt.Print("conn read error : ", err)
			return
		}
		if n == 0 {
			continue
		}
		ss := string(body[:n])
		fmt.Printf("body %s: ,num : %d", ss, n)
		conn.Write(body)
	}
}
