package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
)

type Test1 struct {
	Num  int
	Name [3]byte
}
type Test2 struct {
	Num  int32
	Name [3]byte
}

/*这个是对go中对二进制序列化对使用。其中遇到了一个坑。我开始用了int类型作为结构体变量因此不能进行正常的序列化。


1.go中提供了两个函数来进行序列化和反序列化 ，分别是binary.read   binary.write
2.以下就是对这两个函数对使用。
3.对于write 来讲，我们可以序列化一个结构体，如上面对test，但是要确保这个结构体能正确计算出大小。
因此这个函数就要求他对序列化对象必须是一个定长的结构，所以结构中不能包含，切片（slice），string，int等非定长对类型。
可以是一个值，或者是它的指针，也可以是一个定长的slice 例如：make([]byte,10)。
还有一点需要注意对是，结构体成员的大小写问题，因为小写是不可导出的所以需要大写,还有是不能使用_.
*/
func main() {

	var a Test1
	a.Num = 1
	a.Name[0] = 'a'
	a.Name[1] = 'b'
	a.Name[2] = 'c'

	var b Test2
	b.Num = 1
	b.Name[0] = 'a'
	b.Name[1] = 'b'
	b.Name[2] = 'c'

	fmt.Println(binary.Size(a))
	fmt.Println(binary.Size(b))
	tmp := new(bytes.Buffer)
	err := binary.Write(tmp, binary.LittleEndian, &a)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(tmp.Len())
	fmt.Println(tmp.Bytes())

	err = binary.Write(tmp, binary.LittleEndian, b)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(tmp.Len())
	fmt.Println(tmp.Bytes())

	//tmp2 := new(bytes.Reader)
	var c Test2
	err = binary.Read(tmp, binary.LittleEndian, &c)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println("read :")
	fmt.Println(c.Num)
	fmt.Println(c.Name)
}
