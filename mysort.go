package main
//排序练习
import (
	"fmt"
	"bufio"
	"os"
	"strconv"
)

type myStack struct {
	stack []int
	stack_end int
}

func (s * myStack)Push(elem int){
	s.stack = append(s.stack,elem)
	s.stack_end++
}
func (s * myStack)Pop() int{
	s.stack_end--
	var ret = s.stack[s.stack_end]
	s.stack = s.stack[0:len(s.stack)-1]
	return ret
}
func (s * myStack)Empty() bool{
	if s.stack_end == 0{
		return true
	}
	return false
}

func bubble_sort(arr []int){
	length := len(arr)
	for index := 0 ;index < length;index++{
		pos := 0
		for index2 := 1 ; index2 < length-index;index2++{
			if arr[pos] > arr[index2]{
				arr[index2],arr[pos] = arr[pos],arr[index2]
				pos = index2
			}else {
				pos = index2
			}
		}
	}

}
func select_sort(arr []int){
	for index,_ := range arr{
		min_pos := index
		for index2:=index+1;index2 < len(arr) ;index2++  {
			if arr[min_pos] > arr[index2] {
				min_pos = index2
			}
		}
		if min_pos != index{
			arr[min_pos] ,arr[index] = arr[index],arr[min_pos]
		}
	}
}


//快速排序,递归方式
func quick_sort(arr []int){
	size := len(arr)
	var left = 0
	var right = size - 1
	if left >= right{
		return
	}

	var base = left
	for left < right{
		if arr[right] >= arr[base]{
			right--
			continue
		}
		if arr[left] <= arr[base]{
			left++
			continue
		}

		arr[left],arr[right] = arr[right],arr[left]
	}
	if	arr[base] > arr[left]{
		arr[base],arr[left] = arr[left],arr[base]
	}

	if base < left{
		quick_sort(arr[base:left])
	}
	if size > left{
		quick_sort(arr[left+1:])
	}
}

//迭代方式
func quick_sortv2(arr []int){
	size := len(arr)
	if size == 0{
		return
	}
	statckInt := myStack{}
	statckInt.Push(0)
	statckInt.Push(size-1)

	for !statckInt.Empty(){
		var right = statckInt.Pop()
		var left = statckInt.Pop()
		if left >= right{
			continue
		}

		var base = left
		for left < right{
			if arr[right] >= arr[base]{
				right--
				continue
			}
			if arr[left] <= arr[base]{
				left++
				continue
			}

			arr[left],arr[right] = arr[right],arr[left]
		}
		if	arr[base] > arr[left]{
			arr[base],arr[left] = arr[left],arr[base]
		}

		//base = left
		if base < left{
			statckInt.Push(base)
			statckInt.Push(left-1)
		}
		if size > left{
			statckInt.Push(left+1)
			statckInt.Push(size-1)
		}
	}

}


func print_arr(arr []int){
	for _,elem := range arr{
		fmt.Printf("elem = %d\n",elem)
	}
}

func input_nums() (arr []int){
	input := bufio.NewScanner(os.Stdin)
	index := 0
	arr = make([]int,0)
	for input.Scan(){
		a := input.Text()
		//fmt.Print(a)
		var err error
		num,err := strconv.Atoi(a)
		if(err !=nil){
		}
		arr  = append(arr,num)
		index++
	}
	return arr
}

func main() {
	//arr_int := []int{4,1,2,5,0,3,9,7}
	arr_int := input_nums()//ctrl + D end input
	print_arr(arr_int)
	fmt.Printf("---------------\n")
	//bubble_sort(arr_int)
	quick_sortv2(arr_int)
	print_arr(arr_int)
}
