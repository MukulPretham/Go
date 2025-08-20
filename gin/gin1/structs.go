package main


type Book struct{
	ID uint
	Title string
	Author string
	Price float64
}

type UpdateReq struct{
	Attribute string 
	Value string 
}