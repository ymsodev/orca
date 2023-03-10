package main

import "fmt"

func main() {
	code := []rune("()+-*/12.345hello-world")
	lexer := NewLexer(code)
	for {
		if !lexer.Next() {
			break
		}
	}
	if lexer.err != nil {
		panic(lexer.err)
	}
	for _, t := range lexer.tokens {
		fmt.Println(t)
	}
}
