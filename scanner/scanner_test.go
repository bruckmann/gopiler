package scanner

import (
	"testing"

	"github.com/bruckmann/gopiler/enums"
)

func TestNextToken(t *testing.T) {
	input := `let five = 5;

	 let ten = 10;

	 let add = fn(a, b) {
		a + b;
	};

	let result = add(five, ten);
	`

	tests := []struct {
		expectedType    enums.TokenType
		expectedLiteral string
	}{
		{enums.LET, "let"},
		{enums.IDENTIFIER, "five"},
		{enums.ASSIGN, "="},
		{enums.INT, "5"},
		{enums.SEMICOLON, ";"},
		{enums.LET, "let"},
		{enums.IDENTIFIER, "ten"},
		{enums.ASSIGN, "="},
		{enums.INT, "10"},
		{enums.SEMICOLON, ";"},
		{enums.LET, "let"},
		{enums.IDENTIFIER, "add"},
		{enums.ASSIGN, "="},
		{enums.FUNCTION, "fn"},
		{enums.LEFT_BRACE, "("},
		{enums.IDENTIFIER, "a"},
		{enums.COMMA, ","},
		{enums.IDENTIFIER, "b"},
		{enums.RIGHT_PARENT, ")"},
		{enums.LEFT_BRACE, "{"},
		{enums.INT, "5"},
		{enums.PLUS, "+"},
		{enums.INT, "10"},
		{enums.SEMICOLON, ";"},
		{enums.RIGHT_BRACE, "}"},
		{enums.SEMICOLON, ";"},
		{enums.LET, "let"},
		{enums.IDENTIFIER, "result"},
		{enums.ASSIGN, "="},
		{enums.IDENTIFIER, "add"},
		{enums.LEFT_BRACE, "("},
		{enums.IDENTIFIER, "five"},
		{enums.COMMA, ","},
		{enums.IDENTIFIER, "ten"},
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
