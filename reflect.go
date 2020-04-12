package main

import (
	"fmt"
	"reflect"
)

func main() {
	var x interface{}

	x = 1

	ty := reflect.TypeOf(x)

	fmt.Println(ty.String())

	tv := reflect.ValueOf(x)

	switch tv.Kind() {
	case reflect.String:
		fmt.Println("string")

	case reflect.Int:
		fmt.Println("int")

	}
}
