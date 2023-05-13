package main

import "fmt"

/*
● Create a func which returns a func
● assign the returned func to a variable
● call the returned func
*/



func main() {
	f := func () func(x int) int{
		return func (x int) int {
			return x + x
		}
	}

	result := f()

	fmt.Println(result(5))
}

