package main

func test(a []int){
	a[0] = 1
	a = append(a,2)
}

func test2(a * []int){
	(*a)[0] =1
	*a = append(*a,2)
}

func main(){
	var tt  = make([]int,1)
	tt[0] = 0
	test(tt[:])
	test2(&tt)
}
