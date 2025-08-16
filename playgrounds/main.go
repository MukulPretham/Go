package main

import (
	"fmt"
	"time"
)

//pg1
// func main(){
// 	nums := []int{
// 		2,3,4,5,6,
// 	}
// 	var p *[]int
// 	p = &nums
// 	ReverseArray(p)
// 	fmt.Print(nums)
// }

//pg2
// func main(){
// 	pair := structonverter(2,3)
// 	p := &pair
// 	nullify(p)
// 	fmt.Print(*p)
// }

//pg3
func main(){
	start := time.Now()
	ch := make(chan int)
	go sum([]int{1,2,3,4,5,34,345,345,34,3455,4,24,435,34,34,345,345,34,},ch)
	go sum([]int{1,2,3,4,5,34,43,34,34,4,4,546,456,5,4,5,4,5,4,435,435,345},ch)
	sum1, sum2 := <-ch,<-ch
	fmt.Print("calcullation time : ",time.Since(start))
	fmt.Print(sum1+sum2)
}

