package ast

import "github.com/efeckgz/Maymun/token"

// Node represents the nodes of the AST.
type Node interface {
	// TokenLiteral returns the literal of the token it is associated with.
	TokenLiteral() string
}

// Statement represents the statement nodes in the AST.
type Statement interface {
	Node
	statementNode()
}

// Expression represents the expression nodes in the AST.
type Expression interface {
	Node
	expressionNode()
}

// Program is the root node of every AST our parser produces.
type Program struct {
	Statements []Statement
}

// TokenLiteral returns the literal of the token it is associated with.
func (p *Program) TokenLiteral() string {
	if len(p.Statements) > 0 {
		return p.Statements[0].TokenLiteral()
	}

	return ""
}

// LetStatement represents the let statement node in the AST.
type LetStatement struct {
	Token token.Token
	Name  *Identifier
	Value Expression
}

func (ls *LetStatement) statementNode() {}

// TokenLiteral returns the literal of the token it is associated with.
func (ls *LetStatement) TokenLiteral() string {
	return ls.Token.Literal
}

type Identifier struct {
	Token token.Token
	Value string
}

func (i *Identifier) expressionNode() {}

// TokenLiteral returns the literal of the token it is associated with.
func (i *Identifier) TokenLiteral() string {
	return i.Token.Literal
}
