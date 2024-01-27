package token

// constants for the tokens.
const (
	Illegal = "Illegal"
	EOF     = "EOF"

	// Identifiers + literals
	Ident = "Ident"
	Int   = "Int"

	// Operators
	Assign   = "="
	Plus     = "+"
	Minus    = "-"
	Bang     = "!"
	Asterisk = "*"
	Slash    = "/"

	Lt    = "<"
	Gt    = ">"
	Eq    = "=="
	Noteq = "!="
	Lteq  = "<="
	Gteq  = ">="

	// Delimiters
	Comma     = ","
	Semicolon = ";"

	Lparen = "("
	Rparen = ")"
	Lbrace = "{"
	Rbrace = "}"

	// Keywords
	Function = "Function"
	Let      = "Let"
	True     = "True"
	False    = "False"
	If       = "If"
	Else     = "Else"
	Return   = "Retrun"
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
