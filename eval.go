package goexpr

import (
	"fmt"
	"go/ast"
	"go/token"
	"reflect"
	"strconv"
)

func Evaluate(parsed *ParsedExpression, scope map[string]interface{}) (float64, error) {

	result, err := evaluate(parsed.Ast, scope)

	if err != nil {
		return 0, err
	}

	fmt.Println("result", result, err)

	return result.(float64), nil
}

func evaluate(node ast.Node, scope map[string]interface{}) (value interface{}, err error) {

	switch node.(type) {

	case *ast.Ident:
		ident := node.(*ast.Ident)
		fmt.Println("IDENT: ", node.(*ast.Ident).Name)

		v, found := scope[ident.Name]
		fmt.Printf("IDENT: %v %v\n", v, found)

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

		lFloat, ok := lValue.(float64)
		//fmt.Println(lValue, reflect.TypeOf(lValue), lFloat, ok)

		if !ok {
			err = fmt.Errorf("could not convert %v (left) to float64", lValue)
			break
		}

		rFloat, ok := rValue.(float64)

		if !ok {
			err = fmt.Errorf("could not convert %v to float64", rValue)
			break
		}

		switch exprNode.Op {
		case token.ADD:
			value = lFloat + rFloat
		case token.SUB:
			value = lFloat - rFloat
		case token.MUL:
			value = lFloat * rFloat
		case token.QUO:
			value = lFloat / rFloat
		}

	case *ast.SelectorExpr:
		sel := node.(*ast.SelectorExpr)
		fmt.Println("SELECTOR", sel.X, sel.Sel)

		ident, e := evaluate(sel.X, scope)

		if e != nil {
			err = e
			break
		}

		value = ident
		fmt.Println("SELECTOR: ident %v", ident)

		r := reflect.ValueOf(ident)
		v := r.FieldByName(sel.Sel.Name)
		f := v.Float()
		fmt.Println(sel.Sel.Name, r, v, f)

		//v := reflect.FieldByName(scope[sel.X])
		value = f
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

	fmt.Println("evaluated:", value, err)

	return value, err
}
