package parser

import (
	"fmt"
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

func TestIdentifierExpression(t *testing.T) {
	input := "foobar;"

	l := lexer.New(input)
	p := New(l)
	program := p.ParseProgram()
	checkParseErrors(t, p)

	if len(program.Statements) != 1 {
		t.Fatalf("Unexpected number of statemens: expected 1, got %d.\n", len(program.Statements))
	}

	// Check if the statement is an expression.
	stmt, ok := program.Statements[0].(*ast.ExpressionStatement)
	if !ok {
		t.Fatalf("Unexpected statement type: expected *ast.ExpressionStatement, got %T\n.", stmt)
	}

	// Check if the expression is an identifier.
	ident, ok := stmt.Expression.(*ast.Identifier)
	if !ok {
		t.Fatalf("Unexpected expression type: expected *ast.Identifier, got %T.\n", ident)
	}

	// Check if the values of the identifier are correct.
	if ident.Value != "foobar" {
		t.Fatalf("Unexpected identifier value: expected 'foobar', got '%s'.\n", ident.Value)
	}

	if ident.TokenLiteral() != "foobar" {
		t.Fatalf("Unexpected identifier token literal: expected 'foobar', got '%s'.\n", ident.TokenLiteral())
	}
}

func TestIntegerExpression(t *testing.T) {
	input := "5;"

	l := lexer.New(input)
	p := New(l)

	program := p.ParseProgram()
	checkParseErrors(t, p)

	s, ok := program.Statements[0].(*ast.ExpressionStatement)
	if !ok {
		t.Fatalf("Unexpected statement type, expected *ast.ExpressionStatement, got %T.\n", s)
	}

	literal, ok := s.Expression.(*ast.IntegerLiteral)
	if !ok {
		t.Fatalf("Unexpected type for literal: expected *ast.IntegerLiteral, got %T.\n", literal)
	}

	if literal.Value != 5 {
		t.Errorf("Unexpected literal value: expected 5, got %d.\n", literal.Value)
	}

	if literal.TokenLiteral() != "5" {
		t.Errorf("Unexpected literal token literal: expected '5', got '%s'.\n", literal.TokenLiteral())
	}
}

func TestFloatExpression(t *testing.T) {
	input := "3.14;"

	l := lexer.New(input)
	p := New(l)

	program := p.ParseProgram()
	checkParseErrors(t, p)

	s, ok := program.Statements[0].(*ast.ExpressionStatement)
	if !ok {
		t.Fatalf("Unexpected statement type: expected *ast.ExpressionStatement, got %T.\n", s)
	}

	literal, ok := s.Expression.(*ast.FloatLiteral)
	if !ok {
		t.Fatalf("Unexpected expression type: expected *ast.FloatLiteral, got %T.\n", literal)
	}

	if literal.TokenLiteral() != "3.14" {
		t.Errorf("Unexpected token literal: expected '3.14', got '%s'.\n", literal.TokenLiteral())
	}

	if literal.Value != 3.14 {
		t.Errorf("Unexpected token value: expected 3.14, got %.2f.\n", literal.Value)
	}
}

func TestParsingPrefixExpressions(t *testing.T) {
	prefixTests := []struct {
		input        string
		operator     string
		integerValue int64
	}{
		{"!5;", "!", 5},
		{"-15", "-", 15},
	}

	for _, tt := range prefixTests {
		l := lexer.New(tt.input)
		p := New(l)
		program := p.ParseProgram()
		checkParseErrors(t, p)

		if len(program.Statements) != 1 {
			t.Fatalf("program.Statements does not contain %d statements. got=%d\n",
				1, len(program.Statements))
		}

		s, ok := program.Statements[0].(*ast.ExpressionStatement)
		if !ok {
			t.Fatalf("Unexpected statement type, expected *ast.ExpressionStatement, got %T.\n", s)
		}

		exp, ok := s.Expression.(*ast.PrefixExpression)
		if !ok {
			t.Fatalf("Unexpected expression type: expected *ast.PrefixExpression, got %T.\n", exp)
		}

		if exp.Operator != tt.operator {
			t.Fatalf("Unexpected operator: expected %s, got %s\n", tt.operator, exp.Operator)
		}

		if !integertLiteralOk(t, exp.Right, tt.integerValue) {
			return
		}
	}
}

func TestParsingInfixExpressions(t *testing.T) {
	infixTests := []struct {
		input      string
		leftValue  int64
		operator   string
		rightValue int64
	}{
		{"5 + 5;", 5, "+", 5},
		{"5 - 5;", 5, "-", 5},
		{"5 * 5;", 5, "*", 5},
		{"5 / 5;", 5, "/", 5},
		{"5 > 5;", 5, ">", 5},
		{"5 < 5;", 5, "<", 5},
		{"5 == 5;", 5, "==", 5},
		{"5 != 5;", 5, "!=", 5},
	}

	for _, tt := range infixTests {
		l := lexer.New(tt.input)
		p := New(l)
		program := p.ParseProgram()
		checkParseErrors(t, p)

		if len(program.Statements) != 1 {
			t.Fatalf("Unexpected number of statements: expected 1, got %d.\n", len(program.Statements))
		}

		stmt, ok := program.Statements[0].(*ast.ExpressionStatement)
		if !ok {
			t.Fatalf("Unexpected type of statement: expected *ast.ExpressionStatement, got %T.\n", program.Statements[0])
		}

		exp, ok := stmt.Expression.(*ast.InfixExpression)
		if !ok {
			t.Fatalf("Unexpected type of expression: expected *ast.InfixExpression, got %T.\n", stmt.Expression)
		}

		if !integertLiteralOk(t, exp.Left, tt.leftValue) {
			return
		}

		if exp.Operator != tt.operator {
			t.Fatalf("Unexpected operator: expected %s, got %s.\n", tt.operator, exp.Operator)
		}

		if !integertLiteralOk(t, exp.Right, tt.rightValue) {
			return
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

func integertLiteralOk(t *testing.T, il ast.Expression, value int64) bool {
	integ, ok := il.(*ast.IntegerLiteral)
	if !ok {
		t.Errorf("Unexpected type: expected *ast.IntegerLiteral, got %T.\n", integ)
		return false
	}

	if integ.Value != value {
		t.Errorf("Unexpected value: expected %d, got %d.\n", value, integ.Value)
		return false
	}

	if integ.TokenLiteral() != fmt.Sprintf("%d", value) {
		t.Errorf("Unexpected token literal: expected %d, got %s.\n", value, integ.TokenLiteral())
		return false
	}

	return true
}
