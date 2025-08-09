package main

import (
	"examples3/greetings"
	"fmt"
	"log"
)

func main(){
	names := []string{
		"Mukul",
		"Pretham",
		"Munna",
	}
	message,err := greetings.Hellos(names)
	if(err != nil){
		log.Fatal(err)
	}
	fmt.Print(message)
	for key, value := range message{
		fmt.Printf("%s ==> %s \n",key,value)
	}
}