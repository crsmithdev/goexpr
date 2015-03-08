package goexpr

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
)

func Parse(expr string) (*ParsedExpression, error) {

	tree, err := parser.ParseExpr(expr)

	if err != nil {
		return nil, err
	}

	vars, err := parse(tree)

	if err != nil {
		return nil, err
	}

	return &ParsedExpression{
		Vars: vars,
		Ast:  tree,
	}, nil
}

func parse(node ast.Node) (vars []string, err error) {

	switch node.(type) {

	case *ast.Ident:
		ident := node.(*ast.Ident)
		vars = []string{ident.Name}

	case *ast.BinaryExpr:
		vars, err = parseBinary(node.(*ast.BinaryExpr))

	case *ast.ParenExpr:
		vars, err = parse(node.(*ast.ParenExpr).X)

	case *ast.BasicLit:
		break

	default:
		err = fmt.Errorf("unsupported node %+v (%d - %d)", node, node.Pos(), node.End())
	}

	return vars, err
}

func parseBinary(node *ast.BinaryExpr) ([]string, error) {

	var vars []string

	switch node.Op {
	case token.ADD, token.SUB, token.MUL, token.QUO:
		break
	default:
		return vars, fmt.Errorf("unsupported binary operation: %s", node.Op)
	}

	lVars, err := parse(node.X)

	if err != nil {
		return vars, err
	}

	rVars, err := parse(node.Y)

	if err != nil {
		return vars, err
	}

	vars = append(lVars, rVars...)

	return vars, err
}
