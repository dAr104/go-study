package main

import "fmt"


func main() {

	x := "Hello World"
	fmt.Println(x)
	fmt.Printf("%T\n", x)
	bx := []byte(x)
	fmt.Println(bx)
	fmt.Printf("%T\n", bx)

	for i := 0; i < len(x); i++ {
		fmt.Printf("%#U ", x[i]) // %#U prints the unicode character
	}

	fmt.Println("")

	for i, v := range x { // range returns the index and the value
		fmt.Printf("at index position %d we have hex %#x\n", i,v)
	}	
}