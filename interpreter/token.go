package interpreter

import "fmt"

type TokenType byte

const (
	TokLeftParen TokenType = iota
	TokRightParen
	TokPlus
	TokMinus
	TokStar
	TokSlash
	TokNumber
	TokIdentifier
	TokEof
)

func (t TokenType) String() string {
	switch t {
	case TokLeftParen:
		return "LeftParen"
	case TokRightParen:
		return "RightParen"
	case TokPlus:
		return "Plus"
	case TokMinus:
		return "Minus"
	case TokStar:
		return "Star"
	case TokSlash:
		return "Slash"
	case TokNumber:
		return "Number"
	case TokIdentifier:
		return "Identifier"
	case TokEof:
		return "EOF"
	default:
		return "undefined"
	}
}

type Token struct {
	Type    TokenType
	Lexeme  string
	Line    int
	Col     int
	Literal any
}

func (t *Token) String() string {
	return fmt.Sprintf("<%s [%d:%d]: `%s`>", t.Type, t.Line, t.Col, t.Lexeme)
}
