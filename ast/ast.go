package ast

import "github.com/bruckmann/gopiler/enums"

// !TODO: Refactor to use small modules for each type

type Node interface {
	TokenLiteral() string
}

type Statment interface {
	Node

	statmentNode()
}

type Expression interface {
	Node

	expressionNode()
}

type Identifier struct {
	Token enums.Token
	Value string
}

func (i *Identifier) expressionNode() {}

func (i *Identifier) TokenLiteral() string {
	return i.Token.Literal
}

type LetStatment struct {
	Token enums.Token
	Name  *Identifier
	Value Expression
}

func (ls *LetStatment) statmentNode() {}

func (ls *LetStatment) TokenLiteral() string {
	return ls.Token.Literal
}

type ReturnStatment struct {
	Token       enums.Token
	ReturnValue Expression
}

func (rs *ReturnStatment) statmentNode() {}

func (rs *ReturnStatment) TokenLiteral() string {
	return rs.Token.Literal
}

type Program struct {
	Statments []Statment
}

func (p *Program) TokenLiteral() string {
	if len(p.Statments) <= 0 {
		return ""
	}

	return p.Statments[0].TokenLiteral()
}
