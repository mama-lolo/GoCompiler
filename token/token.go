package token

type TokenType string
type Token struct {
	Type    TokenType
	Literal string
}
 const {
	ILLEGAL = "ILLEGAL"
	EOF = "EOF"

	//Identifiers/Literals
	IDENTIFIER = "IDENTIFIER"
	INTEGER = "INTEGER"

	//Operators
	ASSIGN = "="
	PLUS = "+"

	//Delimiters
	COMMA = ","
	SEMICOLON = ";"

	PAREN_RIGHT = ")"
	PAREN_LEFT = "("
	
 }