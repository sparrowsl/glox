package scanner

import (
	"fmt"
	"glox/lox"
	"glox/token"
)

type Scanner struct {
	source  string
	tokens  []token.Token
	start   int
	current int
	line    int
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
		} else {
			sc.addToken(token.SLASH, nil)
		}
	case ' ':
	case '\r':
	case '\t':
		break
	case '\n':
		sc.line += 1
	default:
		lox.Error(sc.line, fmt.Sprintf("Unexpected character '%c'", char))
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
