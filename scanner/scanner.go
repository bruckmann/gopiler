package scanner

import (
	"github.com/bruckmann/gopiler/enums"
	e "github.com/bruckmann/gopiler/enums"
)

type Scanner struct {
	input        string
	position     int
	readPosition int
	currentChar  byte
}

func New(input string) *Scanner {
	s := &Scanner{input: input}
	s.readChar()

	return s
}

// This function has the responsability to get the next char to read
// Case we reach the end of the file set current char to zero (ASCII null)
func (s *Scanner) readChar() {
		s.currentChar = 0
	if s.readPosition >= len(s.input) {
	} else {
		s.currentChar = s.input[s.readPosition]
	}

	s.position = s.readPosition
	s.readPosition += 1
}

func (s *Scanner) newToken(tokenType e.TokenType, ch byte) e.Token {
	return e.Token{
		Type:    tokenType,
		Literal: string(ch),
	}
}

func (s *Scanner) isLetter(ch byte) bool {
	return 'a' <= ch && ch <= 'z'||'A' <= ch && ch <= 'Z'|| ch == '_'
}

func (s *Scanner) readIdentifier() string {
	position := s.position
	

	for s.isLetter(s.currentChar) {
			s.readChar()
	}

	return s.input[position:s.position]
}


func (s *Scanner) NextToken() e.Token {
	var token e.Token

	switch s.currentChar {
	case '=':
		token = s.newToken(e.ASSIGN, s.currentChar)
	case '+':
		token = s.newToken(e.PLUS, s.currentChar)
	case '-':
		token = s.newToken(e.MINUS, s.currentChar)
	case '(':
		token = s.newToken(e.LEFT_PARENT, s.currentChar)
	case '{':
		token = s.newToken(e.LEFT_BRACE, s.currentChar)
	case ')':
		token = s.newToken(e.RIGHT_PARENT, s.currentChar)
	case '}':
		token = s.newToken(e.RIGHT_BRACE, s.currentChar)
	case ';':
		token = s.newToken(e.SEMICOLON, s.currentChar)
	case ',':
		token = s.newToken(e.COMMA, s.currentChar)
	case '0':
		token.Literal = ""
		token.Type = e.EOF
	default: 
		if s.isLetter(s.currentChar){
			token.Literal = s.readIdentifier()
			token.Type = e.IsKeywordOrIdentifier(token.Literal)
		} else {
			token = s.newToken(enums.ILLEGAL, s.currentChar)
		}
	}

	s.readChar()
	return token
}
