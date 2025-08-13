package main

import "fmt"

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
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}

