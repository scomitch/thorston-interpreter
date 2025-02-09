package parser

import (
	"thorston-interpreter/ast"
	"thorston-interpreter/lexer"
	"thorston-interpreter/token"
)

// l is pointer to lexer instance 
// curToken and peekToken are the current and next tokens
type Parser struct {
	l *lexer.Lexer

	curToken token.Token
	peekToken token.Token
}

func New(l *lexer.Lexer) *Parser {
	p := &Parser{l: l}

	p.nextToken()
	p.nextToken()

	return p
}

// Helper to advance the tokens
func (p *Parser) nextToken(){
	p.curToken = p.peekToken
	p.peekToken = p.l.NextToken()
}

func (p *Parser) ParseProgram() *ast.Program {
	return nil
}