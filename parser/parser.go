package parser

import (
	"github.com/viilis/ast"
	"github.com/viilis/lexer"
	"github.com/viilis/token"
)

type Parser struct {
	lexer        *lexer.Lexer
	currentToken token.Token
	peekToken    token.Token
}

func New(lexer *lexer.Lexer) *Parser {
	parser := &Parser{
		lexer: lexer,
	}

	parser.nextToken()

	return parser
}

func (parser *Parser) nextToken() {
	parser.currentToken = parser.peekToken
	parser.peekToken = parser.lexer.NextToken()
}

func (parser *Parser) ParseProgram() *ast.Program {
	program := &ast.Program{}
	program.Statements = []ast.Statement{}

	for parser.currentToken.Type != token.EOF {
		statement := parser.parseStatement()

		if statement != nil {
			program.Statements = append(program.Statements, statement)
		}

		parser.nextToken()
	}

	return program
}

func (parser *Parser) parseStatement() ast.Statement {
	switch parser.currentToken.Type {
	case token.LET:
		return parser.parseLetStatement()
	default:
		return nil
	}
}

func (parser *Parser) parseLetStatement() *ast.LetStatement {
	statement := &ast.LetStatement{Token: parser.currentToken}

	if !parser.expectPeek(token.IDENTIFIER) {
		return nil
	}

	statement.Name = &ast.Identifier{Token: parser.currentToken, Value: parser.currentToken.Literal}

	if !parser.expectPeek(token.ASSIGN) {
		return nil
	}

	// TODO: We're skipping the expressions until we
	// encounter a semicolon
	for !parser.curTokenIs(token.SEMICOLON) {
		parser.nextToken()
	}

	return statement
}

func (parser *Parser) curTokenIs(tokenType token.Tokentype) bool {
	return parser.currentToken.Type == tokenType
}

func (parser *Parser) peekTokenIs(tokenType token.Tokentype) bool {
	return parser.peekToken.Type == tokenType
}

func (parser *Parser) expectPeek(tokenType token.Tokentype) bool {
	if parser.peekTokenIs(tokenType) {
		parser.nextToken()

		return true
	} else {
		return false
	}
}
