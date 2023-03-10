package main

import (
	"errors"
	"fmt"
	"strconv"
	"unicode"
)

type Lexer struct {
	code    []rune
	line    int
	start   int
	current int
	tokens  []*Token
	err     error
}

func NewLexer(code []rune) *Lexer {
	return &Lexer{code: code}
}

func (l *Lexer) Next() bool {
	switch r := l.advance(); r {
	case '(':
		l.token(TokLeftParen, nil)
	case ')':
		l.token(TokRightParen, nil)
	case '+':
		l.token(TokPlus, nil)
	case '-':
		l.token(TokMinus, nil)
	case '*':
		l.token(TokStar, nil)
	case '/':
		l.token(TokSlash, nil)
	case ' ', '\t', '\r':
	case '\n':
		l.line++
	default:
		if unicode.IsDigit(r) {
			l.number()
		} else if unicode.IsLetter(r) {
			l.identifier()
		}
	}
	l.start = l.current
	return !l.eof() && l.err == nil
}

func (l *Lexer) eof() bool {
	return l.current >= len(l.code)
}

func (l *Lexer) peek() rune {
	return l.code[l.current]
}

func (l *Lexer) advance() rune {
	defer func() { l.current++ }()
	return l.code[l.current]
}

func (l *Lexer) text() string {
	return string(l.code[l.start:l.current])
}

func (l *Lexer) token(tokType TokenType, literal any) {
	l.tokens = append(l.tokens, &Token{
		Type:    tokType,
		Lexeme:  l.text(),
		Line:    l.line,
		Col:     l.current,
		Literal: literal,
	})
}

func (l *Lexer) number() {
	for !l.eof() && unicode.IsDigit(l.peek()) {
		l.current++
	}

	if l.peek() == '.' {
		l.current++
		if !unicode.IsDigit(l.peek()) {
			l.err = errors.New("expected a digit after .")
			return
		}
		for !l.eof() && unicode.IsDigit(l.peek()) {
			l.current++
		}

		lexeme := l.text()
		num, err := strconv.ParseFloat(lexeme, 64)
		if err != nil {
			l.err = fmt.Errorf("failed to convert %s to number", lexeme)
			return
		}
		l.token(TokNumber, num)
		return
	}

	lexeme := l.text()
	num, err := strconv.Atoi(lexeme)
	if err != nil {
		l.err = fmt.Errorf("failed to convert %s to a number", lexeme)
		return
	}
	l.token(TokNumber, num)
}

func (l *Lexer) identifier() {
	for !l.eof() && (unicode.IsLetter(l.peek()) || l.peek() == '-') {
		l.current++
	}
	l.token(TokIdentifier, nil)
}
