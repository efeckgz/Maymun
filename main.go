package main

import (
	"fmt"
	"os"

	"github.com/efeckgz/Maymun/repl"
)

func main() {
	fmt.Println("Maymun programming language v0.0.0")
	fmt.Println("Please start by typing commands.")

	repl.Start(os.Stdin, os.Stdout)
}
