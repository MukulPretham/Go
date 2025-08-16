package main

// Sum of 2 sumations

// go routines
//Channels
// Buffered channels
func sum(nums []int,ch chan int){
	sum := 0
	for vsl,_ := range nums{
		sum += vsl
	}
	ch <- sum
}