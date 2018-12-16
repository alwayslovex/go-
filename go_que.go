package main

import (
	"fmt"
	"sync/atomic"
	"time"
)

/*
1.简单的无锁的循环队列
实现多线程队列，这里使用了cas。
2.多此一举，channel很好用。
*/

type MyQue struct {
	que  []interface{}
	tail uint64
	head uint64
	size uint64
}

func (m *MyQue) Init(queSize uint64) {
	m.size = 1 << queSize
	m.que = make([]interface{}, m.size)
	m.tail = 0
	m.head = 1
}
func (m *MyQue) Put(elem interface{}) bool {
	var ret bool = true
	head := m.head
	pos := (head + 1) & (m.size - 1)
	fmt.Printf("put pos := %d\n", pos)
	//判断是否已经满了。nil代表空值
	if m.que[head] != nil {
		return false
	}

	for !atomic.CompareAndSwapUint64(&m.head, head, pos) { //使用cas进行入队列
		head = m.head
		pos = head&(m.size-1) + 1
	}

	m.que[head] = elem
	return ret
}

func (m *MyQue) Get() (bool, interface{}) {
	tail := m.tail
	pos := (tail + 1) & (m.size - 1)

	fmt.Printf("tail = %d,pos = %d\n", tail, pos)

	if m.que[pos] == nil { //full
		return false, nil
	}

	for !atomic.CompareAndSwapUint64(&m.tail, tail, pos) {
		tail = m.tail
		pos = tail&(m.size-1) + 1
	}
	res := m.que[pos]
	m.que[pos] = nil
	return true, res
}

var cirQue MyQue
var fin chan int

func Push() {
	for i := 0; i < 1000000; i++ {
		if cirQue.Put(i) {
			continue
		} else {
			time.Sleep(time.Second * 1)
		}
	}
	fin <- 1
}
func Get() {
	for j := 0; j < 1000000; j++ {
		suc, elem := cirQue.Get()
		if suc {
			fmt.Println(elem)
			continue
		}
	}
	fin <- 1
}

var que chan int

func channelPut() {

	for i := 0; i < 1000000; i++ {
		que <- i
	}
}
func channelGet() {
	for j := 0; j < 1000000; j++ {
		el := <-que
		fmt.Printf("chan:%d\n", el)
	}
	fin <- 1
}

func main() {
	cirQue.Init(2)
	fin = make(chan int)
	que = make(chan int, 4)
	go Push()
	go Get()
	go Get()

	//go channelPut()
	//go channelGet()

	defer close(fin)
	defer close(que)
	var kk int
	for {
		select {
		case k := <-fin:
			kk += k
		}

		if kk >= 2 {
			break
		}
	}
}
