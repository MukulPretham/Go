package main

import "fmt"

import "examples/mod1"

func main(){
	fmt.Print("hello")
	fmt.Println(mod1.Hello("World"))
	mod1.Hello("Mukul")
}



