package parser

import (
	"fmt"

	a "github.com/SyMind/helang-go/ast"
	e "github.com/SyMind/helang-go/env"
	l "github.com/SyMind/helang-go/lexer"
)

type BadStatementPanic struct{}

type Parser struct {
	lexer l.Lexer
}

func NewParser(lexer l.Lexer) Parser {
	return Parser{
		lexer: lexer,
	}
}

func (p *Parser) Parse() {
	/*
		root
			: print
			| sprint
			| u8_set
			| var_def
			| var_declare
			| var_assign
			| var_increment
			| expr
			| test_5g
			| semicolon
			| cyberspaces
	*/
	asts := make([]a.AST, 0, 1)
	for p.lexer.Token != l.TEndOfFile {
		ast := p.parseStatement()
		p.lexer.Next()
		asts = append(asts, ast)
	}

	for _, ast := range asts {
		fmt.Printf("%s", ast)
		env := e.NewEnv()
		ast.Eval(&env)
	}
}

func (p *Parser) except(token l.Token) {
	p.lexer.Next()
	if p.lexer.Token != token {
		panic(BadStatementPanic{})
	}
}

func (p *Parser) parseStatement() a.AST {
	switch p.lexer.Token {
	case l.TPrint:
		return p.parsePrint()
	case l.TSprint:
	case l.TOpenBracket:
		return p.parseEU8()
	default:
		break
	}
	panic(BadStatementPanic{})
}

func (p *Parser) parseExpr() a.Expr {
	switch p.lexer.Token {
	case l.TSprint:
	case l.TOpenBracket:
		return p.parseEU8()
	default:
		break
	}
	panic(BadStatementPanic{})
}

func (p *Parser) parsePrint() a.AST {
	p.lexer.Next()
	ast := a.Print{
		Value: p.parseExpr(),
	}
	p.except(l.TColon)
	return ast
}

func (p *Parser) parseEU8() a.EU8 {
	p.lexer.Next()
	values := make([]int, 0, 1)
	for p.lexer.Token != l.TCloseBracket {
		switch p.lexer.Token {
		case l.TNumber:
			values = append(values, p.lexer.Number)
			p.lexer.Next()
			continue
		case l.TBar:
			p.lexer.Next()
			continue
		default:
			panic(BadStatementPanic{})
		}
	}
	return a.EU8{
		Values: values,
	}
}
