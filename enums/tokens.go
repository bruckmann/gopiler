package enums

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}

const (
	// Defining keyword tokens
	FUNCTION = "FUNCTION"
	LET      = "LET"

	// Defining Identifier and Literals
	IDENTIFIER = "IDENTIFIER"
	INT        = "INT"

	// Defining Single char tokens
	ASSIGN       = "="
	COMMA        = ","
	DOT          = "."
	LEFT_BRACE   = "{"
	LEFT_PARENT  = "("
	MINUS        = "-"
	PLUS         = "+"
	RIGHT_BRACE  = "}"
	RIGHT_PARENT = ")"
	SEMICOLON    = ":"

	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"
)
