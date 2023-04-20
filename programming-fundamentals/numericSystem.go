package main

import "fmt"

func main() {
	s := "H"
	fmt.Println(s)

	bs := []byte(s)
	fmt.Println(bs)

	n := bs[0]
	fmt.Println(n)
	fmt.Printf("%T\n", n)
	fmt.Printf("%b\n", n) // binary
	fmt.Printf("%x\n", n) // hex
	fmt.Printf("%#x\n", n) // # put the 0x before the number
}