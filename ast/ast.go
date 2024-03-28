package ast

import "github.com/viilis/token"

type Node interface {
	TokenLiteral() string // Used for debugging
}

type Statement interface {
	Node
	statementNode()
}

type Expression interface {
	Node
	expressionNode()
}

type Program struct {
	Statements []Statement
}

func (program *Program) TokenLiteral() string {
	if len(program.Statements) > 0 {
		return program.Statements[0].TokenLiteral()
	} else {
		return ""
	}
}

type LetStatement struct {
	Token token.Token // token.LET
	Name  *Identifier
	Value Expression
}

func (letState *LetStatement) statementNode() {}
func (letState *LetStatement) TokenLiteral() string {
	return letState.Token.Literal
}

type Identifier struct {
	Token token.Token // token.IDENTIFIER
	Value string
}

func (identifier *Identifier) expressionNode() {}
func (identifier *Identifier) TokenLiteral() string {
	return identifier.Token.Literal
}
