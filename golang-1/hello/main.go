package main

import ("fmt"
	"examples/greetings"
)

func main() {
	message := greetings.Hello("mukul")
	fmt.Println(message)
}