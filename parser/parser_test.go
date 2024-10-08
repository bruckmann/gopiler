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
	checkParserErrors(t, p)

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

func TestReturnStatments(t *testing.T) {
	input := `
	return 5;

	return 10;

	return 1000;
	`
	l := lexer.New(input)

	p := New(l)

	program := p.parseProgram()
	checkParserErrors(t, p)

	if len(program.Statments) != 3 {
		t.Fatalf("Program.stataments does not contain 3 statments, got = %d",
			len(program.Statments))
	}

	for _, stmt := range program.Statments {
		returnStmt, ok := stmt.(*ast.ReturnStatment)

		if !ok {
			t.Errorf("Statment not *ast.ReturnStatment, got=%T", stmt)
			continue
		}

		if returnStmt.TokenLiteral() != "return" {
			t.Errorf("Return statment TokenLiteral not 'return', got=%q", returnStmt.TokenLiteral())
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

func checkParserErrors(t *testing.T, p *Parser) {
	errors := p.Errors()

	if len(errors) == 0 {
		return
	}

	t.Errorf("Parser has %d errors", len(errors))

	for _, msg := range errors {
		t.Errorf("Parser error: %q", msg)
	}

	t.FailNow()
}
