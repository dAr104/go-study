package main

import "fmt"

func main() {
	// func (r receiver) identifier(parameters) (return(s)) { ... }
	foo()
	bar("James")
	fmt.Println(woo("Dario"))
	x, y := mouse("Ian","Fleming")
	fmt.Println(x)
	fmt.Println(y)
}

func foo(){
	fmt.Println("Hello from foo")
}

// everythng in go is pass by value
func bar(s string) {
	fmt.Println("Hello,", s)
}

// return a string
func woo(s string) string {
	return fmt.Sprint("Hello from woo,", s)
}

func mouse(fn string, ln string) (string, bool) {
	a := fmt.Sprint(fn, " ", ln, `, says "Hello"`)
	b := false
	return  a, b
}