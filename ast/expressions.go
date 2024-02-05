package ast

import (
	"bytes"

	"github.com/efeckgz/Maymun/token"
)

// Identifier represents the left side of the assignment statement. It is the name of the variable.
// For simplicity, Identifier implements the Expression interface. This means the parser treats the variable
// names as expressions in Maymun.
type Identifier struct {
	Token token.Token // token.Ident
	Value string      // value should be the same as Token.Literal
}

// TokenLiteral returns the literal of the token it is associated with.
func (i *Identifier) TokenLiteral() string {
	return i.Token.Literal
}

func (i *Identifier) String() string {
	return i.Value
}

// IntegerLiteral represents the integer expressions.
type IntegerLiteral struct {
	Token token.Token
	Value int64
}

// FloatLiteral represents the float expression.
type FloatLiteral struct {
	Token token.Token
	Value float64
}

// TokenLiteral returns the literal of the token it is associated with.
func (fl *FloatLiteral) TokenLiteral() string {
	return fl.Token.Literal
}

func (fl *FloatLiteral) String() string {
	return fl.Token.Literal
}

// TokenLiteral returns the literal of the token it is assciated with.
func (il *IntegerLiteral) TokenLiteral() string {
	return il.Token.Literal
}

func (il *IntegerLiteral) String() string {
	return il.Token.Literal
}

// PrefixExpression represents expressions in the form of !ok or -4.
type PrefixExpression struct {
	Token    token.Token
	Operator string
	Right    Expression
}

// TokenLiteral returns the literal of the token it is associated with.
func (pe *PrefixExpression) TokenLiteral() string {
	return pe.Token.Literal
}

func (pe *PrefixExpression) String() string {
	var out bytes.Buffer

	out.WriteString("(")
	out.WriteString(pe.Operator)
	out.WriteString(pe.Right.String())
	out.WriteString(")")

	return out.String()
}

// Implement the Expression interface
func (i *Identifier) expressionNode()        {}
func (il *IntegerLiteral) expressionNode()   {}
func (fl *FloatLiteral) expressionNode()     {}
func (pe *PrefixExpression) expressionNode() {}
