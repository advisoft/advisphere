package lox

import (
	"strconv"
)

type Lexer struct {
	source string
	tokens []Token

	start   int
	current int
	line    int

	keywords map[string]TokenType
}

func CreateLexer(source string) *Lexer {
	return &Lexer{source: source, tokens: make([]Token, 0), line: 1,
		keywords: map[string]TokenType{
			"and":    AND,
			"class":  CLASS,
			"else":   ELSE,
			"false":  FALSE,
			"for":    FOR,
			"fun":    FUN,
			"if":     IF,
			"nil":    NIL,
			"or":     OR,
			"print":  PRINT,
			"return": RETURN,
			"super":  SUPER,
			"this":   THIS,
			"true":   TRUE,
			"var":    VAR,
			"while":  WHILE,
		},
	}
}

func (lx *Lexer) ScanTokens() []Token {
	for !lx.isAtEnd() {
		lx.start = lx.current
		lx.scanToken()
	}

	lx.tokens = append(lx.tokens, Token{Type: EOF, Line: lx.line})

	return lx.tokens
}

func (lx *Lexer) scanToken() {
	c := lx.advance()
	switch c {
	case '(':
		lx.addToken(LEFT_PAREN)
	case ')':
		lx.addToken(RIGHT_PAREN)
	case '{':
		lx.addToken(LEFT_BRACE)
	case '}':
		lx.addToken(RIGHT_BRACE)
	case ',':
		lx.addToken(COMMA)
	case '.':
		lx.addToken(DOT)
	case '-':
		lx.addToken(MINUS)
	case '+':
		lx.addToken(PLUS)
	case ';':
		lx.addToken(SEMICOLON)
	case '*':
		lx.addToken(STAR)
	case '!':
		if lx.match('=') {
			lx.addToken(BANG_EQUAL)
		} else {
			lx.addToken(BANG)
		}
	case '=':
		if lx.match('=') {
			lx.addToken(EQUAL_EQUAL)
		} else {
			lx.addToken(EQUAL)
		}
	case '<':
		if lx.match('=') {
			lx.addToken(LESS_EQUAL)
		} else {
			lx.addToken(LESS)
		}
	case '>':
		if lx.match('=') {
			lx.addToken(GREATER_EQUAL)
		} else {
			lx.addToken(GREATER)
		}
	case '/':
		if lx.match('/') {
			// A comment goes until the end of the line
			for lx.peak() != '\n' && !lx.isAtEnd() {
				lx.advance()
			}
		} else {
			lx.addToken(SLASH)
		}

	case ' ':
		break // Ignore whitespace
	case '\r':
		break // Ignore whitespace
	case '\t':
		break // Ignore whitespace

	case '\n':
		lx.line++

	case '"':
		lx.string()

	default:
		if isDigit(c) {
			lx.number()
		} else if isAlpha(c) {
			lx.identifier()
		} else {
			lexError(lx.line, "Unexpected character.")
		}

	}
}

func (lx *Lexer) isAtEnd() bool {
	return lx.current >= len(lx.source)
}

func (lx *Lexer) advance() byte {
	b := lx.source[lx.current]
	lx.current++
	return b
}

func (lx *Lexer) match(expected byte) bool {
	if lx.isAtEnd() {
		return false
	}
	if lx.source[lx.current] != expected {
		return false
	}
	lx.current++
	return true
}

func (lx *Lexer) peak() byte {
	if lx.isAtEnd() {
		return 0
	}
	return lx.source[lx.current]
}

func (lx *Lexer) peakNext() byte {
	if lx.current+1 >= len(lx.source) {
		return 0
	}
	return lx.source[lx.current+1]
}

func (lx *Lexer) string() {
	for lx.peak() != '"' && !lx.isAtEnd() {
		if lx.peak() == '\n' {
			lx.line++
		}
		lx.advance()
	}

	if lx.isAtEnd() {
		lexError(lx.line, "Unterminated string.")
		return
	}

	// The closing ""
	lx.advance()

	value := lx.source[lx.start+1 : lx.current-1]
	lx.addTokenLiteral(STRING, value)
}

func isDigit(c byte) bool {
	return c >= '0' && c <= '9'
}

func isAlpha(c byte) bool {
	return (c >= 'a' && c <= 'z') || (c >= 'A' && c <= 'Z') || (c == '_')
}

func isAlphaNumeric(c byte) bool {
	return isAlpha(c) || isDigit(c)
}

func (lx *Lexer) number() {
	for isDigit(lx.peak()) {
		lx.advance()
	}

	if lx.peak() == '.' && isDigit(lx.peakNext()) {
		lx.advance()

		for isDigit(lx.peak()) {
			lx.advance()
		}
	}

	value, _ := strconv.ParseFloat(lx.source[lx.start:lx.current], 64)

	lx.addTokenLiteral(NUMBER, value)

}

func (lx *Lexer) identifier() {
	for isAlphaNumeric(lx.peak()) {
		lx.advance()
	}

	str := lx.source[lx.start:lx.current]

	val, ok := lx.keywords[str]
	if ok {
		lx.addToken(val)
	} else {
		lx.addToken(IDENTIFIER)
	}

}

func (lx *Lexer) addToken(tokenType TokenType) {
	lx.addTokenLiteral(tokenType, nil)
}

func (lx *Lexer) addTokenLiteral(tokenType TokenType, literal interface{}) {
	lx.tokens = append(lx.tokens, Token{Type: tokenType, Lexeme: lx.source[lx.start:lx.current], Literal: literal, Line: lx.line})
}
