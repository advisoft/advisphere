package main

import (
	"advisphere/internal/lox"
	"fmt"
	"os"

	"github.com/kr/pretty"
)

func main() {
	if len(os.Args) > 2 {
		fmt.Println("Usage: prog [script]")
		os.Exit(64)
	}

	src := "2 */ 2;"
	lexer := lox.CreateLexer(src)
	tokens := lexer.ScanTokens()

	// expr := lox.BinaryExpr{
	// 	Left: lox.UnaryExpr{
	// 		Operator: lox.Token{Type: lox.TokenType(lox.MINUS), Lexeme: "-", Literal: nil, Line: 1},
	// 		Right:    lox.LiteralExpr{Value: 123},
	// 	},
	// 	Operator: lox.Token{Type: lox.TokenType(lox.STAR), Lexeme: "*", Literal: nil, Line: 1},
	// 	Right:    lox.GroupingExpr{Expression: lox.LiteralExpr{Value: 45.67}},
	// }

	printer := lox.AstPrinter{}
	// fmt.Println(printer.Print(expr))

	parser := lox.CreateParser(tokens)
	expr := parser.Parse()
	fmt.Println(src)
	fmt.Println(printer.Print(expr))
	pretty.Println(expr)
}
