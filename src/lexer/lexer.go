package lexer

import (
	"itstejas.com/monkey-go/src/token"
)

type Lexer struct {
	input        string
	position     int  // Current position in input (points to the current char)
	readPosition int  // Current reading position in input (after current char)
	ch           byte // current char under examination
}

func NewLexer(input string) *Lexer {
	l := &Lexer{input: input}
	l.readChar()

	return l
}

func (l *Lexer) NextToken() token.Token {
	var t token.Token

	l.skipWhitespace()

	switch l.ch {
	case '=':
		if l.peekChar() == '=' {
			l.readChar()

			t = token.NewToken(token.Equals, string(token.Equals))
		} else {
			t = token.NewToken(token.Assign, string(token.Assign))
		}
	case ';':
		t = token.NewToken(token.Semicolon, string(token.Semicolon))
	case '(':
		t = token.NewToken(token.LeftParen, string(token.LeftParen))
	case ')':
		t = token.NewToken(token.RightParen, string(token.RightParen))
	case ',':
		t = token.NewToken(token.Comma, string(token.Comma))
	case '+':
		t = token.NewToken(token.Plus, string(token.Plus))
	case '-':
		t = token.NewToken(token.Minus, string(token.Minus))
	case '!':
		if l.peekChar() == '=' {
			l.readChar()

			t = token.NewToken(token.NotEquals, string(token.NotEquals))
		} else {
			t = token.NewToken(token.Bang, string(token.Bang))
		}
	case '*':
		t = token.NewToken(token.Asterisk, string(token.Asterisk))
	case '/':
		t = token.NewToken(token.Slash, string(token.Slash))
	case '<':
		t = token.NewToken(token.LessThan, string(token.LessThan))
	case '>':
		t = token.NewToken(token.GreaterThan, string(token.GreaterThan))
	case '{':
		t = token.NewToken(token.LeftBrace, string(token.LeftBrace))
	case '}':
		t = token.NewToken(token.RightBrace, string(token.RightBrace))
	case 0:
		t = token.NewToken(token.Eof, "")
	default:
		if isLetter(l.ch) {
			return token.NewTokenForIdentifier(l.readIdentifier())
		} else if isDigit(l.ch) {
			return token.NewToken(token.Integer, l.readNumber())
		} else {
			t = token.NewToken(token.Illegal, string(l.ch))
		}
	}

	l.readChar()

	return t
}

func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) {
		l.ch = 0
	} else {
		l.ch = l.input[l.readPosition]
	}

	l.position = l.readPosition
	l.readPosition++
}

func (l *Lexer) peekChar() byte {
	if l.readPosition >= len(l.input) {
		return 0
	}

	return l.input[l.readPosition]
}

func (l *Lexer) readIdentifier() string {
	pos := l.position

	for isLetter(l.ch) {
		l.readChar()
	}

	return l.input[pos:l.position]
}

func (l *Lexer) readNumber() string {
	pos := l.position
	for isDigit(l.ch) {
		l.readChar()
	}

	return l.input[pos:l.position]
}

func (l *Lexer) skipWhitespace() {
	for l.ch == ' ' || l.ch == '\t' || l.ch == '\n' || l.ch == '\r' {
		l.readChar()
	}
}

func isLetter(ch byte) bool {
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_'
}

func isDigit(ch byte) bool {
	return '0' <= ch && ch <= '9'
}
