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

	return result.(float64), nil
}

func evaluate(node ast.Node, scope map[string]interface{}) (value interface{}, err error) {

	switch node.(type) {

	case *ast.Ident:
		ident := node.(*ast.Ident)

		v, found := scope[ident.Name]

		if !found {
			err = fmt.Errorf("no value for %s", ident.Name)
			break
		} else {

			switch v.(type) {
			case int:
				value = float64(v.(int))
			case float32, float64:
				value = v
			default:
				value = v

				t := reflect.TypeOf(v).Kind()
				if t == reflect.Struct {
					value = v
				} else {
					err = fmt.Errorf("unsupported type %v", t)
				}
			}
			fmt.Println(reflect.TypeOf(v), reflect.TypeOf(value))
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

		ident, e := evaluate(sel.X, scope)

		if e != nil {
			err = e
			break
		}

		value = ident

		r := reflect.ValueOf(ident)
		v := r.FieldByName(sel.Sel.Name)
		f := v.Float()

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
	default:
		err = fmt.Errorf("unsupported node %+v (%d - %d)", node, node.Pos(), node.End())
	}

	return value, err
}
