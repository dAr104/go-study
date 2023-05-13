package main

import (
	"fmt"
	"math"
)

/*
● create a type SQUARE
● create a type CIRCLE
● attach a method to each that calculates AREA and returns it
	○ circle area= π r 2
	○ square area = L * W
● create a type SHAPE that defines an interface as anything that has the AREA method
● create a func INFO which takes type shape and then prints the area
● create a value of type square
● create a value of type circle
● use func info to print the area of square
● use func info to print the area of circle*/

func (s square) area() float32{
	return s.length * s.length
}

func (c circle) area() float32{
	return (c.radius * c.radius) * math.Pi
}

type square struct {
	length float32
}

type circle struct {
	radius float32
}

type shape interface {
	area() float32
}

func info(s shape) {
	fmt.Println(s.area())
}


func main() {
	c1 := circle{
		radius: 4.5,
	}

	s1 := square{
		length: 3,
	
	}

	info(c1)
	info(s1)
}