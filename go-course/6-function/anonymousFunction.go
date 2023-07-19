package main

import "fmt"

func main() {
	foo()

	// anonymou fuunction no have identifier
	func(){
		fmt.Println("Anonymous func ran")
	}()

	func(x int){
		fmt.Println("Anonymous func ran with argument", x)
	}(42)
}

func foo(){
	fmt.Println("Hello from foo")
}
