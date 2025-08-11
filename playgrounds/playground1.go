package main

func ReverseArray(arr *[]int) {
	i := 0
	j := len(*arr) - 1
	var temp int = 0
	for i <= j {
		temp = (*arr)[i]
		(*arr)[i] = (*arr)[j]
		(*arr)[j] = temp
		i++
		j--
	} 
}