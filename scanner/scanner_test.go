package scanner

import (
	"testing"

	"github.com/bruckmann/gopiler/enums"
)

func TestNextToken(t *testing.T) {
	input := `=+(){},;`

	tests := []struct {
		expectedType    enums.TokenType
		expectedLiteral string
	}{
		{enums.ASSIGN, "="},
		{enums.PLUS, "+"},
		{enums.LEFT_BRACE, "("},
		{enums.RIGHT_BRACE, ")"},
		{enums.LEFT_PARENT, "("},
		{enums.RIGHT_PARENT, ")"},
		{enums.COMMA, ","},
		{enums.SEMICOLON, ";"},
		{enums.EOF, ""},
	}

	s := New(input)

	for i, tt := range tests {
		token := s.NextToken()

		if token.Type != tt.expectedType {
			t.Fatalf("test %d: expected type %q, got %q", i, tt.expectedType, token.Type)
		}

		if token.Literal != tt.expectedLiteral {
			t.Fatalf("test %d: expected literal %q, got %q", i, tt.expectedLiteral, token.Literal)
		}

	}

}
