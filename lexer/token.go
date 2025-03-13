package lexer

type TokenType string

const (
	MEOWSTART TokenType = "meowstart"
	MEOWEND   TokenType = "meowend"
	MEOW      TokenType = "meow"
	PURR      TokenType = "purr"
	IDENT     TokenType = "IDENT"
	NUMBER    TokenType = "NUMBER"
	STRING    TokenType = "STRING"
	OPERATOR  TokenType = "OPERATOR"
	ASSIGN    TokenType = "ASSIGN"
	SEMICOLON TokenType = "SEMICOLON"
	LPAREN    TokenType = "LPAREN"
	RPAREN    TokenType = "RPAREN"
	EOF       TokenType = "EOF"
)
