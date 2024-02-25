package token

type Tokentype string

type Token struct {
	Type    Tokentype
	Literal string
}

const (
	// Illegal tokens and end of line
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"

	// Indentifiers and literals
	IDENTIFIER = "IDENTIFIER"
	INT        = "INT"

	// Operators
	ASSIGN = "="
	PLUS   = "+"
	MINUS  = "-"

	// Delimiters
	COMMA     = ","
	SEMICOLON = ";"
	LPAREN    = "("
	RPAREN    = ")"
	LBRACE    = "{"
	RBRACE    = "}"

	// Keywords
	FUNCTION = "FUNCTION"
	LET      = "LET"
)
