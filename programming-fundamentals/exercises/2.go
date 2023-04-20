package main

import "fmt"

/*
Using the following operators, write expressions and assign their values to variables:
g. ==
h. <=
i. >=
j. !=
k. <
l. >
Now print each of the variables
*/

func main() {
	x := 33
	y := 42

	g := x == y
	h := x <= y
	i := x >= y	
	j := x != y
	k := x < y
	l := x > y

	fmt.Println(g, h, i, j, k, l)
}