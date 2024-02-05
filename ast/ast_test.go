package ast

import (
	"testing"

	"github.com/efeckgz/Maymun/token"
)

func TestString(t *testing.T) {
	input := "let myVar = anotherVar;"

	program := &Program{
		Statements: []Statement{
			&LetStatement{
				Token: token.FromString(token.Let, "let"),
				Name: &Identifier{
					Token: token.FromString(token.Ident, "myVar"),
					Value: "myVar",
				},
				Value: &Identifier{
					Token: token.FromString(token.Ident, "anotherVar"),
					Value: "anotherVar",
				},
			},
		},
	}

	if program.String() != input {
		t.Errorf("Unexpected program: expected '%s', got '%s'.\n", input, program.String())
	}
}
