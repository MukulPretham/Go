package main

func startInc(init int) func(int) int{
	counter := init
	return func(i int) int {
		counter += i
		return counter
	}
} 

func fibonacci() func() int {
	n1 := 0
	n2 := 1
	called := 0
	return func() int {
		if called == 0{
			called++
			return 0
		}
		if called == 1{
			called++
			return 1
		}
		next := n1 + n2
		n1 = n2
		n2 = next
		return next
	}
}