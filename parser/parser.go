package parser

import (
	"fmt"
	"mama-lolo/GoCompiler/ast"
	"mama-lolo/GoCompiler/lexer"
	"mama-lolo/GoCompiler/token"
)

type Parser struct {
	l *lexer.Lexer

	currentToken token.Token
	peekedToken  token.Token
}

func New(l *lexer.Lexer) *Parser {
	p := &Parser{l: l}
	p.nextToken()
	p.nextToken()
	return p
}

func (p *Parser) nextToken() {
	p.currentToken = p.peekedToken
	p.peekedToken = p.l.NextToken()
}

func (p *Parser) ParseProgram() *ast.Program {
	fmt.Println("Started Parsing the Program")
	program := &ast.Program{}
	program.Statements = []ast.Statement{}
	for p.currentToken.Type != token.EOF {
		statement := p.parseStatement()
		if statement != nil {
			program.Statements = append(program.Statements, statement)
		}
		p.nextToken()
	}
	return program
}

func (p *Parser) parseStatement() ast.Statement {
	if p.currentToken.Type == token.LET {
		return p.parseLetStatement()
	} else {
		return nil
	}

}
func (p *Parser) parseLetStatement() *ast.LetStatement {
	statement := &ast.LetStatement{Token: p.currentToken}
	if !p.expectPeekedToken(token.IDENTIFIER) {
		return nil
	}
	statement.Name = &ast.Identifier{Token: p.currentToken, Value: p.currentToken.Literal}

	if !p.expectPeekedToken(token.ASSIGN) {
		return nil
	}
	//TODO the actual Expression does not get evaluated
	for !(p.currentToken.Type == token.SEMICOLON) {
		fmt.Printf("Passed token %v\n", p.currentToken)
		p.nextToken()
	}
	return statement
}

func (p *Parser) expectPeekedToken(t token.TokenType) bool {
	if p.peekedToken.Type == t {
		p.nextToken()
		return true
	}
	return false
}
