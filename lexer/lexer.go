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
	l.skipWhitespace()

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
	case '-':
		tkn = token.New(token.MINUS, l.ch)
	case '!':
		tkn = token.New(token.BANG, l.ch)
	case '*':
		tkn = token.New(token.ASTERISK, l.ch)
	case '/':
		tkn = token.New(token.SLASH, l.ch)
	case '<':
		tkn = token.New(token.LT, l.ch)
	case '>':
		tkn = token.New(token.GT, l.ch)
	case '{':
		tkn = token.New(token.LBRACE, l.ch)
	case '}':
		tkn = token.New(token.RBRACE, l.ch)
	case 0:
		tkn.Literal = ""
		tkn.Type = token.EOF
	default:
		if isLetter(l.ch) {
			tkn.Literal = l.readAll(isLetter) // set the literal first to use it in the table
			tkn.Type = token.IdentLookup(tkn.Literal)
			return
		} else if isDigit(l.ch) {
			tkn.Type = token.INT
			tkn.Literal = l.readAll(isDigit)
			return
		}

		tkn = token.New(token.ILLEGAL, l.ch)
	}

	l.readChar()
	return
}

// readAll reads all the chars that satisfy the given condition.
func (l *Lexer) readAll(condition func(ch byte) bool) string {
	pos := l.position
	for condition(l.ch) {
		l.readChar()
	}

	return l.input[pos:l.position]
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

func (l *Lexer) skipWhitespace() {
	for l.ch == ' ' || l.ch == '\t' || l.ch == '\r' || l.ch == '\n' {
		l.readChar()
	}
}

func isLetter(ch byte) bool {
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_'
}

func isDigit(ch byte) bool {
	return '0' <= ch && ch <= '9'
}
