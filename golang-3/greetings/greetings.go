package greetings

import ("fmt"
 "errors")

func Hello(name string) (string,error){
	if name==""{
		return "",errors.New("empty string")
	}
	message := fmt.Sprintf("Hi %s",name) 
	return message,nil
}