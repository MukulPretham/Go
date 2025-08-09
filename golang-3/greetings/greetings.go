package greetings

import (
	"errors"
	"fmt"
	"math/rand"
)

func Hello(name string) (string,error){
	if name==""{
		return "",errors.New("empty string")
	}
	message := fmt.Sprintf(randomGreeting(),name) 
	return message,nil
}

func randomGreeting()string{
	greetings:=[]string{
		"Hi %v",
		"Hello %v",
		"welcome %v",
	}
	return greetings[rand.Intn(len(greetings))]
}

func Hellos(names[]string) (map[string]string,error){
	messages := map[string]string{}
	for _,name := range names{
		currMsg,err := Hello(name)
		if err != nil {
            return nil, err
        }
		messages[name] = currMsg
	} 
	return messages,nil
}