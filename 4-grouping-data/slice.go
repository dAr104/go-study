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
	fmt.Println(x[1])
	fmt.Println(len(x))


	// for looping on a slice use range
	for i, v := range x { 
		fmt.Println(i, v) // loop each element in the slice
	}

	// slicing a slice
	fmt.Println(x[0:2])
	fmt.Println(x[1:])
	fmt.Println(x[:4])

	// append value to slice: you can append a slice but you have to put elements
	x = append(x, 77, 88, 99, 1014)
	fmt.Println(x)

	y := []int{22, 34, 73, 33}
	x = append(x, y...)  // y... take all elements of slice y and put in the slice x 
	// ...T is variadic parameters aka illimitate
	fmt.Println(x)

	// delete from a slice
	x = append(x[:2], x[4:]...) 
	// take first and second element and append from to fifth to the end:
	// then delete thrid and fourth
	fmt.Println(x)

	// built-in function make: build a slice of certain number of elements 
	z := make([]int, 10, 11) // type: slice of int, len 10, capacity 100
	fmt.Println(z)
	fmt.Println(len(z))
	fmt.Println(cap(z))

	// difference beetwen lenght and capacity https://teivah.medium.com/slice-length-vs-capacity-in-go-af71a754b7d8
	/* The slice length is the number of available elements in the slice,
	whereas the slice capacity is the number of elements in the backing array.
	Adding an element to a full slice (length == capacity) leads to creating a new backing array with a new capacity,
	copying all the elements from the previous array, and updating the slice pointer to the new array.
	*/

	z = append(z, 333, 2)
	fmt.Println(z)
	fmt.Println(len(z))
	fmt.Println(cap(z))

	// when exceed the capacity go create runtime a new backend array with double capacity, and the slice point to it

	// multidimensional slice aka matrice
	jb := []string{"James", "Bond", "Mario", "Swim"}
	fmt.Println(jb)
	mb := []string{"Dario", "Andrea", "Pietro", "Nadia"}
	fmt.Println(mb)

	xp := [][]string{jb, mb} // matrice
	fmt.Println(xp)
}