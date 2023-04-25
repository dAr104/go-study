package main

import "fmt"

/*
● Using a COMPOSITE LITERAL:
	● create a SLICE of TYPE int
	● assign 10 VALUES
● Range over the slice and print the values out.
● Using format printing
	○ print out the TYPE of the slice
*/

func main() {

	x := []int{23, 44, 42, 27, 8, 55, 63, 7, 11, 223}

	for i, v := range x {
		fmt.Println(i, v)
	}

	fmt.Printf("%T\n", x)
}