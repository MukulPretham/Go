package main

import "fmt"

type Pair struct{
	x int
	y int
}

func structonverter(i int, j int) Pair{
	var pairInstance Pair
	pairInstance.x = i
	pairInstance.y = j 
	pairInstance = Pair{i,j}
	return pairInstance
}

func nullify(p *Pair){
	fmt.Print(*p)
	(*p).x = 0
	(*p).y = 0
}