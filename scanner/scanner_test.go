package scanner

import (
	"fmt"
	"os"
	"testing"

	"github.com/bruckmann/gopiler/enums"
)

func TestNextToken(t *testing.T) {
	filePath := "../testfiles/tokentest.sm"
	input, err := os.ReadFile(filePath)

	if err != nil {
		t.Fatal(err)
	}

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
		{enums.LEFT_PARENT, "("},
		{enums.IDENTIFIER, "a"},
		{enums.COMMA, ","},
		{enums.IDENTIFIER, "b"},
		{enums.RIGHT_PARENT, ")"},
		{enums.LEFT_BRACE, "{"},
		{enums.IDENTIFIER, "a"},
		{enums.PLUS, "+"},
		{enums.IDENTIFIER, "b"},
		{enums.SEMICOLON, ";"},
		{enums.RIGHT_BRACE, "}"},
		{enums.SEMICOLON, ";"},
		{enums.LET, "let"},
		{enums.IDENTIFIER, "result"},
		{enums.ASSIGN, "="},
		{enums.IDENTIFIER, "add"},
		{enums.LEFT_PARENT, "("},
		{enums.IDENTIFIER, "five"},
		{enums.COMMA, ","},
		{enums.IDENTIFIER, "ten"},
		{enums.RIGHT_PARENT, ")"},
		{enums.SEMICOLON, ";"},
		{enums.BANG, "!"},
		{enums.MINUS, "-"},
		{enums.SLASH, "/"},
		{enums.ASTERISK, "*"},
		{enums.INT, "5"},
		{enums.SEMICOLON, ";"},
		{enums.INT, "5"},
		{enums.LT, "<"},
		{enums.INT, "10"},
		{enums.GT, ">"},
		{enums.INT, "5"},
		{enums.SEMICOLON, ";"},
		{enums.IF, "if"},
		{enums.LEFT_PARENT, "("},
		{enums.INT, "5"},
		{enums.LT, "<"},
		{enums.INT, "10"},
		{enums.RIGHT_PARENT, ")"},
		{enums.LEFT_BRACE, "{"},
		{enums.RETURN, "return"},
		{enums.TRUE, "true"},
		{enums.SEMICOLON, ";"},
		{enums.RIGHT_BRACE, "}"},
		{enums.ELSE, "else"},
		{enums.LEFT_BRACE, "{"},
		{enums.RETURN, "return"},
		{enums.FALSE, "false"},
		{enums.SEMICOLON, ";"},
		{enums.RIGHT_BRACE, "}"},
		{enums.INT, "10"},
		{enums.EQUAL, "=="},
		{enums.INT, "10"},
		{enums.SEMICOLON, ";"},
		{enums.INT, "9"},
		{enums.NOT_EQUAL, "!="},
		{enums.INT, "10"},
		{enums.SEMICOLON, ";"},
		{enums.EOF, ""},
	}

	s := New(string(input))

	for i, tt := range tests {
		token := s.NextToken()

		if token.Type != tt.expectedType {
			fmt.Println(token.Literal)
			t.Fatalf("test %d: expected type %q, got %q", i, tt.expectedType, token.Type)
		}

		if token.Literal != tt.expectedLiteral {
			t.Fatalf("test %d: expected literal %q, got %q", i, tt.expectedLiteral, token.Literal)
		}

	}

}
