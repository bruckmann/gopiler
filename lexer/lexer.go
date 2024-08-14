package lexer

import (
	"github.com/bruckmann/gopiler/enums"
	e "github.com/bruckmann/gopiler/enums"
)

type Lexer struct {
	input        string
	position     int
	readPosition int
	currentChar  byte
}

type guardianFunction func(byte) bool

func New(input string) *Lexer {
	l := &Lexer{input: input}
	l.readChar()

	return l
}

// This function has the responsability to get the next char to read
// Case we reach the end of the file set current char to zero (ASCII null)
func (l *Lexer) readChar() {
	l.currentChar = 0
	if l.readPosition >= len(l.input) {
	} else {
		l.currentChar = l.input[l.readPosition]
	}

	l.position = l.readPosition
	l.readPosition += 1
}

func (l *Lexer) newToken(tokenType e.TokenType, ch byte) e.Token {
	return e.Token{
		Type:    tokenType,
		Literal: string(ch),
	}
}

func (l *Lexer) newDoubleCharToken(tokenType e.TokenType, literal string) e.Token {
	return e.Token{
		Type:    tokenType,
		Literal: literal,
	}
}

func (l *Lexer) createDoubleCharToken(tokenType e.TokenType) e.Token {
		ch := l.currentChar
		l.readChar()
		literal := string(ch) + string(l.currentChar)
		return l.newDoubleCharToken(tokenType, literal) 
}

func (l *Lexer) isLetter(ch byte) bool {
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_'
}

func (l *Lexer) isDigit(ch byte) bool {
	return '0' <= ch && ch <= '9'
}

func (l *Lexer) readValue(gf guardianFunction) string {

	position := l.position

	for gf(l.currentChar) {
		l.readChar()
	}

	return l.input[position:l.position]
}

func (l *Lexer) peekChar() byte{

	if l.readPosition >= len(l.input){
		return 0
	} else {
		return l.input[l.readPosition]
	}

}

func (l *Lexer) eatWhitespaces() {
	for l.currentChar == ' ' ||
		l.currentChar == '\n' ||
		l.currentChar == '\t' ||
		l.currentChar == '\r' {
		l.readChar()
	}
}

func (l *Lexer) NextToken() e.Token {
	var token e.Token

	l.eatWhitespaces()

	switch l.currentChar {
	case '=':
		if(l.peekChar() == '='){
			token = l.createDoubleCharToken(e.EQUAL) 
		} else {
			token = l.newToken(e.ASSIGN, l.currentChar)
		}
	case '+':
		token = l.newToken(e.PLUS, l.currentChar)
	case '-':
		token = l.newToken(e.MINUS, l.currentChar)
	case '(':
		token = l.newToken(e.LEFT_PARENT, l.currentChar)
	case '{':
		token = l.newToken(e.LEFT_BRACE, l.currentChar)
	case ')':
		token = l.newToken(e.RIGHT_PARENT, l.currentChar)
	case '}':
		token = l.newToken(e.RIGHT_BRACE, l.currentChar)
	case ';':
		token = l.newToken(e.SEMICOLON, l.currentChar)
	case ',':
		token = l.newToken(e.COMMA, l.currentChar)
	case '/':
		token = l.newToken(e.SLASH, l.currentChar)
	case '!':
		if(l.peekChar() == '='){
			token = l.createDoubleCharToken(e.NOT_EQUAL)
		} else {
			token = l.newToken(e.BANG, l.currentChar)
		}
	case '>':
		token = l.newToken(e.GT, l.currentChar)
	case '<':
		token = l.newToken(e.LT, l.currentChar)
	case '*':
		token = l.newToken(e.ASTERISK, l.currentChar)
	case 0:
		token.Literal = ""
		token.Type = e.EOF
	default:
		if l.isLetter(l.currentChar) {
			token.Literal = l.readValue(l.isLetter)
			token.Type = e.IsKeywordOrIdentifier(token.Literal)
			return token
		} else if l.isDigit(l.currentChar) {
			  token.Literal = l.readValue(l.isDigit)
				token.Type = e.INT
				return token
		} else {
				token = l.newToken(enums.ILLEGAL, l.currentChar)
		}
	}

	l.readChar()
	return token
}
