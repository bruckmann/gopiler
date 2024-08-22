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

func (p *Parser) currentTokenIs(tokenType enums.TokenType) bool {
	return p.currentToken.Type == tokenType 
}



func (p *Parser) peekTokenIs(tokenType enums.TokenType) bool {
	return p.peekToken.Type == tokenType 
}

func (p *Parser) expectPeek(tokenType enums.TokenType) bool {

	if !p.peekTokenIs(tokenType) {
		return false
	}

	p.nextToken()
	return true
}

func (p *Parser) parseLetStatment() *ast.LetStatment {
	stmt := &ast.LetStatment{
		Token: p.currentToken,
	}

	if !p.expectPeek(enums.IDENTIFIER) {
		return nil
	}

	stmt.Name = &ast.Identifier{Token: p.currentToken, Value: p.currentToken.Literal}

	if !p.expectPeek(enums.ASSIGN) {
		return nil
	}

	for !p.currentTokenIs(enums.SEMICOLON) {
		p.nextToken()
	}

	return stmt
}

func (p *Parser) parseStatment() ast.Statment {
	switch p.currentToken.Type {
	case enums.LET:
		return p.parseLetStatment()
	default:
		return nil
		
	}
}

func (p *Parser) parseProgram() *ast.Program {
	program := &ast.Program{}

	for !p.currentTokenIs(enums.EOF) {
		stmt := p.parseStatment()

		if stmt != nil {
			program.Statments = append(program.Statments, stmt)
		}

		p.nextToken()
	}
	
	return program
}

func New(l *lexer.Lexer) *Parser {
	p := &Parser{l: l}

	p.nextToken()

	p.nextToken()

	return p
}
