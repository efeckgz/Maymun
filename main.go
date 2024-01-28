package main

import (
	"fmt"
	"os"

	"github.com/efeckgz/Maymun/repl"
)

func main() {
	fmt.Println("Welcome to Maymun programming language.")
	fmt.Println("Please start by typing commands.")

	repl.Start(os.Stdin, os.Stdout)
}
