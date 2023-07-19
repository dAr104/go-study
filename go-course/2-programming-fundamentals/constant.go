package main

import "fmt"

const a = 42
const b = 42.43
const c = "James Bond"

// or

const (
	d int = 42
	e float64 = 43.33
	f string = "Dario"
)

func main() {
	fmt.Println(a, b, c)
	fmt.Printf("%T\n", a)
	fmt.Printf("%T\n", b)
	fmt.Printf("%T\n", c)
}