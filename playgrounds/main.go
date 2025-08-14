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
	list := arrayList{}
	list.cons()
	// list.arr = append(list.arr, 1)
	list.appendVal(99)
	fmt.Print(list.arr)
	appendToList(&list,1)
	appendToList(&list,10)
	fmt.Print(list.arr)
	fmt.Print("length:",len(list.arr))
}

