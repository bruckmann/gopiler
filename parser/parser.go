package parser

import (
	"github.com/bruckmann/gopiler/ast"
	"github.com/bruckmann/gopiler/enums"
	"github.com/bruckmann/gopiler/lexer"
)

type Parser struct {
	l            *lexer.Lexer
	currentToken enums.Token
	peekToken    enums.Token
}

func (p *Parser) nextToken() {
	p.currentToken = p.peekToken

	p.peekToken = p.l.NextToken()
}

func (p *Parser) parseProgram() *ast.Program {
	return nil
}

func New(l *lexer.Lexer) *Parser {
	p := &Parser{l: l}

	p.nextToken()

	p.nextToken()

	return p
}
