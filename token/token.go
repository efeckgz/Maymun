package token

// constants for the tokens.
const (
	Illegal Type = "Illegal"
	EOF     Type = "EOF"

	// Identifiers + literals
	Ident Type = "Ident"
	Int   Type = "Int"
	Float Type = "Float"

	// Operators
	Assign         Type = "="
	Plus           Type = "+"
	Minus          Type = "-"
	Bang           Type = "!"
	Asterisk       Type = "*"
	Slash          Type = "/"
	Modulo         Type = "%"
	PlusAssign     Type = "+="
	MinusAssign    Type = "-="
	AsteriskAssign Type = "*="
	SlashAssign    Type = "/="
	ModuloAssign   Type = "%="

	Lt    Type = "<"
	Gt    Type = ">"
	Eq    Type = "=="
	Noteq Type = "!="
	Lteq  Type = "<="
	Gteq  Type = ">="

	// Delimiters
	Comma     Type = ","
	Semicolon Type = ";"

	Lparen Type = "("
	Rparen Type = ")"
	Lbrace Type = "{"
	Rbrace Type = "}"

	// Keywords
	Function Type = "Function"
	Let      Type = "Let"
	True     Type = "True"
	False    Type = "False"
	If       Type = "If"
	Else     Type = "Else"
	Return   Type = "Retrun"
)

var keywords = map[string]Type{
	"fn":     Function,
	"let":    Let,
	"true":   True,
	"false":  False,
	"if":     If,
	"else":   Else,
	"return": Return,
}

// Type represents the type of the token. Setting it to string allows to use many things as types.
type Type string

// A Token has a Type and a Literal.
type Token struct {
	Type    Type
	Literal string
}

// FromChar creates a new token from one char.
func FromChar(tokenType Type, ch byte) Token {
	return Token{Type: tokenType, Literal: string(ch)}
}

// FromString returns a token from a string literal.
func FromString(tokenType Type, literal string) Token {
	return Token{Type: tokenType, Literal: literal}
}

// IdentLookup checks the keywords table to see weather the given identifier is in fact a keyword.
// If it is, the TokenType of that keyword is returned. If not, than the identifier is user defined, so
// we return token.IDENT.
func IdentLookup(ident string) Type {
	token, reserved := keywords[ident]
	if reserved {
		return token
	}

	return Ident
}
