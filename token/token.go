package token

type TokenType string
type Token struct {
	Type    TokenType
	Literal string
}

const (
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"

	//Identifiers/Literals
	IDENTIFIER = "IDENTIFIER"
	INTEGER    = "INTEGER"

	//Operators
	ASSIGN = "="
	PLUS   = "+"

	//Delimiters
	COMMA     = ","
	SEMICOLON = ";"

	PAREN_RIGHT = ")"
	PAREN_LEFT  = "("
	CURLY_RIGHT = "}"
	CURLY_LEFT  = "{"

	// Keywords
	FUNCTION = "FUNCTION"
	LET      = "LET"
)

var keywords = map[string]TokenType{
	"fn":  FUNCTION,
	"let": LET,
}

func LookupIdent(ident string) TokenType {
	tok, ok := keywords[ident]
	if ok {
		return tok
	}
	return IDENTIFIER
}
