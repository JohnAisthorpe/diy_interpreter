package token

type TokenType string

type Token struct {
	Type	TokenType
	Literal	string
}

const (
	ILLEGAL = "ILLEGAL"
	EOF		= "EOF"
	
	// identifiers + literals
	IDENT	= "IDENT" // add, foobar, x, y, ...
	INT		= "INT" // 134356

	//Operators
	ASSIGN	= "="
	PLUS	= "+"

	//DELIMITERS
	COMMA		= ","
	SEMICOLON	= ";"
	LPAREN 		= "("
	RPAREN		= ")"
	LBRACE		= "{"
	RBRACE		= "}"

	//KEYWORDS

	FUNCTION	= "FUNCTION"
	LET			= "LET"

)