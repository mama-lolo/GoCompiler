package lexer

import (
	"fmt"
	"mama-lolo/GoCompiler/token"
	"testing"
)

func TestNextToken(t *testing.T) {

	input := `let five = 5;
			let ten = 10;
			let add = fn(x, y) {
			x + y;
			};
			let result = add(five, ten);
			!-/*5;
			5 < 10 > 5;
			if (5 < 10) {
				return true;
			} else {
				return false;
			}
			== != <= >=`
	tests := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		{token.LET, "let"},
		{token.IDENTIFIER, "five"},
		{token.ASSIGN, "="},
		{token.INTEGER, "5"},
		{token.SEMICOLON, ";"},
		{token.LET, "let"},
		{token.IDENTIFIER, "ten"},
		{token.ASSIGN, "="},
		{token.INTEGER, "10"},
		{token.SEMICOLON, ";"},
		{token.LET, "let"},
		{token.IDENTIFIER, "add"},
		{token.ASSIGN, "="},
		{token.FUNCTION, "fn"},
		{token.PAREN_LEFT, "("},
		{token.IDENTIFIER, "x"},
		{token.COMMA, ","},
		{token.IDENTIFIER, "y"},
		{token.PAREN_RIGHT, ")"},
		{token.CURLY_LEFT, "{"},
		{token.IDENTIFIER, "x"},
		{token.PLUS, "+"},
		{token.IDENTIFIER, "y"},
		{token.SEMICOLON, ";"},
		{token.CURLY_RIGHT, "}"},
		{token.SEMICOLON, ";"},
		{token.LET, "let"},
		{token.IDENTIFIER, "result"},
		{token.ASSIGN, "="},
		{token.IDENTIFIER, "add"},
		{token.PAREN_LEFT, "("},
		{token.IDENTIFIER, "five"},
		{token.COMMA, ","},
		{token.IDENTIFIER, "ten"},
		{token.PAREN_RIGHT, ")"},
		{token.SEMICOLON, ";"},
		{token.BANG, "!"},
		{token.MINUS, "-"},
		{token.SLASH, "/"},
		{token.ASTERISK, "*"},
		{token.INTEGER, "5"},
		{token.SEMICOLON, ";"},
		{token.INTEGER, "5"},
		{token.LT, "<"},
		{token.INTEGER, "10"},
		{token.GT, ">"},
		{token.INTEGER, "5"},
		{token.SEMICOLON, ";"},
		{token.IF, "if"},
		{token.PAREN_LEFT, "("},
		{token.INTEGER, "5"},
		{token.LT, "<"},
		{token.INTEGER, "10"},
		{token.PAREN_RIGHT, ")"},
		{token.CURLY_LEFT, "{"},
		{token.RETURN, "return"},
		{token.TRUE, "true"},
		{token.SEMICOLON, ";"},
		{token.CURLY_RIGHT, "}"},
		{token.ELSE, "else"},
		{token.CURLY_LEFT, "{"},
		{token.RETURN, "return"},
		{token.FALSE, "false"},
		{token.SEMICOLON, ";"},
		{token.CURLY_RIGHT, "}"},
		{token.EQUAL, "=="},
		{token.NOT_EQUAL, "!="},
		{token.LEQ, "<="},
		{token.GEQ, ">="},
		{token.EOF, ""},
	}

	l := New(input)
	for i, testCase := range tests {
		tok := l.NextToken()
		fmt.Printf("Current Token: %q\n", tok)
		if tok.Type != testCase.expectedType {
			t.Fatalf("tests[%d] - wrong token type. Expected : %q, but got:%q",
				i, testCase.expectedType, tok.Type)
		}
		if tok.Literal != testCase.expectedLiteral {
			t.Fatalf("tests[%d] - wrong literal. Expected : %q, but got:%q",
				i, testCase.expectedLiteral, tok.Literal)
		}
	}
}
