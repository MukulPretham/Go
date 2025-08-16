package main

type Image interface{
	Bounds() int
}

type Rect struct{
	x int
	y int
}

type Circle struct{
	r int
}

func (r Rect) Bounds()int{
	return r.x * r.y
}

func (r Circle) Bounds()int{
	return r.r
}

