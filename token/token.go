package token

const (
	ILLEGAL = "ILLEGAL"
	EOF 	= "EOF"

	// Identifiers + literals
	IDENT 	= "IDENT"
	INT	= "INT"

	// Operators
	ASSIGN 	= "="
	PLUS	= "+"

	// Delimiters
	COMMA	  = ","
	SEMICOLON = ";"

	LPAREN = "("
	RPAREN = ")"
	LBRACE = "{"
	RBRACE = "}"

	// Keywords
	FUNCTION = "FUNCTION"
	LET	 = "LET"
)

type Token string

type TokenType struct {
	Type 	TokenType
	Literal string
}

