package interpreter

func Interpret(code string) {
	lexer := newLexer(code)
	lexer.tokenize()

	parser := newParser(tokens)
}
