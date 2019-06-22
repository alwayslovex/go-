### go 中的 waitgroup 怎么用

* 首先说个问题，那就是使用的时候遇到的坑
```
var n = 0

func Count(ctx context.Context,wg sync.WaitGroup){
	for{
		select {
		case <- ctx.Done():
			fmt.Println("done")
			wg.Done()
			return
		default:
			n++
			time.Sleep(time.Second)
		}
	}
}

func main(){
	wg := sync.WaitGroup{}
	ctx,cancelfunc := context.WithCancel(context.Background())
	wg.Add(1)
	go Count(ctx,wg)

	time.Sleep(10*time.Second)
	cancelfunc()
	fmt.Println(n)
	wg.Wait()
}

这个代码看上取没什么问题，但是在运行的时候会报错：error: all goroutines are asleep - deadlock!，说句人话，就是死锁了，程序会一直在wait中等待，但是没有协程了。

后来我仔细一看发现，go中是值传递，所以在Count函数中应该进行地址传递。
Count(ctx context.Context,wg * sync.WaitGroup)
```

