package parser

import (
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

var precedences = map[token.Type]int{
	token.Eq:       equals,
	token.Noteq:    equals,
	token.Lt:       comparison,
	token.Gt:       comparison,
	token.Lteq:     comparisonEquals,
	token.Gteq:     comparisonEquals,
	token.Plus:     sum,
	token.Minus:    sum,
	token.Slash:    product,
	token.Asterisk: product,
}

func (p *Parser) peekPrecedence() int {
	if p, ok := precedences[p.nextToken.Type]; ok {
		return p
	}

	return lowest
}

func (p *Parser) curPrecedence() int {
	if p, ok := precedences[p.curToken.Type]; ok {
		return p
	}

	return lowest
}
