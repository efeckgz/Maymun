package lexer

import (
	"testing"

	"github.com/efeckgz/Maymun/token"
)

type testToken struct {
	expectedType    token.TokenType
	expectedLiteral string
}

func TestNextToken(t *testing.T) {
	input := `=+(){},;`

	tests := []testToken{
		{token.ASSIGN, "="},
		{token.PLUS, "+"},
		{token.LPAREN, "("},
		{token.RPAREN, ")"},
		{token.LBRACE, "{"},
		{token.RBRACE, "}"},
		{token.COMMA, ","},
		{token.SEMICOLON, ";"},
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
