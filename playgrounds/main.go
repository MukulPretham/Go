package main

import "fmt"

// func main(){
// 	nums := []int{
// 		2,3,4,5,6,
// 	}
// 	var p *[]int
// 	p = &nums
// 	ReverseArray(p)
// 	fmt.Print(nums)
// }

func main(){
	pair := structonverter(2,3)
	p := &pair
	nullify(p)
	fmt.Print(*p)
}

// func main(){
// 	var x int = 1
// 	var p *int = &x
// 	fmt.Println(x)
// 	fmt.Println(p)
// 	fmt.Println(*p)
// }