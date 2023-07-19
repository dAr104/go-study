package main

import "fmt"

/*
Interfaces allow us to define behavior and often also allow us to do polymorphism
Interface is a type: a value can be of more than one type
*/

type person struct{
	first string
	last string
}

type secretAgent struct{
	person
	ltk bool
}

// func (r receiver) identifier(parametes (return(s)) {code}

// when specify a receiver, you attach the function to any value of that type (it is a method)
func (s secretAgent) speak() {
	fmt.Println("I'm", s.first, s.last, "- the secretAgent speak")
}

func (p person) speak() {
	fmt.Println("I'm", p.first, p.last, " - the person speak")
}

// any type has this method speack is also a type human
type human interface {
	speak()
}

// if you have an empty interface, every type have that value; every value implements empty interface

// asserting: verify the typ and then turn to the concrete type
func foo(h human) {
	switch h.(type){
	case person: 
		fmt.Println("I'm a human and a person", h.(person).first)
	case secretAgent:
		fmt.Println("I'm a human and a secretAgent", h.(secretAgent).first)
	}
 }

func main(){
	sa1 := secretAgent{
		person: person{
			"James",
			"Bond",
		},
		ltk: true,
	}

	sa2 := secretAgent{
		person: person{
			"Dario",
			"Mario",
		},
		ltk: true,
	}

	p1 := person{
		first: "Dr.",
		last: "Yes",
	}

	fmt.Println(sa1)
	sa1.speak()
	sa2.speak()

	fmt.Println(p1)

	foo(sa1)
	foo(sa2)
	foo(p1)
	
}

