package main

import "fmt"

/*
create a func with the identifier foo that
	○ takes in a variadic parameter of type int
	○ pass in a value of type []int into your func (unfurl the []int)
	○ returns the sum of all values of type int passed in
create a func with the identifier bar that
	○ takes in a parameter of type []int
	○ returns the sum of all values of type int passed in
*/

func main() {
	slc := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	sum := foo(slc...)
	fmt.Println(sum)
	s := bar(slc)
	fmt.Println(s)
}

func foo(x ...int) int {
	sum := 0
	for _, v := range x {
		sum += v
	}
	return sum
}

func bar(x []int) int {
	sum := 0
	for _, v := range x {
		sum += v
	}
	return sum
}