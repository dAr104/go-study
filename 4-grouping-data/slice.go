package main		

import "fmt"

// a SLICE allows you to group toghether values of the same type
// construct values for structs, arrays, slices, and maps

/*
The basic difference between a slice and an array is that a slice is a reference to a contiguous segment of an array.
Unlike an array, which is a value-type, slice is a reference type. A slice can be a complete array or a part of an array,
indicated by the start and end index.
In short, Go arrays are not flexible, and are not meant to be associated with dynamic memory allocation.
Slices, on the other hand, are abstractions built on top of these array types and are flexible and dynamic.
 */

func main() {
	// x := type{values} // composite literal
	x := []int{4,6,7,8,42} // slice
	fmt.Println(x)
	fmt.Println(len(x))


	// for looping on a slice use range
	for i, v := range x { 
		fmt.Println(i, v) // loop each element in the slice
	}
}