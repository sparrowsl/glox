package scanner

import (
	"fmt"
	"glox/lox"
	"glox/token"
	"strconv"
)

type Scanner struct {
	source  string
	tokens  []token.Token
	start   int
	current int
	line    int
}

var keywords = map[string]token.TokenType{
	"and":    token.AND,
	"class":  token.CLASS,
	"else":   token.ELSE,
	"false":  token.FALSE,
	"for":    token.FOR,
	"fun":    token.FUN,
	"if":     token.IF,
	"nil":    token.NIL,
	"or":     token.OR,
	"print":  token.PRINT,
	"return": token.RETURN,
	"super":  token.SUPER,
	"this":   token.THIS,
	"true":   token.TRUE,
	"var":    token.VAR,
	"while":  token.WHILE,
}

func NewScanner(source string) Scanner {
	scanner := Scanner{
		source:  source,
		tokens:  []token.Token{},
		start:   0,
		current: 0,
		line:    1,
	}

	return scanner
}

func (sc *Scanner) ScanTokens() []token.Token {
	for !sc.isAtEnd() {
		sc.start = sc.current
		sc.scanToken()
	}

	sc.tokens = append(sc.tokens, token.NewToken(token.EOF, "", nil, sc.line))

	return sc.tokens
}

func (sc *Scanner) isAtEnd() bool {
	return sc.current >= len(sc.source)
}

func (sc *Scanner) scanToken() {
	char := sc.advance()

	switch char {
	case '(':
		sc.addToken(token.LEFT_PAREN, nil)
	case ')':
		sc.addToken(token.RIGHT_PAREN, nil)
	case '{':
		sc.addToken(token.LEFT_BRACE, nil)
	case '}':
		sc.addToken(token.RIGHT_BRACE, nil)
	case ',':
		sc.addToken(token.COMMA, nil)
	case '.':
		sc.addToken(token.DOT, nil)
	case '-':
		sc.addToken(token.MINUS, nil)
	case '+':
		sc.addToken(token.PLUS, nil)
	case ';':
		sc.addToken(token.SEMICOLON, nil)
	case '*':
		sc.addToken(token.STAR, nil)
	case '!':
		if sc.match('=') {
			sc.addToken(token.BANG_EQUAL, nil)
		} else {
			sc.addToken(token.BANG, nil)
		}
	case '=':
		if sc.match('=') {
			sc.addToken(token.EQUAL_EQUAL, nil)
		} else {
			sc.addToken(token.EQUAL, nil)
		}
	case '<':
		if sc.match('=') {
			sc.addToken(token.LESS_EQUAL, nil)
		} else {
			sc.addToken(token.LESS, nil)
		}
	case '>':
		if sc.match('=') {
			sc.addToken(token.GREATER_EQUAL, nil)
		} else {
			sc.addToken(token.GREATER, nil)
		}
	case '/':
		if sc.match('/') {
			for sc.peek() != '\n' && !sc.isAtEnd() {
				sc.advance()
			}
		} else if sc.match('*') {
			for !sc.isAtEnd() {
				if sc.peek() == '*' && sc.peekNext() == '/' {
					sc.advance()
					sc.advance()
					return
				}

				sc.advance()
			}
			// If we got here, comment wasn’t closed
			fmt.Println("Unterminated multi-line comment")
		} else {
			sc.addToken(token.SLASH, nil)
		}
	case '"':
		sc.string()
	case ' ':
	case '\r':
	case '\t':
		break
	case '\n':
		sc.line += 1
	default:
		if isDigit(char) {
			sc.number()
		} else if isAlpha(char) {
			sc.identifier()
		} else {

			lox.Error(sc.line, fmt.Sprintf("Unexpected character '%c'", char))
		}
	}
}

func (sc *Scanner) advance() byte {
	char := sc.source[sc.current]
	sc.current++
	return char
}

func (sc *Scanner) addToken(tkType token.TokenType, literal any) {
	text := sc.source[sc.start:sc.current]
	sc.tokens = append(sc.tokens, token.NewToken(tkType, text, literal, sc.line))
}

func (sc *Scanner) match(expected byte) bool {
	if sc.isAtEnd() {
		return false
	}

	if sc.source[sc.current] != expected {
		return false
	}

	sc.current += 1
	return true
}

func (sc *Scanner) peek() byte {
	if sc.isAtEnd() {
		return 0
	}

	return sc.source[sc.current]
}

func (sc *Scanner) string() {
	for sc.peek() != '"' && !sc.isAtEnd() {
		if sc.peek() == '\n' {
			sc.line += 1
		}

		sc.advance()
	}

	if sc.isAtEnd() {
		lox.Error(sc.line, "Unterminated string.")
		return
	}

	sc.advance()
	value := sc.source[sc.start+1 : sc.current-1]
	sc.addToken(token.STRING, value)
}

func isDigit(char byte) bool {
	return char >= '0' && char <= '9'
}

func (sc *Scanner) number() {
	for isDigit(sc.peek()) {
		sc.advance()
	}

	if sc.peek() == '.' && isDigit(sc.peekNext()) {
		sc.advance()

		for isDigit(sc.peek()) {
			sc.advance()
		}
	}

	value, _ := strconv.ParseFloat(sc.source[sc.start:sc.current], 64)
	sc.addToken(token.NUMBER, value)
}

func (sc Scanner) peekNext() byte {
	if sc.current+1 >= len(sc.source) {
		return 0
	}

	return sc.source[sc.current+1]

}

func (sc *Scanner) identifier() {
	for isAlphaNumeric(sc.peek()) {
		sc.advance()
	}

	text := sc.source[sc.start:sc.current]
	word, ok := keywords[text]

	if !ok {
		word = token.IDENTIFIER
	}

	sc.addToken(word, nil)
}

func isAlpha(char byte) bool {
	return (char >= 'a' && char <= 'z') || (char >= 'A' && char <= 'Z') || char == '_'
}

func isAlphaNumeric(char byte) bool {
	return isAlpha(char) || isDigit(char)
}
