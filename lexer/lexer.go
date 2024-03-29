package lexer

import "github.com/viilis/token"

type Lexer struct {
	input        string // Input aka give monkey code itself
	position     int    // Points to current character
	readPosition int    // Points to next character which allows "peeking"
	character    byte   // Current character
}

func New(input string) *Lexer {
	lexer := &Lexer{input: input}
	lexer.readChar()
	return lexer
}

// Turns current character to token and reads next one.
func (lexer *Lexer) NextToken() token.Token {
	var tok token.Token

	lexer.skipWhitespace()

	switch lexer.character {
	case '=':
		nextChar := lexer.peekChar()
		if nextChar == '=' {
			tokenLiteral := string(lexer.character) + string(nextChar)
			lexer.readChar()
			tok = token.Token{Type: token.EQUAL, Literal: tokenLiteral}
		} else {
			tok = newToken(token.ASSIGN, lexer.character)
		}
	case ';':
		tok = newToken(token.SEMICOLON, lexer.character)
	case '(':
		tok = newToken(token.LPAREN, lexer.character)
	case ')':
		tok = newToken(token.RPAREN, lexer.character)
	case ',':
		tok = newToken(token.COMMA, lexer.character)
	case '+':
		tok = newToken(token.PLUS, lexer.character)
	case '-':
		tok = newToken(token.MINUS, lexer.character)
	case '{':
		tok = newToken(token.LBRACE, lexer.character)
	case '}':
		tok = newToken(token.RBRACE, lexer.character)
	case '!':
		nextChar := lexer.peekChar()
		if nextChar == '=' {
			tokenLiteral := string(lexer.character) + string(nextChar)
			lexer.readChar()
			tok = token.Token{Type: token.NOT_EQUAL, Literal: tokenLiteral}
		} else {
			tok = newToken(token.BANG, lexer.character)
		}
	case '*':
		tok = newToken(token.ASTERISK, lexer.character)
	case '/':
		tok = newToken(token.SLASH, lexer.character)
	case '<':
		tok = newToken(token.LT, lexer.character)
	case '>':
		tok = newToken(token.GT, lexer.character)
	case 0:
		tok.Literal = ""
		tok.Type = token.EOF
	default:
		if isLetter(lexer.character) {
			tok.Literal = lexer.readIdentifiers()
			tok.Type = token.LookupIdentifier(tok.Literal)
			return tok
		} else if isNumber(lexer.character) {
			tok.Literal = lexer.readNumber()
			tok.Type = token.INT
			return tok
		} else {
			tok = newToken(token.ILLEGAL, lexer.character)
		}
	}

	lexer.readChar()

	return tok
}

func (l *Lexer) peekChar() byte {
	if l.readPosition >= len(l.input) {
		return 0
	} else {
		return l.input[l.readPosition]
	}
}

// Reads current character and sets positions ready for reading next one.
// Supports only ASCII. For UTF-8 support, implement lexer to support runes and read of multiple bytes
func (lexer *Lexer) readChar() {
	if lexer.readPosition >= len(lexer.input) {
		lexer.character = 0 // ASCII "NUL"
	} else {
		lexer.character = lexer.input[lexer.readPosition]
	}

	lexer.position = lexer.readPosition
	lexer.readPosition += 1
}

func newToken(tokenType token.Tokentype, character byte) token.Token {
	return token.Token{Type: tokenType, Literal: string(character)}
}

//TODO: Generalize reading indentifiers and numbers.

// Used for reading indentifiers from given code. Supports ASCII identifier names (a-z, A-Z and '_')
func (lexer *Lexer) readIdentifiers() string {
	position := lexer.position

	for isLetter(lexer.character) {
		lexer.readChar()
	}

	return lexer.input[position:lexer.position]
}

// Used for reading integers from given code.
func (lexer *Lexer) readNumber() string {
	position := lexer.position

	for isNumber(lexer.character) {
		lexer.readChar()
	}

	return lexer.input[position:lexer.position]
}

// Helper func for reading indentifiers. Checks if give argument is letter between a-z, A-Z or _
func isLetter(char byte) bool {
	return 'a' <= char && char <= 'z' || 'A' <= char && char <= 'Z' || char == '_'
}

// Consumes whitespaces.
func (lexer *Lexer) skipWhitespace() {
	for lexer.character == ' ' || lexer.character == '\t' || lexer.character == '\n' || lexer.character == '\r' {
		lexer.readChar()
	}
}

// Helper func for reading numbers.
func isNumber(char byte) bool {
	return '0' <= char && char <= '9'
}
