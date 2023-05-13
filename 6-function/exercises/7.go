package main

import "fmt"

/*
Assign a func to a variable, then call that func
*/



func main() {
	x := func (x int) int {
		return x + x
	}

	fmt.Println(x(5))
}

