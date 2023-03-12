package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/ymsodev/orca/interpreter"
)

func main() {
	code := `
(* (+ 1 2 3 4 5)
   hello-world)
`
	ep := NewErrorPrinter(code)
	lexer := NewLexer(code)
	for lexer.Next() {
		// hmmmmmmmm, do we want this?
	}
	if lexer.err != nil {
		ep.PrintError(lexer.err)
	}

	for _, t := range lexer.tokens {
		fmt.Println(t)
	}

	ep.PrintError(&Error{
		Msg:   "this is a test error!",
		Line:  1,
		Start: 5,
		End:   12,
	})

	repl()
}

func repl() {
	scanner := bufio.NewScanner(os.Stdin)
	prompt := func() bool {
		fmt.Print("> ")
		return scanner.Scan()
	}
	for prompt() {
		code := scanner.Text()

		ep := NewErrorPrinter(code)
		lexer := NewLexer(code)
		for lexer.Next() {
		}
		if lexer.err != nil {
			ep.PrintError(lexer.err)
		}
		for _, t := range lexer.tokens {
			fmt.Println(t)
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}
}
