package main

import (
	"fmt"
	"net"
	"strconv"
	"strings"
)

//数字与字符串间的互相转换
//还有增加了两个拷贝的ip转换的函数

func main() {
	snum := "100"
	num, err := strconv.ParseInt(snum, 10, 32) //这个只会将结果转换成int64类型，无论bitsize填写32还是64。。所以如果需要转换int，再做强转（）
	fmt.Println(int(num))                      //强转成int
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(strconv.Atoi("20"))
	nsnum := strconv.Itoa(10) //其实底层调用了下面的函数，将数字转换成字符串。
	nsnum = strconv.FormatInt(1000, 10)
	fmt.Println(nsnum)
	fmt.Println(snum, num)

	//string到float32(float64)
	f, err := strconv.ParseFloat("3.14", 32/64)
	//float到string
	s := strconv.FormatFloat(3.14, 'E', -1, 32)
	s2 := strconv.FormatFloat(4.28, 'E', -1, 64)
	// 'b' (-ddddp±ddd，二进制指数)
	// 'e' (-d.dddde±dd，十进制指数)
	// 'E' (-d.ddddE±dd，十进制指数)
	// 'f' (-ddd.dddd，没有指数)
	// 'g' ('e':大指数，'f':其它情况)
	// 'G' ('E':大指数，'f':其它情况)
	fmt.Println(f, s, s2)
}

func IPToUInt32(ipnr net.IP) uint32 {

	bits := strings.Split(ipnr.String(), ".")

	b0, _ := strconv.Atoi(bits[0])
	b1, _ := strconv.Atoi(bits[1])
	b2, _ := strconv.Atoi(bits[2])
	b3, _ := strconv.Atoi(bits[3])

	var sum uint32
	sum += uint32(b0) << 24
	sum += uint32(b1) << 16
	sum += uint32(b2) << 8
	sum += uint32(b3)
	return sum
}

func UInt32ToIP(intIP uint32) net.IP {
	var bytes [4]byte
	bytes[0] = byte(intIP & 0xFF)
	bytes[1] = byte((intIP >> 8) & 0xFF)
	bytes[2] = byte((intIP >> 16) & 0xFF)
	bytes[3] = byte((intIP >> 24) & 0xFF)

	return net.IPv4(bytes[3], bytes[2], bytes[1], bytes[0])
}
