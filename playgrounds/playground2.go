package main

type Pair struct{
	x int
	y int
}

func structonverter(i int, j int) Pair{
	var pairInstance Pair
	pairInstance.x = i
	pairInstance.y = j 
	return pairInstance
}

func nullify(p *Pair){
	(*p).x = 0
	(*p).y = 0
}