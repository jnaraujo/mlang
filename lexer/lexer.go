package lexer

import "mlang/token"

type Lexer struct {
	input        string
	position     int  // current char position
	readPosition int  // after current char position
	char         byte // current char
}

func New(input string) *Lexer {
	lexer := &Lexer{
		input: input,
	}
	lexer.readChar()
	return lexer
}

func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) {
		l.char = 0
	} else {
		l.char = l.input[l.readPosition]
	}

	l.position = l.readPosition
	l.readPosition++
}

func (l *Lexer) NextToken() token.Token {
	var tkn token.Token

	l.skipWhitespace()

	switch l.char {
	case '=':
		tkn = newToken(token.ASSIGN, l.char)
	case ';':
		tkn = newToken(token.SEMICOLON, l.char)
	case ',':
		tkn = newToken(token.COMMA, l.char)
	case '(':
		tkn = newToken(token.LPAREN, l.char)
	case ')':
		tkn = newToken(token.RPAREN, l.char)
	case '{':
		tkn = newToken(token.LBRACE, l.char)
	case '}':
		tkn = newToken(token.RBRACE, l.char)
	case '+':
		tkn = newToken(token.PLUS, l.char)
	case 0:
		tkn.Type = token.EOF
		tkn.Literal = ""
	default:
		if isLetter(l.char) {
			tkn.Literal = l.readIdentifier()
			tkn.Type = token.LookupIdent(tkn.Literal)
			return tkn
		}
		if isDigit(l.char) {
			tkn.Type = token.INT
			tkn.Literal = l.readNumber()
			return tkn
		}
		tkn = newToken(token.ILLEGAL, l.char)
	}

	l.readChar()
	return tkn
}

func (l *Lexer) readIdentifier() string {
	position := l.position
	for isLetter(l.char) {
		l.readChar()
	}
	return l.input[position:l.position]
}

func (l *Lexer) readNumber() string {
	position := l.position
	for isDigit(l.char) {
		l.readChar()
	}
	return l.input[position:l.position]
}

func (l *Lexer) skipWhitespace() {
	for l.char == ' ' || l.char == '\t' || l.char == '\n' || l.char == '\r' {
		l.readChar()
	}
}

func isLetter(char byte) bool {
	return (char >= 'a' && char <= 'z') || (char >= 'A' && char <= 'Z') || char == '_'
}

func isDigit(char byte) bool {
	return char >= '0' && char <= '9'
}

func newToken(tokenType token.TokenType, char byte) token.Token {
	return token.Token{
		Type:    tokenType,
		Literal: string(char),
	}
}
