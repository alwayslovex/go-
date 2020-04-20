package main

/*
#include "add.h"
*/
import "C"
import (
	"fmt"
	"unsafe"
)

/*
go 中可以直接调用C的代码，但是格式是有要求的。C的代码要写在注释里。
并且注释完要紧跟着import "C" 不能有空格否则会报错could not determine kind of name for C.
必须给 import "C" 单独一行，且必须放在注释的 C 代码后面一行。
*/
func main() {
	var a = 10
	var b = 20

	fmt.Println(int(C.sub(C.int(a), C.int(b))))

	var hello = "hello world"
	cs := C.CString(hello)
	defer C.free(unsafe.Pointer(cs))
	fmt.Println(C.GoString(C.echo(cs)))

	//使用C中的结构体
	var args C.struct_AddArgs
	args.a = 1
	args.b = 2

	fmt.Println(int(C.newadd(args)))
}
