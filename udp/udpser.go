package main

import (
	"fmt"
	"net"
	"time"
)

func main() {
	addre, err := net.ResolveUDPAddr("udp", "localhost:7890")
	if err != nil {
		panic(err)
	}
	go clientReq()
	skt, err := net.ListenUDP("udp", addre)
	if err != nil {
		panic(err)
	}
	defer skt.Close()
	for {
		data := make([]byte, 10)
		n, addr, err2 := skt.ReadFromUDP(data)
		if err2 != nil {
			panic(err2)
		}
		if n > 0 {
			fmt.Printf("recv data length is %d,content is %s,address is %s,%d\n", n, string(data[0:n]), addr.IP, addr.Port)
			skt.WriteToUDP(data[:n], addr)
		}
	}
}

func clientReq() {
	time.Sleep(3 * time.Second)
	//conn,err := net.Dial("udp","localhost:7890")
	laddr, err := net.ResolveUDPAddr("udp", "localhost:6789")
	raddr, err := net.ResolveUDPAddr("udp", "localhost:7890")
	conn, err := net.DialUDP("udp", laddr, raddr)
	if err != nil {
		panic(err)
	}
	defer conn.Close()
	conn.Write([]byte("hello"))
	recv := make([]byte, 10)
	conn.Read(recv)
	fmt.Println(string(recv[0:5]))
}
