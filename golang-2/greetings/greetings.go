package greetings

import "fmt"

func Hello(name string) string{
	message := fmt.Sprintf("Hi %s",name) 
	return message
}