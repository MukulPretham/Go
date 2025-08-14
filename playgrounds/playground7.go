package main

// import "fmt"

type arrayList struct{
	arr []int
}

func (a *arrayList)cons(){
	a.arr = make([]int, 0)
}

func (a*arrayList)appendVal(val int){
	a.arr = append(a.arr, val)
}

func appendToList(a * arrayList,va int){
	a.arr = append(a.arr, va)
}

