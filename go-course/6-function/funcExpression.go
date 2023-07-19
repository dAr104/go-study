package main

import "fmt"

func main() {

	// function can be stored in variable, are "first class citizen"
	f := func(){
		fmt.Println("my first func expression")
	}

	g := func(x int, y int){
		fmt.Println("Sum of", x, "and", y, "is", x+y)
	}

	f()
	g(2, 5)
}
