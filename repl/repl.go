package repl

import (
	"bufio"
	"fmt"
	"io"

	"github.com/efeckgz/Maymun/lexer"
	"github.com/efeckgz/Maymun/token"
)

// Prompt is the input prompt to be printed to the console.
const Prompt = ">> "

// Start initiates the read-eval-print-loop.
func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)

	for {
		fmt.Print(Prompt)

		scanned := scanner.Scan()
		if !scanned {
			return
		}

		line := scanner.Text()
		l := lexer.New(line)
		for tkn := l.NextToken(); tkn.Type != token.EOF; tkn = l.NextToken() {
			fmt.Printf("%+v\n", tkn)
		}
	}
}
