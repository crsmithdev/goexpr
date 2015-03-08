package goexpr

import "go/ast"

// Expression is a mathematic expression that can be evaluated, given a scope.
type Expression struct {
	String string   // The expression string.
	Vars   []string // The variable names in the expression.
	Ast    ast.Node // the root of the expression AST.
}
