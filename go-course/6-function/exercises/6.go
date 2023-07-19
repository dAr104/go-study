package main

import "fmt"

/*
Build and use an anonymous func
*/



func main() {
	x := func (x int) int {
		return x + x
	}(2)

	
	fmt.Println(x)
}

