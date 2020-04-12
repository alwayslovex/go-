package main

import (
	"fmt"
	"sync/atomic"
)

//练习使用原子操作

func main() {

	var age int32 = 18
	atomic.AddInt32(&age, 1)

	var num uint64 = 199
	atomic.AddUint64(&num, 299)

	fmt.Println(age, num)
	var zero = 0
	//原子减法
	atomic.AddUint64(&num, ^uint64(zero-1)) //对于无符号的数，进行减法时需要做^uint64(待减数-1)
	fmt.Println(num)

	//v:= atomic.LoadInt32(&age) 原子读取一个数的值
	//atomic.StoreInt32(&age,10)原子地写入一个值
	atomic.SwapInt32(&age, 10) //原子交换
	atomic.CompareAndSwapInt32(&age, 3, 0)
}
