package main

import (
	l "github.com/SyMind/helang-go/lexer"
	p "github.com/SyMind/helang-go/parser"
)

func main() {
	lexer := l.NewLexer("print [1 | 2 | 3];")
	parser := p.NewParser(lexer)
	parser.Parse()
}
