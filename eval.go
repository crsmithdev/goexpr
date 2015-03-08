package goexpr

import (
	"fmt"
	"go/ast"
	"go/token"
	"strconv"
)

func Evaluate(parsed *ParsedExpression, scope map[string]float64) (float64, error) {
	return evaluate(parsed.Ast, scope)
}

func evaluate(node ast.Node, scope map[string]float64) (value float64, err error) {

	switch node.(type) {

	case *ast.Ident:
		ident := node.(*ast.Ident)
		fmt.Println("IDENT", node.(*ast.Ident).Name)

		v, found := scope[ident.Name]

		if !found {
			err = fmt.Errorf("no value for %s", ident.Name)
			break
		} else {
			value = v
		}

	case *ast.BinaryExpr:
		exprNode := node.(*ast.BinaryExpr)

		lValue, e := evaluate(exprNode.X, scope)

		if e != nil {
			err = e
			break
		}

		rValue, e := evaluate(exprNode.Y, scope)

		if e != nil {
			err = e
			break
		}

		switch exprNode.Op {
		case token.ADD:
			value = lValue + rValue
		case token.SUB:
			value = lValue - rValue
		case token.MUL:
			value = lValue * rValue
		case token.QUO:
			value = lValue / rValue
		}

		evaluate(exprNode.Y, scope)

		/*
			case *ast.SelectorExpr:
				sel := node.(*ast.SelectorExpr)
				evaluate(sel.X, scope)
				fmt.Println("SELECTOR", sel.X, sel.Sel)
		*/
	case *ast.ParenExpr:
		value, err = evaluate(node.(*ast.ParenExpr).X, scope)
	case *ast.BasicLit:
		lit := node.(*ast.BasicLit)
		float, e := strconv.ParseFloat(lit.Value, 64)

		if e != nil {
			err = e
			break
		}

		value = float
		//value = node.(*ast.BasicLit).Value
	default:
		err = fmt.Errorf("unsupported node %+v (%d - %d)", node, node.Pos(), node.End())
	}

	return value, err
}
