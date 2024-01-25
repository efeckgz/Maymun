package lexer

import "github.com/efeckgz/Maymun/token"

// Lexer represents the lexer type.
type Lexer struct {
	input        string
	position     int  // current position in input (points to current char)
	readPosition int  // current reading postion in input (after current char)
	ch           byte // current char under examination
}

// New returns a new lexer from a given input.
func New(input string) *Lexer {
	l := &Lexer{input: input}
	l.readChar()
	return l
}

// NextToken goes through the input and output the next token it recognizes.
func (l *Lexer) NextToken() (tkn token.Token) {
	switch l.ch {
	case '=':
		tkn = token.New(token.ASSIGN, l.ch)
	case ';':
		tkn = token.New(token.SEMICOLON, l.ch)
	case '(':
		tkn = token.New(token.LPAREN, l.ch)
	case ')':
		tkn = token.New(token.RPAREN, l.ch)
	case ',':
		tkn = token.New(token.COMMA, l.ch)
	case '+':
		tkn = token.New(token.PLUS, l.ch)
	case '{':
		tkn = token.New(token.LBRACE, l.ch)
	case '}':
		tkn = token.New(token.RBRACE, l.ch)
	case 0:
		tkn.Literal = ""
		tkn.Type = token.EOF
	}

	l.readChar()
	return
}

func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) {
		l.ch = 0
	} else {
		l.ch = l.input[l.readPosition]
	}

	l.position = l.readPosition
	l.readPosition++
}
