package ast

import "github.com/efeckgz/Maymun/token"

// Node represents the nodes of the AST.
type Node interface {
	TokenLiteral() string
}

// Statement represents the statement nodes in the AST. It has one dummy method, and is only used to make writing easier.
type Statement interface {
	Node
	statementNode()
}

// Expression represents the expression nodes in the AST. It has one dummy method, and is only used to make writing easier.
type Expression interface {
	Node
	expressionNode()
}

// Program is the root node of every AST our parser produces.
type Program struct {
	// Evey Maymun program is a series of statements. These statements are contained in Program.Statements.
	Statements []Statement
}

// TokenLiteral for Program returns the Literal of the root node of the AST.
func (p *Program) TokenLiteral() string {
	if len(p.Statements) > 0 {
		return p.Statements[0].TokenLiteral()
	}

	return ""
}

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
