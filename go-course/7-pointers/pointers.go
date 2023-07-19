package main

import "fmt"

func main() {
	a := 42
	fmt.Println(a)
	fmt.Println(&a) // & gives you the address

	fmt.Printf("%T\n", a)
	fmt.Printf("%T\n", &a) // the type is *int, a pointer to a int

	var b *int = &a
	fmt.Println(b)
	fmt.Println(*b) // * give the value stored in that address
}


// We se that we can attach a method to a type. Those methode attached to a type are known as it's method set