package ast

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
