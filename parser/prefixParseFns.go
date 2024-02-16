package parser

import (
	"fmt"
	"strconv"

	"github.com/efeckgz/Maymun/ast"
	"github.com/efeckgz/Maymun/token"
)

func (p *Parser) noPrefixParseFnError(t token.Type) {
	err := fmt.Sprintf("no prefix parse function for %s found.", t)
	p.errors = append(p.errors, err)
}

func (p *Parser) parseIdentifier() ast.Expression {
	return &ast.Identifier{Token: p.curToken, Value: p.curToken.Literal}
}

func (p *Parser) parseIntegerLiteral() ast.Expression {
	il := &ast.IntegerLiteral{Token: p.curToken}

	value, err := strconv.ParseInt(p.curToken.Literal, 0, 64)
	if err != nil {
		errMsg := fmt.Sprintf("could not parse %q as integer", p.curToken.Literal)
		p.errors = append(p.errors, errMsg)
		return nil
	}

	il.Value = value

	return il
}

func (p *Parser) parseFloatLiteral() ast.Expression {
	fl := &ast.FloatLiteral{Token: p.curToken}

	val, err := strconv.ParseFloat(p.curToken.Literal, 64)
	if err != nil {
		errMsg := fmt.Sprintf("could not parse %q as a float", p.curToken.Literal)
		p.errors = append(p.errors, errMsg)
		return nil
	}

	fl.Value = val
	return fl
}

func (p *Parser) parsePrefixExpression() ast.Expression {
	expression := &ast.PrefixExpression{
		Token:    p.curToken,
		Operator: p.curToken.Literal,
	}

	p.readToken()

	expression.Right = p.parseExpression(prefix)

	return expression
}
