package lexer

import (
	"bellamy/token"
	"bellamy/utils"
)

type Lexer struct {
	input        string
	position     int
	readPosition int
	ch           byte
}

func New(input string) *Lexer {
	l := &Lexer{input: input}
	l.readChar() // read in first byte to initialize
	return l
}

func (l *Lexer) NextToken() token.Token {
	var t token.Token

	l.skipWhitespace()

	switch l.ch {
	case '=':
		if l.peekChar() == '=' {
			ch := l.ch
			l.readChar()
			literal := string(ch) + string(l.ch)
			t = token.Token{Type: token.EQ, Literal: literal}
		} else {
			t = newToken(token.ASSIGN, l.ch)
		}
	case ';':
		t = newToken(token.SEMICOLON, l.ch)
	case ',':
		t = newToken(token.COMMA, l.ch)
	case '(':
		t = newToken(token.LPAREN, l.ch)
	case ')':
		t = newToken(token.RPAREN, l.ch)
	case '{':
		t = newToken(token.LBRACE, l.ch)
	case '}':
		t = newToken(token.RBRACE, l.ch)
	case '+':
		t = newToken(token.PLUS, l.ch)
	case '-':
		t = newToken(token.MINUS, l.ch)
	case '<':
		t = newToken(token.LT, l.ch)
	case '>':
		t = newToken(token.GT, l.ch)
	case '!':
		if l.peekChar() == '=' {
			ch := l.ch
			l.readChar()
			literal := string(ch) + string(l.ch)
			t = token.Token{Type: token.NE, Literal: literal}
		} else {
			t = newToken(token.BANG, l.ch)
		}
	case '*':
		t = newToken(token.ASTERISK, l.ch)
	case '/':
		t = newToken(token.SLASH, l.ch)
	case 0:
		t = newToken(token.EOF, '0')
	default:
		if utils.IsLetter(l.ch) {
			t.Literal = l.readIdentifier()
			t.Type = token.LookupIdent(t.Literal)
			return t
		} else if utils.IsDigit(l.ch) {
			t.Literal = l.readNumber()
			t.Type = token.INT
			return t
		} else {
			t = newToken(token.ILLEGAL, l.ch)
		}
	}
	l.readChar()
	return t
}

func newToken(tokenType token.TokenType, ch byte) token.Token {
	return token.Token{Type: tokenType, Literal: string(ch)}
}

func (l *Lexer) peekChar() byte {
	if l.readPosition >= len(l.input) {
		return 0
	}
	return l.input[l.readPosition]
}

func (l *Lexer) readChar() {
	l.ch = l.peekChar()
	l.position = l.readPosition
	l.readPosition += 1
}

func (l *Lexer) readIdentifier() string {
	position := l.position
	for utils.IsLetter(l.ch) {
		l.readChar()
	}
	return l.input[position:l.position]
}

func (l *Lexer) readNumber() string {
	position := l.position
	for utils.IsDigit(l.ch) {
		l.readChar()
	}
	return l.input[position:l.position]
}

func (l *Lexer) skipWhitespace() {
	for utils.IsWhitespace(l.ch) {
		l.readChar() // advance the token ahead
	}
}
