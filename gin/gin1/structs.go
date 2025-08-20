package main

import "gorm.io/gorm"

type Book struct{
	gorm.Model
	ID uint
	Title string
	Author string
	Price float64
}

type UpdateReq struct{
	attribute string
	value string
}