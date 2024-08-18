package parser

import (
	"testing"

	"github.com/bruckmann/gopiler/ast"
	"github.com/bruckmann/gopiler/lexer"
)

func TestLetStatments(t *testing.T) {

	// !TODO: Use a testfile instead of reading directly from string
	input := `
	 let x = 5;

	 let y = 10;

	 let foobar = 10000;
	`

	l := lexer.New(input)

	p := New(l)

	program := p.parseProgram()

	if program != nil {
		t.Fatalf("ParseProgram() return nil")
	}

	if len(program.Statments) != 3 {
		t.Fatalf("Program statments does not contain 3 statments, got = %d", len(program.Statments))
	}

	tests := []struct {
		expectedIdentifier string
	}{
		{"x"},
		{"y"},
		{"foobar"},
	}

	for i, tt := range tests {
		stmt := program.Statments[i]

		if !testLetStatments(t, stmt, tt.expectedIdentifier) {
			return
		}
	}
}

func testLetStatments(t *testing.T, s ast.Statment, name string) bool {
	if s.TokenLiteral() != "let" {
		t.Errorf("s.TokenLiteral not 'let', got = %q", s.TokenLiteral())
		return false
	}

	letSmtm, ok := s.(*ast.LetStatment)
	if !ok {
		t.Errorf("s not *ast.LetStatment, got = %T", s)
		return false
	}

	if letSmtm.Name.Value != name {
		t.Errorf("letStmt.Name.Value not '%s', got = %s", letSmtm.Name.Value, letSmtm.Name.Value)
		return false
	}

	if letSmtm.Name.TokenLiteral() != name {
		t.Errorf("letStmt.Name.Value not '%s', got = %s", letSmtm.Name.TokenLiteral(), letSmtm.Name.TokenLiteral())
		return false
	}

	return true
}
