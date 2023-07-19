package main

import "fmt"

/*
Closure: when enclosing a variable so that we limi its scope
*/

func main() {
	
	{
		y := 42
		fmt.Println(y)
	}
	// fmt.Println(y) don't works

	a := incrementor()
	b := incrementor() 
 
	fmt.Println(a())
	fmt.Println(a())
	fmt.Println(a())
	fmt.Println(a())
	fmt.Println(b())
	fmt.Println(b())
}

func incrementor() func() int {
	var x int
	// the function returned is enclosing in the scope of x
	return func() int {
		x++
		return x
	}
}
