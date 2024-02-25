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
	ASSIGN    = "="
	PLUS      = "+"
	MINUS     = "-"
	BANG      = "!"
	ASTERISK  = "*"
	SLASH     = "/"
	LT        = "<"
	GT        = ">"
	EQUAL     = "=="
	NOT_EQUAL = "!="

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
	RETURN   = "RETURN"
	IF       = "IF"
	ELSE     = "ELSE"
	TRUE     = "TRUE"
	FALSE    = "FALSE"
)

// Maps keywords to proper token type
var keywords = map[string]Tokentype{
	"fn":     FUNCTION,
	"let":    LET,
	"return": RETURN,
	"if":     IF,
	"else":   ELSE,
	"true":   TRUE,
	"false":  FALSE,
}

// Separates keywords from identifiers.
func LookupIdentifier(ident string) Tokentype {
	tokType, ok := keywords[ident]

	if ok {
		return tokType
	}

	return IDENTIFIER
}
