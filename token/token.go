package token

import "fmt"

type Token struct {
	TkType  TokenType
	Lexeme  string
	Literal any
	Line    int
}

func NewToken(tkType TokenType, lexeme string, literal any, line int) Token {
	token := Token{
		TkType:  tkType,
		Lexeme:  lexeme,
		Literal: literal,
		Line:    line,
	}

	return token
}

func (tk Token) ToString() string {
	return fmt.Sprintf("%v %s %v", tk.TkType, tk.Lexeme, tk.Literal)
}
