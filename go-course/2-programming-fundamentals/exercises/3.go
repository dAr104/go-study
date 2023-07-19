package main

import "fmt"

/*
Create TYPED and UNTYPED constants. Print the values of the constants
*/

const (
	a = 23
	b = 233.455
	c = "Ciao"
)

const (
	d int = 32
	e float64 = 4234.44
	f string = "buonasera"
)

func main() {

	fmt.Println(a, b, c)
	fmt.Println(d, e, f)

	
}