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
	ASSIGN   = "="
	PLUS     = "+"
	MINUS    = "-"
	BANG     = "!"
	ASTERISK = "*"
	SLASH    = "/"
	LT       = "<"
	GT       = ">"
	// DoubleSpae Operators
	EQUAL     = "=="
	NOT_EQUAL = "!="
	LEQ       = "<="
	GEQ       = ">="

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
	TRUE     = "true"
	FALSE    = "false"
	IF       = "if"
	ELSE     = "else"
	RETURN   = "return"
)

var keywords = map[string]TokenType{
	"fn":     FUNCTION,
	"let":    LET,
	"true":   TRUE,
	"false":  FALSE,
	"if":     IF,
	"else":   ELSE,
	"return": RETURN,
}

func LookupIdent(ident string) TokenType {
	tok, ok := keywords[ident]
	if ok {
		return tok
	}
	return IDENTIFIER
}
