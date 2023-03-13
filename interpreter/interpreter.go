package interpreter

func Interpret(code string) {
	tokens, err := newLexer(code).tokenize()
	if err != nil {

	}
	parser := newParser(tokens)
}
