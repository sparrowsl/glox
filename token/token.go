package token

import "fmt"

type Token struct {
	tokenType TokenType
	lexeme    string
	literal   any
	line      int
}

func NewToken(tkType TokenType, lexeme string, literal any, line int) Token {
	token := Token{
		tokenType: tkType,
		lexeme:    lexeme,
		literal:   literal,
		line:      line,
	}

	return token
}

func (tk Token) ToString() string {
	return fmt.Sprintf("%v %s %v", tk.tokenType, tk.lexeme, tk.literal)
}
