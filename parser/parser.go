package parser

import (
	"github.com/efeckgz/Maymun/ast"
	"github.com/efeckgz/Maymun/lexer"
	"github.com/efeckgz/Maymun/token"
)

// Parser represents the inner state of the parser during the parsing of the source code.
type Parser struct {
	l         *lexer.Lexer // Lexer to lex the source code.
	curToken  token.Token  // The current token the parses is evaluating.
	nextToken token.Token  // The next token the parser will evaluate.
}

// New creates a new parser.
func New(l *lexer.Lexer) *Parser {
	p := &Parser{l: l}

	// Read two tokens so that both curToken and peekToken are set
	p.readToken()
	p.readToken()

	return p
}

// ParseProgram generates the AST from an input program
func (p *Parser) ParseProgram() *ast.Program {
	program := &ast.Program{}
	program.Statements = []ast.Statement{}

	for p.curToken.Type != token.EOF {
		stmt := p.parseStatement()
		if stmt != nil {
			program.Statements = append(program.Statements, stmt)
		}

		p.readToken()
	}

	return program
}

func (p *Parser) readToken() {
	p.curToken = p.nextToken
	p.nextToken = p.l.NextToken()
}

func (p *Parser) curTokenIs(t token.Type) bool {
	return p.curToken.Type == t
}

func (p *Parser) nextTokenIs(t token.Type) bool {
	return p.nextToken.Type == t
}

func (p *Parser) expectNext(t token.Type) bool {
	if p.nextTokenIs(t) {
		p.readToken()
		return true
	}

	return false
}

func (p *Parser) parseStatement() ast.Statement {
	switch p.curToken.Type {
	case token.Let:
		return p.parseLetStatement()
	default:
		return nil
	}
}

func (p *Parser) parseLetStatement() *ast.LetStatement {
	stmt := &ast.LetStatement{Token: p.curToken}

	if !p.expectNext(token.Ident) {
		return nil // If the next token after a token.Let is not an identifier return early.
	}

	stmt.Name = &ast.Identifier{Token: p.curToken, Value: p.curToken.Literal}

	if !p.expectNext(token.Assign) {
		return nil // If the next token after variable name is not '=', return early.
	}

	for !p.curTokenIs(token.Semicolon) {
		p.readToken() // TODO: We are skipping the expression to the right for now.
	}

	return stmt
}