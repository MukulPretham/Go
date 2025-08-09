package main

import (
	"examples3/greetings"
	"fmt"
	"log"
)

func main(){
	message,err := greetings.Hello("")
	if(err != nil){
		log.Fatal(err)
	}
	fmt.Print(message)
}