package interpreter

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

type Error struct {
	Msg   string // error message
	Line  int    // line number in the code
	Start int    // highlight start
	End   int    // highlight end
}

type ErrorHandler struct {
	lines []string
}

func NewErrorPrinter(code string) *ErrorPrinter {
	return &ErrorPrinter{strings.Split(code, "\n")}
}

func (e *ErrorPrinter) PrintError(err *Error) {
	fmt.Printf("error: %v\n", err.Msg)
	if err.Line > 0 {
		fmt.Println(e.prettyLine(err.Line - 1))
	}
	pretty := e.prettyLine(err.Line)
	fmt.Println(pretty)
	pad := len(pretty) - len(e.lines[err.Line])
	fmt.Println(strings.Repeat(" ", pad) + strings.Repeat("^", err.End-err.Start))
}

func (e *ErrorPrinter) prettyLine(i int) string {
	return fmt.Sprintf("%s | %s", e.lineNumStr(i), e.lines[i])
}

func (e *ErrorPrinter) lineNumStr(line int) string {
	n := float64(len(e.lines))
	spaces := int(math.Floor(math.Log10(n))) + 1
	return fmt.Sprintf("% "+strconv.Itoa(spaces)+"d", line)
}
