package goexpr

import "go/ast"

type Expression struct {
	String string
}

type ParsedExpression struct {
	Expression
	Vars []string
	Ast  ast.Node
}
