package expr

import (
	"expr/expr/node"
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
)

type Engine struct {
	Expr ast.Expr
}

func NewEngine(expr string) (*Engine, error) {
	parseExpr, err := parser.ParseExpr(expr)
	if err != nil {
		return nil, err
	}

	return &Engine{Expr: parseExpr}, nil
}

func (e *Engine) Run(m map[string]interface{}) node.ValueNode {
	pm := parseControlMap(m)
	if len(pm) == 0 {
		fmt.Println("[Engine.Run]parseControlMap empty")
		return node.NewBadNode("map empty")
	}
	n := eval(pm, e.Expr)
	bValue, _ := n.(node.BoolNode)
	fmt.Println(fmt.Sprintf("%#v", bValue))
	return n
}

func parseControlMap(controlMap map[string]interface{}) map[string]node.ValueNode {
	nodeMap := make(map[string]node.ValueNode, len(controlMap))
	for key, value := range controlMap {
		switch value.(type) {
		case int:
			n := node.NewFloat64(float64(value.(int)), key)
			nodeMap[key] = n
		case int64:
			n := node.NewFloat64(float64(value.(int64)), key)
			nodeMap[key] = n
		case float64:
			// value from json will be always float64
			n := node.NewFloat64(value.(float64), key)
			nodeMap[key] = n
		}
	}
	return nodeMap
}

func eval(mem map[string]node.ValueNode, expr ast.Expr) (y node.ValueNode) {
	switch x := expr.(type) {
	case *ast.BasicLit: //表达式的值 类似于 a > 1 的1
		fmt.Println("qweqwe", x)
		return node.NewNodeByLit(x)
	case *ast.Ident: //表达式的右边的值 类似于a > 1的a
		return mem[x.Name]
	case *ast.ParenExpr: //表达式中的括号
		return eval(mem, x.X)
	case *ast.BinaryExpr:
		left := eval(mem, x.X)
		right := eval(mem, x.Y)
		if left == nil || right == nil {
			return node.NewBadNode(fmt.Sprintf("%+v, %+v is nil", right, left))
		}
		fmt.Println(fmt.Sprintf("left:%#v,right:%#v", left, right))
		op := x.Op
		switch left.GetType() {
		case node.TypeInt:
			return BinaryIntExpr{}.Invoke(left, right, op, expr)
		case node.TypeFloat:
			return BinaryFloatExpr{}.Invoke(left, right, op, expr)
		case node.TypeBool:
			return BinaryBoolExpr{}.Invoke(left, right, op, expr)
		default:
			return node.NewBadNode("a:" + right.GetTextValue() + "b:" + left.GetTextValue())
		}
	case *ast.UnaryExpr:
		n := eval(mem, x.X)
		if x.Op == token.SUB {
			fn, ok := n.(node.FloatNode)
			if ok {
				fn.Value = -fn.Value
				return fn
			}
		}
		return node.NewBadNode(fmt.Sprintf("ast.UnaryExpr :%x type is not suppoort,%T", x))
	default:
		return node.NewBadNode(fmt.Sprintf("%x type is not suppoort,%T", x, x))
	}
}
