package lexer

import (
	"fmt"
	"testing"

	"github.com/viilis/token"
)

type expectedToken struct {
	expectedType    token.Tokentype
	expectedLiteral string
}

func TestNextToken(t *testing.T) {
	input := `
	let five = 5;
	let ten = 10;
	
	let add = fn(x, y) {
		x + y;
	};
	let result = add(five, ten);
	`

	tests := []expectedToken{
		{token.LET, "let"},
		{token.IDENTIFIER, "five"},
		{token.ASSIGN, "="},
		{token.INT, "5"},
		{token.SEMICOLON, ";"},
		{token.LET, "let"},
		{token.IDENTIFIER, "ten"},
		{token.ASSIGN, "="},
		{token.INT, "10"},
		{token.SEMICOLON, ";"},
		{token.LET, "let"},
		{token.IDENTIFIER, "add"},
		{token.ASSIGN, "="},
		{token.FUNCTION, "fn"},
		{token.LPAREN, "("},
		{token.IDENTIFIER, "x"},
		{token.COMMA, ","},
		{token.IDENTIFIER, "y"},
		{token.RPAREN, ")"},
		{token.LBRACE, "{"},
		{token.IDENTIFIER, "x"},
		{token.PLUS, "+"},
		{token.IDENTIFIER, "y"},
		{token.SEMICOLON, ";"},
		{token.RBRACE, "}"},
		{token.SEMICOLON, ";"},
		{token.LET, "let"},
		{token.IDENTIFIER, "result"},
		{token.ASSIGN, "="},
		{token.IDENTIFIER, "add"},
		{token.LPAREN, "("},
		{token.IDENTIFIER, "five"},
		{token.COMMA, ","},
		{token.IDENTIFIER, "ten"},
		{token.RPAREN, ")"},
		{token.SEMICOLON, ";"},
		{token.EOF, ""},
	}

	lexer := New(input)

	for i, testToken := range tests {
		token := lexer.NextToken()
		fmt.Println(token)

		if token.Type != testToken.expectedType {
			t.Fatalf(
				"tests[%d] - wrong tokentype. expected=%q got=%q",
				i,
				testToken.expectedType,
				token.Type,
			)
		}

		if token.Literal != testToken.expectedLiteral {
			t.Fatalf(
				"tests[%d - wrong literal. expected=%q got=%q",
				i,
				testToken.expectedLiteral,
				token.Literal,
			)
		}
	}
}
