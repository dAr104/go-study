package main

import (
	"fmt"
	"io"
	"os"
)

func main() {

	fmt.Println("Ciao")
	fmt.Fprintln(os.Stdout, "Ciao")
	io.WriteString(os.Stdout, "Ciao")
}