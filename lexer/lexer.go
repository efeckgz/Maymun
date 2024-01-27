package lexer

import "github.com/efeckgz/Maymun/token"

// Lexer represents the lexer type.
type Lexer struct {
	input   string
	pos     int  // current position in input (points to current char)
	nextPos int  // current reading postion in input (after current char)
	ch      byte // current char under examination
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
		if l.peekChar() == '=' {
			ch := l.ch                                                // save the currently matched '='
			l.readChar()                                              // read one more character, l.ch is now the next '='
			tkn = token.FromString(token.EQ, string(ch)+string(l.ch)) // build the token.
		} else {
			tkn = token.FromChar(token.ASSIGN, l.ch)
		}
	case ';':
		tkn = token.FromChar(token.SEMICOLON, l.ch)
	case '(':
		tkn = token.FromChar(token.LPAREN, l.ch)
	case ')':
		tkn = token.FromChar(token.RPAREN, l.ch)
	case ',':
		tkn = token.FromChar(token.COMMA, l.ch)
	case '+':
		tkn = token.FromChar(token.PLUS, l.ch)
	case '-':
		tkn = token.FromChar(token.MINUS, l.ch)
	case '!':
		if l.peekChar() == '=' {
			ch := l.ch
			l.readChar()
			tkn = token.FromString(token.NOTEQ, string(ch)+string(l.ch))
		} else {
			tkn = token.FromChar(token.BANG, l.ch)
		}
	case '*':
		tkn = token.FromChar(token.ASTERISK, l.ch)
	case '/':
		tkn = token.FromChar(token.SLASH, l.ch)
	case '<':
		if l.peekChar() == '=' {
			ch := l.ch
			l.readChar()
			tkn = token.FromString(token.LTEQ, string(ch)+string(l.ch))
		} else {
			tkn = token.FromChar(token.LT, l.ch)
		}
	case '>':
		if l.peekChar() == '=' {
			ch := l.ch
			l.readChar()
			tkn = token.FromString(token.GTEQ, string(ch)+string(l.ch))
		} else {
			tkn = token.FromChar(token.GT, l.ch)
		}
	case '{':
		tkn = token.FromChar(token.LBRACE, l.ch)
	case '}':
		tkn = token.FromChar(token.RBRACE, l.ch)
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

		tkn = token.FromChar(token.ILLEGAL, l.ch)
	}

	l.readChar()
	return
}

// readAll reads all the chars that satisfy the given condition.
func (l *Lexer) readAll(condition func(ch byte) bool) string {
	pos := l.pos
	for condition(l.ch) {
		l.readChar()
	}

	return l.input[pos:l.pos]
}

func (l *Lexer) readChar() {
	if l.nextPos >= len(l.input) {
		l.ch = 0
	} else {
		l.ch = l.input[l.nextPos]
	}

	l.pos = l.nextPos
	l.nextPos++
}

func (l *Lexer) skipWhitespace() {
	for l.ch == ' ' || l.ch == '\t' || l.ch == '\r' || l.ch == '\n' {
		l.readChar()
	}
}

func (l *Lexer) peekChar() byte {
	if l.nextPos > len(l.input) {
		return 0
	} else {
		return l.input[l.nextPos] // do not increment the next pos as we are not moving to it.
	}
}

func isLetter(ch byte) bool {
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_'
}

func isDigit(ch byte) bool {
	return '0' <= ch && ch <= '9'
}
