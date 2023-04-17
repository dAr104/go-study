package main

import (
	"fmt"
	"runtime"
)

var a byte // alias for int8
var b rune // alias for int32 --> UTF8 4 byte for all character in the world

func main() {
	x := 42 // int 
	y := 42.3424 // float64
	fmt.Println(x)
	fmt.Println(y)
	fmt.Printf("%T\n", x)
	fmt.Printf("%T\n", y)
	fmt.Printf("%T\n", a)
	fmt.Printf("%T\n", b)

	fmt.Println(runtime.GOOS)
	fmt.Println(runtime.GOARCH)
}