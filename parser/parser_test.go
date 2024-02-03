package parser

import (
	"testing"

	"github.com/efeckgz/Maymun/ast"
	"github.com/efeckgz/Maymun/lexer"
)

func TestLetStatements(t *testing.T) {
	input := `
	let x = 5;
	let y = 10;
	let foobar = 838383;
	`

	l := lexer.New(input)
	p := New(l)

	program := p.ParseProgram()
	if program == nil {
		t.Fatal("ParseProgram() returned nil.")
	}

	if len(program.Statements) != 3 {
		t.Fatalf("Unexpected statement count: expected 3, got %d.\n", len(program.Statements))
	}

	tests := []struct {
		expectedIdentifier string
	}{
		{"x"},
		{"y"},
		{"foobar"},
	}

	for i, tt := range tests {
		stmt := program.Statements[i]

		if stmt == nil {
			t.Fatalf("The statement %d is nil.", i)
		}

		if !testLetStatement(t, stmt, tt.expectedIdentifier) {
			return
		}
	}
}

func testLetStatement(t *testing.T, s ast.Statement, name string) bool {
	if s.TokenLiteral() != "let" {
		t.Errorf("Unexpected statement token literal: expected let, got %s.\n", s.TokenLiteral())
		return false
	}

	letStmt, ok := s.(*ast.LetStatement)
	if !ok {
		t.Errorf("Unexpected statement type: expected 'let', got '%T'.\n", s)
		return false
	}

	if letStmt.Name.Value != name {
		t.Errorf("Unexpected name for the let statement: expected '%s', got '%s',\n", name, letStmt.Name.Value)
		return false
	}

	if letStmt.Name.TokenLiteral() != name {
		t.Errorf("Unexpected s.Name: expected '%s', got '%s'.\n", name, letStmt.Name)
		return false
	}

	return true
}
