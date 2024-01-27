package lexer

import (
	"testing"

	"github.com/efeckgz/Maymun/token"
)

type testToken struct {
	expectedType    token.Type
	expectedLiteral string
}

func TestNextToken(t *testing.T) {
	input := `let five = 5;
let ten = 10;

let add = fn(x, y) {
	x + y;
};

let result = add(five, ten);
!-/*5;
5 < 10 > 5;

if (5 < 10) {
	return true;
} else {
	return false;
}

10 == 10;
10 != 9;
10 >= 9;
10 <= 10;
`

	tests := []testToken{
		{token.Let, "let"},
		{token.Ident, "five"},
		{token.Assign, "="},
		{token.Int, "5"},
		{token.Semicolon, ";"},
		{token.Let, "let"},
		{token.Ident, "ten"},
		{token.Assign, "="},
		{token.Int, "10"},
		{token.Semicolon, ";"},
		{token.Let, "let"},
		{token.Ident, "add"},
		{token.Assign, "="},
		{token.Function, "fn"},
		{token.Lparen, "("},
		{token.Ident, "x"},
		{token.Comma, ","},
		{token.Ident, "y"},
		{token.Rparen, ")"},
		{token.Lbrace, "{"},
		{token.Ident, "x"},
		{token.Plus, "+"},
		{token.Ident, "y"},
		{token.Semicolon, ";"},
		{token.Rbrace, "}"},
		{token.Semicolon, ";"},
		{token.Let, "let"},
		{token.Ident, "result"},
		{token.Assign, "="},
		{token.Ident, "add"},
		{token.Lparen, "("},
		{token.Ident, "five"},
		{token.Comma, ","},
		{token.Ident, "ten"},
		{token.Rparen, ")"},
		{token.Semicolon, ";"},
		{token.Bang, "!"},
		{token.Minus, "-"},
		{token.Slash, "/"},
		{token.Asterisk, "*"},
		{token.Int, "5"},
		{token.Semicolon, ";"},
		{token.Int, "5"},
		{token.Lt, "<"},
		{token.Int, "10"},
		{token.Gt, ">"},
		{token.Int, "5"},
		{token.Semicolon, ";"},
		{token.If, "if"},
		{token.Lparen, "("},
		{token.Int, "5"},
		{token.Lt, "<"},
		{token.Int, "10"},
		{token.Rparen, ")"},
		{token.Lbrace, "{"},
		{token.Return, "return"},
		{token.True, "true"},
		{token.Semicolon, ";"},
		{token.Rbrace, "}"},
		{token.Else, "else"},
		{token.Lbrace, "{"},
		{token.Return, "return"},
		{token.False, "false"},
		{token.Semicolon, ";"},
		{token.Rbrace, "}"},
		{token.Int, "10"},
		{token.Eq, "=="},
		{token.Int, "10"},
		{token.Semicolon, ";"},
		{token.Int, "10"},
		{token.Noteq, "!="},
		{token.Int, "9"},
		{token.Semicolon, ";"},
		{token.Int, "10"},
		{token.Gteq, ">="},
		{token.Int, "9"},
		{token.Semicolon, ";"},
		{token.Int, "10"},
		{token.Lteq, "<="},
		{token.Int, "10"},
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
