package main

import "fmt"

/*
Recursion: a function that call itself several times 
*/

func main() {
	n := factioral(10)
	fmt.Println(n)

	n1 := factorialLoop(10)
	fmt.Println(n1)
	
}

func factioral(n int) int {
	if n == 0{
		return 1
	}
	return n * factioral(n-1)
}

func factorialLoop(n int) int {
	x := 1
	for ; n > 0; n-- {
		x *= n
	}
	return x
}