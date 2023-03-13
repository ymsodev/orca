package interpreter

import (
	"fmt"
	"strconv"
	"unicode"
)

type lexer struct {
	code    string
	runes   []rune
	lineNum int
	start   int
	current int
	tokens  []*Token
	err     *Error
}

func newLexer(code string) *lexer {
	return &lexer{
		code:    code,
		runes:   []rune(code),
		lineNum: 0,
		start:   0,
		current: 0,
		tokens:  make([]*Token, 0),
	}
}

func (l *lexer) tokenize() {
	for !l.eof() && l.err != nil {
		l.next()
	}
	l.token(TokEof, nil)
}

func (l *lexer) next() {
	defer func() { l.start = l.current }()
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
		l.lineNum++
	default:
		if unicode.IsDigit(r) {
			l.number()
		} else if unicode.IsLetter(r) {
			l.identifier()
		} else {
			l.error(fmt.Sprintf("invalid character: `%s`", string(r)))
		}
	}
}

func (l *lexer) eof() bool {
	return l.current >= len(l.code)
}

func (l *lexer) peek() rune {
	if l.eof() {
		return '\000'
	}
	return l.runes[l.current]
}

func (l *lexer) advance() rune {
	defer func() { l.current++ }()
	return l.peek()
}

func (l *lexer) text() string {
	return l.code[l.start:l.current]
}

func (l *lexer) token(tokType TokenType, literal any) {
	l.tokens = append(l.tokens, &Token{
		Type:    tokType,
		Lexeme:  l.text(),
		Line:    l.lineNum,
		Col:     l.current,
		Literal: literal,
	})
}

func (l *lexer) number() error {
	for !l.eof() && unicode.IsDigit(l.peek()) {
		l.current++
	}
	if l.peek() == '.' {
		return l.float()
	}
	lexeme := l.text()
	num, err := strconv.Atoi(lexeme)
	if err != nil {
		return fmt.Errorf("invalid number: %s", lexeme)
	}
	l.token(TokNumber, num)
	return nil
}

func (l *lexer) float() error {
	l.current++ // .
	if !unicode.IsDigit(l.peek()) {
		return fmt.Errorf("invalid character: %s", string(l.peek()))
	}
	for !l.eof() && unicode.IsDigit(l.peek()) {
		l.current++
	}

	lexeme := l.text()
	num, err := strconv.ParseFloat(lexeme, 64)
	if err != nil {
		return fmt.Errorf("invalid number: %s", lexeme)
	}
	l.token(TokNumber, num)
	return nil
}

func (l *lexer) identifier() error {
	for !l.eof() && (unicode.IsLetter(l.peek()) || l.peek() == '-') {
		l.current++
	}
	l.token(TokIdentifier, nil)
	return nil
}

func (l *lexer) error(msg string) {
	l.err = &Error{
		Msg:   msg,
		Line:  l.lineNum,
		Start: l.start,
		End:   l.current,
	}
}
