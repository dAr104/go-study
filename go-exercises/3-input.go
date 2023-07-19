package main

import (
	"bufio"
	"fmt"
	"os"
)

/*
	Package bufio implements buffered I/O. It wraps an io.Reader or io.Writer object,
	creating another object (Reader or Writer) that also implements the interface but provides buffering and some help for textual I/O.
*/

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter your name: ")
	name, _ := reader.ReadString('\n')
	fmt.Println("Your name is " + name)
}