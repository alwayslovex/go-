### 笔记1.map,slice,channel等在函数间传递的是值还是地址?

在go中,所有的都是值引用,也就是传值.如果需要进行传引用,要进行取地址传递.
但是在我使用的时候,发现

```
func test(kv map[string]string){
	fmt.Println(kv["he"])  //打印not me
	kv["he"] = "me"
}

func main(){
	mmp := make(map[string]string,0)
	mmp["he"] = "not  me"
	test(mmp)
	fmt.Println(mmp["he"])  //打印me
}
```
我一开始以为是不会更改map中的内容...但是没成想改了..于是开始查询语法.发现,在make产生map的时候,是产生的一个指针,也就是说,
var mmp * map[string]string = make(map[string]string,0)
所以其实传递的是一个指针...同样的,slice ,channel 通过make产生的也是一个指针..在函数间传递的也是指针..
不过slice还是特殊一点.
其实用c来描述slice会比较清楚.

```
struct slice {
	int * data;
	int len;
	int cap;
};

void test(struct slice arr)
{
	arr.data[0] = 1;//这里能够修改,并且影响外部.
	arr.len  = 10;//但是这里的修改是不会影响外部.
}

int main()
{

	struct slice t;
	int a = 1;
	t.data = &a;
	t.len = 1;
	t.cap = 2;
	test(t);
	return 0;
}
```
