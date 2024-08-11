package enums

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}

const (
	// Defining keyword tokens
	FUNCTION = "FUNCTION"
	IF       = "IF"
	RETURN   = "RETURN"
	ELSE     = "ELSE"
	LET      = "LET"
	TRUE     = "TRUE"
	FALSE    = "FALSE"

	// Defining Identifier and Literals
	IDENTIFIER = "IDENTIFIER"
	INT        = "INT"
	BOOL       = "BOOL"

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
	GT           = ">"
	LT           = "<"
	SLASH        = "/"
	ASTERISK     = "*"
	BANG         = "!"

	// Defining double char tokens
	EQUAL     = "=="
	NOT_EQUAL = "!="

	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"
)

var KEYWORDS = map[string]TokenType{
	"fn":     FUNCTION,
	"let":    LET,
	"return": RETURN,
	"if":     IF,
	"else":   ELSE,
	"true":   TRUE,
	"false":  FALSE,
}

func IsKeywordOrIdentifier(value string) TokenType {
	if token, ok := KEYWORDS[value]; ok {
		return token
	}
	return IDENTIFIER
}
