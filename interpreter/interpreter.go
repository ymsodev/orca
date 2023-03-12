package interpreter

type Interpreter struct {
	code  string
	lines []string
}

func Interpret(code string) {
	lexer := newLexer(code)
	for lexer.Next() {
	}
	if lexer.err != nil {

	}
}

func printError(err *Error) {
	fmt.Printf("error: %v\n", err.Msg)
	if err.Line > 0 {
		fmt.Println(e.prettyLine(err.Line - 1))
	}
	pretty := e.prettyLine(err.Line)
	fmt.Println(pretty)
	pad := len(pretty) - len(e.lines[err.Line])
	fmt.Println(strings.Repeat(" ", pad) + strings.Repeat("^", err.End-err.Start))
}

func prettyLine(i int) string {
	return fmt.Sprintf("%s | %s", e.lineNumStr(i), e.lines[i])
}

func lineNumStr(line int) string {
	n := float64(len(e.lines))
	spaces := int(math.Floor(math.Log10(n))) + 1
	return fmt.Sprintf("% "+strconv.Itoa(spaces)+"d", line)
}
