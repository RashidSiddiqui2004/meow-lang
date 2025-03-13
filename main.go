package main

import (
	"fmt"
	"meowlang/lexer"
	"meowlang/parser"
	"os"
)

func ReadFile(fileName string) string {
	data, err := os.ReadFile(fileName)
	if err != nil {
		fmt.Printf("Error reading file: %v\n", err)
		os.Exit(1)
	}
	return string(data)
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: meowlang <source_file.meow>")
		os.Exit(1)
	}

	fileName := os.Args[1]
	sourceCode := ReadFile(fileName)

	lex := lexer.NewLexer(sourceCode)
	var tokens []lexer.Token

	for {
		tok := lex.NextToken()
		tokens = append(tokens, tok)
		if tok.Type == lexer.EOF {
			break
		}
	}

	fmt.Println("Tokens:", tokens)

	p := parser.NewParser(tokens)
	p.Parse()
}
