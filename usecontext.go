package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

/*主要是练习使用下context这个包，
主要的使用场景举例，比如http请求，中需要去使用一个新的goroutine 去处理一个比较耗时的操作时，
如果我们需要设置超时并取消这些goroutine，那么除了自己封装，还可以使用context包。下面是练习使用
*/

var n = 0

func Count(ctx context.Context, wg *sync.WaitGroup) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("done")
			wg.Done()
			return
		default:
			n++
			time.Sleep(time.Second)
		}
	}
}

func main() {
	wg := sync.WaitGroup{}
	ctx, cancelfunc := context.WithCancel(context.Background())
	wg.Add(1)
	go Count(ctx, &wg)
	time.Sleep(1 * time.Second)
	cancelfunc()
	fmt.Println(n)

	ctx2, canc := context.WithTimeout(context.Background(), time.Second*10) //这里的是设置了一个自动超时的context，如果超时了，则自动调用cancel
	wg.Add(1)
	go Count(ctx2, &wg)
	canc()
	wg.Wait()
}
