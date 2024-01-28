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
			tkn = l.makeTwoCharTokenEQ(token.Eq)
		} else {
			tkn = token.FromChar(token.Assign, l.ch)
		}
	case ';':
		tkn = token.FromChar(token.Semicolon, l.ch)
	case '(':
		tkn = token.FromChar(token.Lparen, l.ch)
	case ')':
		tkn = token.FromChar(token.Rparen, l.ch)
	case ',':
		tkn = token.FromChar(token.Comma, l.ch)
	case '+':
		if l.peekChar() == '=' {
			tkn = l.makeTwoCharTokenEQ(token.PlusAssign)
		} else {
			tkn = token.FromChar(token.Plus, l.ch)
		}
	case '-':
		if l.peekChar() == '=' {
			tkn = l.makeTwoCharTokenEQ(token.MinusAssign)
		} else {
			tkn = token.FromChar(token.Minus, l.ch)
		}
	case '!':
		if l.peekChar() == '=' {
			tkn = l.makeTwoCharTokenEQ(token.Noteq)
		} else {
			tkn = token.FromChar(token.Bang, l.ch)
		}
	case '*':
		if l.peekChar() == '=' {
			tkn = l.makeTwoCharTokenEQ(token.AsteriskAssign)
		} else {
			tkn = token.FromChar(token.Asterisk, l.ch)
		}
	case '/':
		if l.peekChar() == '=' {
			tkn = l.makeTwoCharTokenEQ(token.SlashAssign)
		} else {
			tkn = token.FromChar(token.Slash, l.ch)
		}
	case '%':
		if l.peekChar() == '=' {
			tkn = l.makeTwoCharTokenEQ(token.ModuloAssign)
		} else {
			tkn = token.FromChar(token.Modulo, l.ch)
		}
	case '<':
		if l.peekChar() == '=' {
			tkn = l.makeTwoCharTokenEQ(token.Lteq)
		} else {
			tkn = token.FromChar(token.Lt, l.ch)
		}
	case '>':
		if l.peekChar() == '=' {
			tkn = l.makeTwoCharTokenEQ(token.Gteq)
		} else {
			tkn = token.FromChar(token.Gt, l.ch)
		}
	case '{':
		tkn = token.FromChar(token.Lbrace, l.ch)
	case '}':
		tkn = token.FromChar(token.Rbrace, l.ch)
	case 0:
		tkn = token.FromString(token.EOF, "")
	default:
		if isLetter(l.ch) {
			tkn.Literal = l.readAll(isLetter) // set the literal first to use it in the table
			tkn.Type = token.IdentLookup(tkn.Literal)
			return
		} else if isDigit(l.ch) {
			tkn.Type = token.Int
			tkn.Literal = l.readAll(isDigit)
			return
		}

		tkn = token.FromChar(token.Illegal, l.ch)
	}

	l.readChar()
	return
}

// makeTwoCharTokenEQ builds a two-char input where the later char is a '='.
func (l *Lexer) makeTwoCharTokenEQ(tokenType token.Type) token.Token {
	ch := l.ch                                                  // save the current char.
	l.readChar()                                                // increment to the next char. l.ch is now '='.
	return token.FromString(tokenType, string(ch)+string(l.ch)) // build the token.
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
