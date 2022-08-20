package ast

import (
	"fmt"

	e "github.com/SyMind/helang-go/env"
)

/*
 print: PRINT expr SEMICOLON

 sprint: SPRINT expr SEMICOLON

 var_def: U8 IDENT ASSIGN expr SEMICOLON

 var_declare: U8 IDENT SEMICOLON

 var_assign: IDENT ASSIGN expr SEMICOLON

 var_increment: IDENT INCREMENT SEMICOLON

 expr_statement: expr SEMICOLON

 test_5g: TEST_5G SEMICOLON

 cyberspaces: CYBERSPACES SEMICOLON

 semicolon: SEMICOLON

  expr
          : empty_u8 expr'
          | or_u8 expr'
          | var expr'


empty_u8: LS NUMBER RS


or_u8
            : NUMBER
            | NUMBER OR or_u8_expr

        :return: or initializer for u8.


		var: IDENT

expr'
            : LS expr RS ASSIGN expr expr'
            | LS expr RS expr'
            | SUB expr expr'
            | ADD expr expr'
            | empty

        :param prev:
        :return:
*/

type CyberNamePanic struct{}

type AST interface {
	Eval(env *e.Env) e.U8
}

type Expr interface {
	AST
	isExpr()
}

/*
	print: PRINT expr SEMICOLON
*/
type Print struct {
	Value Expr
}

func (ast Print) Eval(env *e.Env) e.U8 {
	val := ast.Value.Eval(env)
	fmt.Println(val.String())
	return val
}

/*
	sprint: SPRINT expr SEMICOLON
*/
type Sprint struct {
	Value Expr
}

func (ast Sprint) Eval(env *e.Env) e.U8 {
	val := ast.Value.Eval(env)
	str := ""
	for _, char := range val.Values {
		str += string(rune(char))
	}
	fmt.Println(val)
	return val
}

/*
	var_def: U8 IDENT ASSIGN expr SEMICOLON
*/
type VarDef struct {
	Identifier string
	valueOrNil Expr
}

func (ast VarDef) Eval(env *e.Env) e.U8 {
	val := ast.valueOrNil.Eval(env)
	env.Values[ast.Identifier] = val
	return e.NewU8()
}

/*
	expr u8
*/
type EU8 struct {
	Values []int
}

func (ast EU8) Eval(env *e.Env) e.U8 {
	u8 := e.NewU8()
	u8.SetValues(ast.Values)
	return u8
}

func (ast EU8) isExpr() {}
