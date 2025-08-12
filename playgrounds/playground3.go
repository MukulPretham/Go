package main

import (
	"fmt"
)

// import "fmt"

//Arrays and Slices

func ArrayExe()([2]int){
	arr := [2]int{0,1,}
	return arr
}

//Nill slice
func ArrayExe2()([]int){
	arr := []int{0,1,}
	return arr
}

//Slice Example
func SliceExe()[]int{
	arr := [5]int{1,2,3,4,5}
	slice := arr[0:2]
	slice[0] = 99
	return slice
}
// Note: slice of an array, is a reference to the original array.

func MakeSliceExe()([]int,[]int){
	a := make([]int,5)
	// What is happening
	// - a array of size 5 is created and all the numbers are initially 0 

	b := make([]int, 0, 5)
	fmt.Println("size or length: ",len(b))
	fmt.Println("capacity",cap(b))
	// here the 3dh argument is capacity of the array, under the hood, go lang creates a array
	// of 5 and slice size = 0. So now the reallocation happens only id we add more then 5
	// elements in the slice
	return a,b
}

func AppemdToSliceExe() []int{
	a := []int{1,2,3}
	a = append(a,2)
	return a
}

func Pic(dx, dy int) [][]uint8 {
	var mat [][]uint8
	for i:=0 ; i<dx; i++{
		var curr []uint8
		for j:=0 ; j<dy ; j++{
			curr = append(curr, uint8(j))
		}
		mat = append(mat, curr)
	}
	return mat
}



