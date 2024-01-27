package token

// constants for the tokens.
const (
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"

	// Identifiers + literals
	IDENT = "IDENT"
	INT   = "INT"

	// Operators
	ASSIGN   = "="
	PLUS     = "+"
	MINUS    = "-"
	BANG     = "!"
	ASTERISK = "*"
	SLASH    = "/"

	LT    = "<"
	GT    = ">"
	EQ    = "=="
	NOTEQ = "!="
	LTEQ  = "<="
	GTEQ  = ">="

	// Delimiters
	COMMA     = ","
	SEMICOLON = ";"

	LPAREN = "("
	RPAREN = ")"
	LBRACE = "{"
	RBRACE = "}"

	// Keywords
	FUNCTION = "FUNCTION"
	LET      = "LET"
	TRUE     = "TRUE"
	FALSE    = "FALSE"
	IF       = "IF"
	ELSE     = "ELSE"
	RETURN   = "RETURN"
)

var keywords = map[string]Type{
	"fn":     FUNCTION,
	"let":    LET,
	"true":   TRUE,
	"false":  FALSE,
	"if":     IF,
	"else":   ELSE,
	"return": RETURN,
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

	return IDENT
}
