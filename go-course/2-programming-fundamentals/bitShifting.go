package main

import "fmt"

/*
	kb := 1024
	mb := 1024 * kb
	gb := 1024 * mb
*/

const (
	_ = iota // skip 0 value
	kb = 1 << (iota * 10)
	mb = 1 << (iota * 10)
	gb = 1 << (iota * 10)

)


func main() {
	x := 4
	fmt.Printf("%d\t\t%b\n", x, x)

	y := x << 1 //left shifting -> shift the bit of one position to left
	fmt.Printf("%d\t\t%b\n", y, y)

	// shift of 10 digits in binary
	fmt.Printf("%d\t\t\t\t%b\n", kb, kb)
	fmt.Printf("%d\t\t\t\t%b\n", mb, mb)
	fmt.Printf("%d\t\t\t%b\n", gb, gb)
}