package main

import (
	"fmt"
)

func main(){
	arr1 := []int{1,2,3}
	arr2 := []int{1,2,3,2,2,3,2,23,32,32,32,12,32,43,2,2}

	ch := make(chan int)

	go sum(arr1,ch)
	go sum(arr2,ch)

	x := <- ch
	y := <- ch

	fmt.Print(x)
	fmt.Print(y)
}

func sum(nums []int,ch chan int){
	sum := 0
	for _,vsl := range nums{
		sum += vsl
	}
	ch <- sum
}