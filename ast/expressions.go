package ast

import "github.com/efeckgz/Maymun/token"

// Identifier represents the left side of the assignment statement. It is the name of the variable.
// For simplicity, Identifier implements the Expression interface. This means the parser treats the variable
// names as expressions in Maymun.
type Identifier struct {
	Token token.Token // token.Ident
	Value string      // value should be the same as Token.Literal
}

func (i *Identifier) expressionNode() {} // dummy method

// TokenLiteral returns the literal of the token it is associated with.
func (i *Identifier) TokenLiteral() string {
	return i.Token.Literal
}
