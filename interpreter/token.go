package interpreter

import "fmt"

type tokenType byte

const (
	tokLeftParen tokenType = iota
	tokRightParen
	tokPlus
	tokMinus
	tokStar
	tokSlash
	tokNumber
	tokIdentifier
	tokEof
)

func (t tokenType) String() string {
	switch t {
	case tokLeftParen:
		return "LeftParen"
	case tokRightParen:
		return "RightParen"
	case tokPlus:
		return "Plus"
	case tokMinus:
		return "Minus"
	case tokStar:
		return "Star"
	case tokSlash:
		return "Slash"
	case tokNumber:
		return "Number"
	case tokIdentifier:
		return "Identifier"
	case tokEof:
		return "EOF"
	default:
		return "undefined"
	}
}

type token struct {
	tokType tokenType
	lexeme  string
	line    int
	col     int
	literal any
}

func (t *token) String() string {
	return fmt.Sprintf("<%s [%d:%d]: `%s`>", t.tokType, t.line, t.col, t.lexeme)
}
