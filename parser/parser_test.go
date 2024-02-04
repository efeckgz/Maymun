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
	let pi = 3.14;
	`

	l := lexer.New(input)
	p := New(l)

	program := p.ParseProgram()
	if program == nil {
		t.Fatal("ParseProgram() returned nil.")
	}

	checkParseErrors(t, p)

	if len(program.Statements) != 4 {
		t.Fatalf("Unexpected statement count: expected 3, got %d.\n", len(program.Statements))
	}

	tests := []struct {
		expectedIdentifier string
	}{
		{"x"},
		{"y"},
		{"foobar"},
		{"pi"},
	}

	for i, tt := range tests {
		stmt := program.Statements[i]
		if !letStatementOk(t, stmt, tt.expectedIdentifier) {
			return
		}
	}
}

func TestReturnStatements(t *testing.T) {
	input := `
	return 5;
	return 10;
	return 993322;
	return 3.14;
	`

	l := lexer.New(input)
	p := New(l)

	program := p.ParseProgram()
	checkParseErrors(t, p)

	if len(program.Statements) != 4 {
		t.Fatalf("Unexpected number of statements: expected 4, got %d.\n", len(program.Statements))
	}

	for _, stmt := range program.Statements {
		rs, ok := stmt.(*ast.ReturnStatement)
		if !ok {
			t.Errorf("Unexpected statement type: expected *ast.ReturnStatement, got %T.\n", stmt)
			continue
		}

		if rs.TokenLiteral() != "return" {
			t.Errorf("Unexpected token literal: expected return, got %s.\n", rs.TokenLiteral())
		}
	}
}

func checkParseErrors(t *testing.T, p *Parser) {
	errors := p.Errors()
	if len(errors) == 0 {
		return
	}

	t.Errorf("The parsers has %d errors.\n", len(errors))
	for _, err := range errors {
		t.Errorf("%q\n", err)
	}

	t.FailNow()
}

func letStatementOk(t *testing.T, s ast.Statement, name string) bool {
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
