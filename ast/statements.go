package ast

import "github.com/efeckgz/Maymun/token"

// LetStatement represents the let statement node in the AST.
type LetStatement struct {
	Token token.Token // token.Let
	Name  *Identifier // name of the variable.
	Value Expression  // right side of the assignment
}

func (ls *LetStatement) statementNode() {} // dummy method

// TokenLiteral returns the literal of the token it is associated with.
func (ls *LetStatement) TokenLiteral() string {
	return ls.Token.Literal
}

// ReturnStatement represents the ast node for return statements.
type ReturnStatement struct {
	Token       token.Token
	ReturnValue Expression
}

func (rs *ReturnStatement) statementNode() {}

// TokenLiteral represents the Literal value of the token that is associated with this statement.
func (rs *ReturnStatement) TokenLiteral() string {
	return rs.Token.Literal
}

// ExpressionStatement represents statements that consist solely of one expression.
type ExpressionStatement struct {
	Token      token.Token
	Expression Expression
}

func (es *ExpressionStatement) statementNode() {}

// TokenLiteral returns the Literal of the token it is associated with.
func (es *ExpressionStatement) TokenLiteral() string {
	return es.Token.Literal
}
