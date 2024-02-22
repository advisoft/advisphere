package main

import (
	"advisphere/internal/frontend"
	"fmt"
	"os"
)

func main() {
	if len(os.Args) > 2 {
		fmt.Println("Usage: prog [script]")
		os.Exit(64)
	}

	lexer := frontend.CreateLexer("(1+  2) \"i am cool\" * 3 / 5 // comment")
	tokens := lexer.ScanTokens()
	for _, token := range tokens {
		fmt.Println(token)
	}
}
