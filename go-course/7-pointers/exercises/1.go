package main

import "fmt"

/*
Create a value and assign it to a variable.
Print the address of that value.
*/

func main() {
	x := 43
	fmt.Println(&x)
}