package lexer

import (
	"testing"

	"github.com/efeckgz/Maymun/token"
)

func TestNextToken(t *testing.T) {
	input := `
		let a = 5;
		let b = 10;
		a += b;
		a -= b;
		a *= b;
		a /= b;
		a %= b;
		a % b;
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
		{token.Ident, "a"},
		{token.PlusAssign, "+="},
		{token.Ident, "b"},
		{token.Semicolon, ";"},
		{token.Ident, "a"},
		{token.MinusAssign, "-="},
		{token.Ident, "b"},
		{token.Semicolon, ";"},
		{token.Ident, "a"},
		{token.AsteriskAssign, "*="},
		{token.Ident, "b"},
		{token.Semicolon, ";"},
		{token.Ident, "a"},
		{token.SlashAssign, "/="},
		{token.Ident, "b"},
		{token.Semicolon, ";"},
		{token.Ident, "a"},
		{token.ModuloAssign, "%="},
		{token.Ident, "b"},
		{token.Semicolon, ";"},
		{token.Ident, "a"},
		{token.Modulo, "%"},
		{token.Ident, "b"},
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
