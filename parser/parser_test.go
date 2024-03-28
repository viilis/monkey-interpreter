package parser

import (
	"testing"

	"github.com/viilis/ast"
	"github.com/viilis/lexer"
)

type Expected struct {
	expectedIdentifier string
}

// Should parse three let statement identifiers properly
func TestLetStatements(t *testing.T) {
	input := `
	let x = 5;
	let y = 10;
	let foobar = 123456;
	`

	lexer := lexer.New(input)
	parser := New(lexer)

	program := parser.ParseProgram()

	if program == nil {
		t.Fatalf("ParseProgram() returned nil")
	}

	if len(program.Statements) != 3 {
		t.Fatalf("program.Statements does not contain 3 statements. got=%d",
			len(program.Statements))
	}

	tests := []Expected{
		{"x"},
		{"y"},
		{"foobar"},
	}

	for i, testToken := range tests {
		statement := program.Statements[i]

		if !TestLetStatement(t, statement, testToken.expectedIdentifier) {
			return
		}
	}
}

// Should parse one individual let statement
func TestLetStatement(t *testing.T, statement ast.Statement, name string) bool {

	if statement.TokenLiteral() != "let" {
		t.Errorf("s.TokenLiteral not 'let'. got=%q", statement.TokenLiteral())
		return false
	}

	letStatement, ok := statement.(*ast.LetStatement)

	if !ok {
		t.Errorf("s not *ast.LetStatement. got=%T", statement)
		return false
	}

	if letStatement.Name.Value != name {
		t.Errorf("letStatement.Name.Value not '%s'. got=%s", name, letStatement.Name.Value)
		return false
	}

	if letStatement.Name.TokenLiteral() != name {
		t.Errorf("letStatement.Name.TokenLiteral() not '%s'. got=%s", name, letStatement.Name.TokenLiteral())
		return false
	}

	return true
}
