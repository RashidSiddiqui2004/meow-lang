package lexer

import (
	"unicode"
)

type Token struct {
	Type    TokenType
	Literal string
}

// Lexer holds the state of the lexer.
type Lexer struct {
	input   string
	pos     int  // current position in input
	readPos int  // current reading position in input
	ch      byte // current character under examination
}

// NewLexer initializes a new instance of Lexer with the input string.
func NewLexer(input string) *Lexer {
	l := &Lexer{input: input}
	l.readChar()
	return l
}

func (l *Lexer) readChar() {
	if l.readPos >= len(l.input) {
		l.ch = 0 // ASCII code for NUL, signifies end of file
	} else {
		l.ch = l.input[l.readPos]
	}
	l.pos = l.readPos
	l.readPos++
}

// peekChar returns the next character without advancing the lexer.
func (l *Lexer) peekChar() byte {
	if l.readPos >= len(l.input) {
		return 0
	}
	return l.input[l.readPos]
}

// NextToken lexes and returns the next token.
func (l *Lexer) NextToken() Token {
	var tok Token

	l.skipWhitespace()

	switch l.ch {
	case '=':
		tok = newToken(ASSIGN, string(l.ch))
	case ';':
		tok = newToken(SEMICOLON, string(l.ch))
	case '(':
		tok = newToken(LPAREN, string(l.ch))
	case ')':
		tok = newToken(RPAREN, string(l.ch))
	case '"':
		tok.Type = STRING
		tok.Literal = l.readString()
		return tok
	case '+':
		tok = newToken(OPERATOR, string(l.ch))
	case '-':
		tok = newToken(OPERATOR, string(l.ch))
	case '*':
		tok = newToken(OPERATOR, string(l.ch))
	case '/':
		tok = newToken(OPERATOR, string(l.ch))
	case '&':
		tok = newToken(OPERATOR, string(l.ch))
	case '|':
		tok = newToken(OPERATOR, string(l.ch))
	case '^':
		tok = newToken(OPERATOR, string(l.ch))
	case '<':
		// Check for the "<<" operator
		if l.peekChar() == '<' {
			ch := l.ch
			l.readChar()
			tok = newToken(OPERATOR, string(ch)+string(l.ch))
		} else {
			tok = newToken(OPERATOR, string(l.ch))
		}
	case '>':
		// Check for the ">>" operator
		if l.peekChar() == '>' {
			ch := l.ch
			l.readChar()
			tok = newToken(OPERATOR, string(ch)+string(l.ch))
		} else {
			tok = newToken(OPERATOR, string(l.ch))
		}
	case 0:
		tok.Literal = ""
		tok.Type = EOF
	default:
		if isLetter(l.ch) {
			literal := l.readIdentifier()
			tok.Type = lookupIdent(literal)
			tok.Literal = literal
			return tok
		} else if isDigit(l.ch) {
			tok.Type = NUMBER
			tok.Literal = l.readNumber()
			return tok
		} else {
			// For unknown characters, create an EOF token to avoid getting stuck.
			tok = newToken(EOF, "")
		}
	}

	l.readChar()
	return tok
}

// newToken creates a new token of a given type and literal.
func newToken(tokenType TokenType, ch string) Token {
	return Token{Type: tokenType, Literal: ch}
}

func (l *Lexer) skipWhitespace() {
	for l.ch == ' ' || l.ch == '\t' || l.ch == '\n' || l.ch == '\r' || l.ch == '~' {
		if l.ch == '~' {
			l.skipComment()
		} else {
			l.readChar()
		}
	}
}

// skipComment advances the lexer until the end of the current line.
func (l *Lexer) skipComment() {
	for l.ch != '\n' && l.ch != 0 {
		l.readChar()
	}
}

// isLetter returns true if the character is a letter or underscore.
func isLetter(ch byte) bool {
	return unicode.IsLetter(rune(ch)) || ch == '_'
}

// readIdentifier reads an identifier and advances the lexer's position.
func (l *Lexer) readIdentifier() string {
	position := l.pos
	for isLetter(l.ch) {
		l.readChar()
	}
	return l.input[position:l.pos]
}

// isDigit returns true if the character is a digit.
func isDigit(ch byte) bool {
	return '0' <= ch && ch <= '9'
}

// readNumber reads a number and advances the lexer's position.
func (l *Lexer) readNumber() string {
	position := l.pos
	for isDigit(l.ch) {
		l.readChar()
	}
	return l.input[position:l.pos]
}

// readString reads a string literal, assuming the current character is a double quote.
func (l *Lexer) readString() string {
	position := l.pos + 1
	for {
		l.readChar()
		if l.ch == '"' || l.ch == 0 {
			break
		}
	}
	str := l.input[position:l.pos]
	l.readChar()
	return str
}

// lookupIdent checks if an identifier is a reserved keyword and returns the appropriate TokenType.
func lookupIdent(ident string) TokenType {
	switch ident {
	case "meowstart":
		return MEOWSTART
	case "meowend":
		return MEOWEND
	case "meow":
		return MEOW
	case "purr":
		return PURR
	default:
		return IDENT
	}
}
