package token

type TokenType string

const (
	// Single-character tokens.
	LEFT_PAREN  TokenType = "("
	RIGHT_PAREN TokenType = ")"
	LEFT_BRACE  TokenType = "{"
	RIGHT_BRACE TokenType = "}"
	COMMA       TokenType = ","
	DOT         TokenType = "."
	MINUS       TokenType = "-"
	PLUS        TokenType = "+"
	SEMICOLON   TokenType = ";"
	SLASH       TokenType = "/"
	STAR        TokenType = "*"

	// One or two character tokens.
	BANG          TokenType = "!"
	BANG_EQUAL    TokenType = "!="
	EQUAL         TokenType = "="
	EQUAL_EQUAL   TokenType = "=="
	GREATER       TokenType = ">"
	GREATER_EQUAL TokenType = ">="
	LESS          TokenType = "<"
	LESS_EQUAL    TokenType = "<="

	// Literals.
	IDENTIFIER TokenType = "IDENTIFIER"
	STRING     TokenType = "STRING"
	NUMBER     TokenType = "NUMBER"

	// Keywords.
	AND    TokenType = "and"
	CLASS  TokenType = "class"
	ELSE   TokenType = "else"
	FALSE  TokenType = "false"
	FUN    TokenType = "fun"
	FOR    TokenType = "for"
	IF     TokenType = "if"
	NIL    TokenType = "nil"
	OR     TokenType = "or"
	PRINT  TokenType = "print"
	RETURN TokenType = "return"
	SUPER  TokenType = "super"
	THIS   TokenType = "this"
	TRUE   TokenType = "true"
	VAR    TokenType = "var"
	WHILE  TokenType = "while"

	EOF TokenType = "EOF"
)
