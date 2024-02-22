package main

import (
	"advisphere/internal/lexer"
	"fmt"
	"log"
)

func main() {
	fmt.Println("Hello")

	result := lexer.Lex("1+2")

	log.Println(result)

}
