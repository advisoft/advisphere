package main

import (
	"advisphere/internal/lox"
	"fmt"
	"os"
)

func main() {
	if len(os.Args) > 2 {
		fmt.Println("Usage: prog [script]")
		os.Exit(64)
	}

	lexer := lox.CreateLexer("(1+  2) \"i am cool\" * 3 / 5 // comment")
	tokens := lexer.ScanTokens()
	for _, token := range tokens {
		fmt.Println(token)
	}

	expr := lox.BinaryExpr{
		Left: lox.UnaryExpr{
			Operator: lox.Token{Type: lox.TokenType(lox.MINUS), Lexeme: "-", Literal: nil, Line: 1},
			Right:    lox.LiteralExpr{Value: 123},
		},
		Operator: lox.Token{Type: lox.TokenType(lox.STAR), Lexeme: "*", Literal: nil, Line: 1},
		Right:    lox.GroupingExpr{Expression: lox.LiteralExpr{Value: 45.67}},
	}

	printer := lox.AstPrinter{}
	fmt.Println(printer.Print(expr))
}
