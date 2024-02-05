package ast

import (
	"bytes"

	"github.com/efeckgz/Maymun/token"
)

// LetStatement represents the let statement node in the AST.
type LetStatement struct {
	Token token.Token // token.Let
	Name  *Identifier // name of the variable.
	Value Expression  // right side of the assignment
}

// TokenLiteral returns the literal of the token it is associated with.
func (ls *LetStatement) TokenLiteral() string {
	return ls.Token.Literal
}

// String returns the string representation of a statement.
func (ls *LetStatement) String() string {
	var out bytes.Buffer

	out.WriteString(ls.TokenLiteral() + " ")
	out.WriteString(ls.Name.String())
	out.WriteString(" = ")

	if ls.Value != nil {
		out.WriteString(ls.Value.String())
	}

	out.WriteString(";")

	return out.String()
}

// ReturnStatement represents the ast node for return statements.
type ReturnStatement struct {
	Token       token.Token
	ReturnValue Expression
}

// TokenLiteral represents the Literal value of the token that is associated with this statement.
func (rs *ReturnStatement) TokenLiteral() string {
	return rs.Token.Literal
}

// String returns the string representation of a statement.
func (rs *ReturnStatement) String() string {
	var out bytes.Buffer

	out.WriteString(rs.TokenLiteral() + " ")
	if rs.ReturnValue != nil {
		out.WriteString(rs.ReturnValue.String())
	}
	out.WriteString(";")

	return out.String()
}

// ExpressionStatement represents statements that consist solely of one expression.
type ExpressionStatement struct {
	Token      token.Token
	Expression Expression
}

// TokenLiteral returns the Literal of the token it is associated with.
func (es *ExpressionStatement) TokenLiteral() string {
	return es.Token.Literal
}

func (es *ExpressionStatement) String() string {
	if es.Expression != nil {
		return es.Expression.String()
	}

	return ""
}

// Implement the Statement interface
func (ls *LetStatement) statementNode()        {}
func (rs *ReturnStatement) statementNode()     {}
func (es *ExpressionStatement) statementNode() {}
