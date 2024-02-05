package parser

import (
	"fmt"

	"github.com/efeckgz/Maymun/ast"
	"github.com/efeckgz/Maymun/lexer"
	"github.com/efeckgz/Maymun/token"
)

const (
	_ int = iota
	lowest
	equals           // ==
	comparisonEquals // >= or <=
	comparison       // > or <
	sum              // +
	product          // *
	prefix           // -x or !x
	call             // square(x)
)

type prefixParseFn func() ast.Expression

// Parser represents the inner state of the parser during the parsing of the source code.
type Parser struct {
	l         *lexer.Lexer // Lexer to lex the source code.
	curToken  token.Token  // The current token the parses is evaluating.
	nextToken token.Token  // The next token the parser will evaluate.
	errors    []string     // string array to hold parsing errors.

	prefixParseFns map[token.Type]prefixParseFn
}

// New creates a new parser.
func New(l *lexer.Lexer) *Parser {
	p := &Parser{l: l, errors: []string{}}

	p.prefixParseFns = make(map[token.Type]prefixParseFn)

	p.registerPrefix(token.Ident, p.parseIdentifier)
	p.registerPrefix(token.Int, p.parseIntegerLiteral)
	p.registerPrefix(token.Float, p.parseFloatLiteral)
	p.registerPrefix(token.Bang, p.parsePrefixExpression)
	p.registerPrefix(token.Minus, p.parsePrefixExpression)

	// Read two tokens so that both curToken and peekToken are set
	p.readToken()
	p.readToken()

	return p
}

// ParseProgram generates the AST from an input program
func (p *Parser) ParseProgram() *ast.Program {
	program := &ast.Program{}
	program.Statements = []ast.Statement{}

	for !p.curTokenIs(token.EOF) {
		stmt := p.parseStatement()
		if stmt != nil {
			program.Statements = append(program.Statements, stmt)
		}

		p.readToken()
	}

	return program
}

// Errors gives the errors encountered at the time of calling.
func (p *Parser) Errors() []string {
	return p.errors
}

func (p *Parser) readToken() {
	p.curToken = p.nextToken
	p.nextToken = p.l.NextToken()
}

// readSpesificToken only reads if the parameter token is the next, and adds an error if otherwise.
func (p *Parser) readSpesificToken(t token.Type) bool {
	if p.nextTokenIs(t) {
		p.readToken()
		return true
	}

	p.peekError(t)
	return false
}

func (p *Parser) curTokenIs(t token.Type) bool {
	return p.curToken.Type == t
}

func (p *Parser) nextTokenIs(t token.Type) bool {
	return p.nextToken.Type == t
}

func (p *Parser) peekError(t token.Type) {
	err := fmt.Sprintf("Unexpected next token: expected %s, got %s.", t, p.nextToken.Type)
	p.errors = append(p.errors, err)
}

func (p *Parser) registerPrefix(tokenType token.Type, fn prefixParseFn) {
	p.prefixParseFns[tokenType] = fn
}

func (p *Parser) parseStatement() ast.Statement {
	switch p.curToken.Type {
	case token.Let:
		return p.parseLetStatement()
	case token.Return:
		return p.parseReturnStatement()
	default:
		return p.parseExpressionStatement()
	}
}

func (p *Parser) parseLetStatement() *ast.LetStatement {
	stmt := &ast.LetStatement{Token: p.curToken}

	if !p.readSpesificToken(token.Ident) {
		return nil // If the next token after a token.Let is not an identifier return early.
	}
	// p.readToken()

	stmt.Name = &ast.Identifier{Token: p.curToken, Value: p.curToken.Literal}

	if !p.readSpesificToken(token.Assign) {
		return nil // If the next token after variable name is not '=', return early.
	}
	// p.readToken()

	for !p.curTokenIs(token.Semicolon) {
		p.readToken() // TODO: We are skipping the expression to the right for now.
	}

	return stmt
}

func (p *Parser) parseReturnStatement() *ast.ReturnStatement {
	rs := &ast.ReturnStatement{Token: p.curToken}
	p.readToken()

	for !p.curTokenIs(token.Semicolon) {
		p.readToken()
	}

	return rs
}

func (p *Parser) parseExpressionStatement() *ast.ExpressionStatement {
	es := &ast.ExpressionStatement{Token: p.curToken}
	es.Expression = p.parseExpression(lowest)

	if p.nextTokenIs(token.Semicolon) {
		p.readToken()
	}

	return es
}

func (p *Parser) parseExpression(precedence int) ast.Expression {
	prefix := p.prefixParseFns[p.curToken.Type]
	if prefix == nil {
		p.noPrefixParseFnError(p.curToken.Type)
		return nil
	}

	leftExp := prefix()
	return leftExp
}
