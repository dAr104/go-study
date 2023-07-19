package main		

import "fmt"

func main() {
	var x [5]int // declare an array of 5 int, the lenght is part of the array's type
	fmt.Println(x)
	x[3] = 42
	fmt.Println(x)

	// get length of array
	fmt.Println(len(x))

	// Go incentive to use slices than arrays. Array are useful in allocation memory
	// They are a bulding block for slicess
}