package lexer

import (
	"testing"

	"github.com/efeckgz/Maymun/token"
)

func TestNextToken(t *testing.T) {
	input := `
	let a = 5;
	let b = 10;
	let pi = 3.14;
	`

	tests := []struct {
		expectedType    token.Type
		expectedLiteral string
	}{
		{token.Let, "let"},
		{token.Ident, "a"},
		{token.Assign, "="},
		{token.Int, "5"},
		{token.Semicolon, ";"},
		{token.Let, "let"},
		{token.Ident, "b"},
		{token.Assign, "="},
		{token.Int, "10"},
		{token.Semicolon, ";"},
		{token.Let, "let"},
		{token.Ident, "pi"},
		{token.Assign, "="},
		{token.Float, "3.14"},
		{token.Semicolon, ";"},
		{token.EOF, ""},
	}

	l := New(input)

	for i, tt := range tests {
		tkn := l.NextToken()

		if tkn.Type != tt.expectedType {
			t.Fatalf("test[%d] - tokentype wrong: expected %q, got %q.\n", i, tt.expectedType, tkn.Type)
		}

		if tkn.Literal != tt.expectedLiteral {
			t.Fatalf("test[%d] - literal wrong: expected %q, got %q.\n", i, tt.expectedLiteral, tkn.Literal)
		}
	}
}
