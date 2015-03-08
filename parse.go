package goexpr

import (
	"fmt"
	"go/ast"
	"go/parser"
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
		exprNode := node.(*ast.BinaryExpr)
		lVars, e := parse(exprNode.X)

		if e != nil {
			err = e
			break
		}

		rVars, err := parse(exprNode.Y)

		if err != nil {
			err = e
			break
		}

		vars = append(lVars, rVars...)

	case *ast.SelectorExpr:
		vars, err = parse(node.(*ast.SelectorExpr).X)

	case *ast.ParenExpr:
		vars, err = parse(node.(*ast.ParenExpr).X)

	case *ast.BasicLit:
		break

	default:
		err = fmt.Errorf("unsupported node %+v (%d - %d)", node, node.Pos(), node.End())
	}

	return vars, err
}
