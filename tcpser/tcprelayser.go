package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"strconv"
)

/*
tcp 包转发服务
将收到的tcp包转发到指定到端口和ip上去
v1.先完成简单的版本，不考虑其他，收到连接就建立一个线程，去处理，直到这个被关闭
v2.支持高并发，接收与回复
*/

var b int = 0

//获取ip和端口
func GetIpAndPort() (string, int64) {
	if b == 1 {
		return "127.0.0.1", 12345
	}
	return "127.0.0.1", 12346
}

//管理连接，
//1。接受包，进行转发
//2。收到回复，转发
func ConnSer(conn net.Conn) {
	defer conn.Close()
	b += 1
	body := make([]byte, 100)
	ip, port := GetIpAndPort()
	addr := ip + ":" + strconv.Itoa(int(port))
	trans, err := net.Dial("tcp", addr)
	if err != nil {
		log.Printf("trans body fail,%s", err.Error())
		return
	}
	defer trans.Close()

	for {
		n, err := conn.Read(body)
		if err != nil {
			log.Printf("read data from connection; error,%s", err.Error())
			break
		}
		if n > 0 {
			log.Printf("recv data from connnection;the length is %d bytes", n)
			//net.Dial()
			trans.Write(body[:n])
		}
	}
}

func ConnSerV2(conn net.Conn) {
	b += 1

	clientRead := make(chan []byte) //读取客户端的数据
	defer close(clientRead)

	servRead := make(chan []byte)
	defer close(servRead)

	close_notify := make(chan bool, 1)
	defer close(close_notify)

	addr, port := GetIpAndPort()
	addr += ":" + strconv.Itoa(int(port))
	serConn, err := net.Dial("tcp", addr)

	if err != nil {
		return
	}
	defer (conn).Close()
	defer serConn.Close()

	go HandleRead(&conn, clientRead, close_notify)
	go HandleRead(&serConn, servRead, close_notify)

	for {
		select {
		case cnt := <-close_notify:
			if cnt == true {
				fmt.Println("recv a close info")
				close_notify <- true
				return
			}
		case data := <-clientRead:
			serConn.Write(data)
		case data2 := <-servRead:
			conn.Write(data2)
		}
	}
}

func HandleRead(conn *net.Conn, readChan chan []byte, sign chan bool) {
	body := make([]byte, 1024)
	var exit bool = false
	for !exit {
		bodyLength, err := (*conn).Read(body)
		fmt.Println((*conn).RemoteAddr())
		if err != nil {
			if err == io.EOF {
				fmt.Println("read a close")
			}
			select {
			case _, ok := <-sign:
				if ok {
					sign <- true
					return
				} else {
					return
				}
			default:
				sign <- true
				return
			}
		} else if bodyLength > 0 {
			readChan <- body[:bodyLength]
		}
	}
}

//比较简单的版本，但是连接不能立马断掉
func ConnRes(conn net.Conn) {
	defer conn.Close()
	ip, port := GetIpAndPort()
	addr := ip + ":" + strconv.Itoa(int(port))
	trans, err := net.Dial("tcp", addr)
	if err != nil {
		log.Printf("trans body fail,%s", err.Error())
		return
	}
	defer trans.Close()
	ch := make(chan bool)
	defer close(ch)
	go func(cancle chan bool, conn net.Conn, trans net.Conn) {

		n, err := io.Copy(conn, trans)
		if err != nil {
			cancle <- true
		}
		fmt.Println(n)
	}(ch, conn, trans)

	go func(cancle chan bool, conn net.Conn, trans net.Conn) {

		n, err := io.Copy(conn, trans)
		if err != nil {
			cancle <- true
		}
		fmt.Println(n)
	}(ch, trans, conn)

	<-ch

}
func main() {
	listen, err := net.Listen("tcp", "localhost:8081")
	if err != nil {
		log.Fatal("create a listen socket fail : %s", err.Error())
	}

	defer listen.Close()

	for {
		conn, err := listen.Accept()
		if err != nil {
			log.Printf("accept a tcp connection fail")
			continue
		}
		go ConnSerV2(conn)
		//go ConnRes(conn)
	}
}
