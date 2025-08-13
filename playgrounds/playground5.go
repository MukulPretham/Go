package main

import "fmt"

//Maps

type vertex struct {
	lat  float64
	long float64
}

func MapExe1() {
	// Initilizing an array
	m := make(map[string]string)
	// Insert an array
	m["1"] = "A"
	fmt.Print("map after inserrtion: ", m)
	// delete an array
	delete(m, "1")
	fmt.Print("map after deletion: ", m)
}

func WordCount(s string) map[string]int {
	mymap := make(map[string]int)
	curr := string(s[0])
	for i := 1; i < len(s); i++ {
		if string(s[i]) == " " || i == len(s)-1  {
			if i == len(s)-1{
				curr = curr + string(s[i])
			} 
			_, status := mymap[curr]

			if status == false {
				mymap[curr] = 0
			}
			mymap[curr] = mymap[curr] + 1
			curr = ""
		} else {
			curr = curr + string(s[i])
		}
	}
	return mymap
}
