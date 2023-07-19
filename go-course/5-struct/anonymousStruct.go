package main

import "fmt"


func main(){

	// not create the type person. This is anonymous struct
	p1 := struct {
		first string
		last string
		age int
	}{
		first: "Dario",
		last: "Rossi",
		age: 28,
	}

	fmt.Println(p1)
	
}