package main	

import "fmt"

/*
Create your own type “person” which will have an underlying type of “struct” so that it can store
the following data:
	● first name
	● last name
	● favorite ice cream flavors
Create two VALUES of TYPE person. Print out the values, ranging over the elements in the slice
which stores the favorite flavors.
*/

type person struct {
	first string
	last string
	iceCreamFlavors []string
}

func main() {

	p1 := person{
		first: "Dario",
		last: "Rossi",
		iceCreamFlavors: []string{"Cocco", "Tiramisù"},
	}

	p2 := person{
		first: "Andrea",
		last: "Espposito",
		iceCreamFlavors: []string{"Stracciatella", "Limone", "Fragola"},
	}

	fmt.Println(p1.first, p1.last)
	for i, v := range p1.iceCreamFlavors {
		fmt.Println(i, v)
	}

	fmt.Println(p2.first, p2.last)
	for i, v := range p2.iceCreamFlavors {
		fmt.Println(i, v)
	}


}