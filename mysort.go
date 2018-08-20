package main
//排序练习
import (
	"fmt"
	"bufio"
	"os"
	"strconv"
)

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

func quick_sort(arr []int){
  //undo
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
	arr_int := input_nums()
	print_arr(arr_int)
	fmt.Printf("---------------\n")
	//bubble_sort(arr_int)
	select_sort(arr_int)
	print_arr(arr_int)
}
