package main

import "fmt"

// a struct is an composite data type (composite data types, aka, aggregate data types, aka, complex data types)
// Struct allow us to compose toghether values of different types
type person struct {
	first string
	last string
	age int
}

// you can embedded structs
type secretAgent struct {
	person  // anonymous field (or embedded field): you have not to give a identifier. The unqualifield type name acts as the field name
	ltk bool
}

func main(){
	p1 := person{
		first: "James", // field
		last: "Bond",
		age: 34,
	}

	p2 := person{
		first: "Miss",
		last: "Moneypenny",
		age: 22,
	}

	sa1 := secretAgent{
		person: person{
			first: "James",
			last: "Bond",
			age: 34,
		},
		ltk: true,
	}

	fmt.Println(p1)
	fmt.Println(p2.first, p2.last, p2.age)
	fmt.Println(sa1)
	fmt.Println(sa1.person)
	// in this case first is in person: the inner type gets promoted to the outer type
	fmt.Println(sa1.first, sa1.last, sa1.age, sa1.ltk)

	// it is similar to an object or class, is a good example but in go we create a "value of type person"
}