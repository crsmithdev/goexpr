package goexpr

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"reflect"
)

// Parse parses a string into an Expression.
func Parse(str string) (*Expression, error) {

	tree, err := parser.ParseExpr(str)

	if err != nil {
		return nil, err
	}

	vars, err := extract(tree)

	if err != nil {
		return nil, err
	}

	return &Expression{
		String: str,
		Vars:   vars,
		Ast:    tree,
	}, nil
}

func extract(node ast.Node) (vars []string, err error) {

	switch node.(type) {

	case *ast.Ident:
		vars = []string{node.(*ast.Ident).Name}

	case *ast.BinaryExpr:
		vars, err = extractBinary(node.(*ast.BinaryExpr))

	case *ast.ParenExpr:
		vars, err = extract(node.(*ast.ParenExpr).X)

	case *ast.BasicLit:
		break

	default:
		err = fmt.Errorf("unsupported node %+v (type %+v)", node, reflect.TypeOf(node))
	}

	return vars, err
}

func extractBinary(node *ast.BinaryExpr) ([]string, error) {

	var vars []string

	switch node.Op {
	case token.ADD, token.SUB, token.MUL, token.QUO:
		break
	default:
		return vars, fmt.Errorf("unsupported binary operation: %s", node.Op)
	}

	lVars, err := extract(node.X)

	if err != nil {
		return vars, err
	}

	rVars, err := extract(node.Y)

	if err != nil {
		return vars, err
	}

	vars = append(lVars, rVars...)

	return vars, err
}
