package expr

import (
	"bytes"
	"expr/expr/node"
	"fmt"
	"github.com/shopspring/decimal"
	"go/ast"
	"go/printer"
	"go/token"
)

type BinaryIntExpr struct{}

type BinaryFloatExpr struct{}

type BinaryBoolExpr struct{}

func (b BinaryIntExpr) Invoke(x, y node.ValueNode, op token.Token, expr ast.Expr) node.ValueNode {
	xs, xok := x.(node.IntNode)
	ys, yok := y.(node.IntNode)

	if !xok || !yok {
		return node.NewBadNode(x.GetTextValue() + y.GetTextValue())
	}
	var filed = x.GetField()
	var n node.ValueNode
	switch op {
	case token.EQL: // ==
		n = node.NewBoolNode(xs.Value == ys.Value, filed)
	case token.LSS: // <
		n = node.NewBoolNode(xs.Value < ys.Value, filed)
	case token.GTR: // >
		n = node.NewBoolNode(xs.Value > ys.Value, filed)
	case token.GEQ: // >=
		n = node.NewBoolNode(xs.Value >= ys.Value, filed)
	case token.LEQ: // <=
		n = node.NewBoolNode(xs.Value <= ys.Value, filed)
	default:
		return node.NewBadNode(fmt.Sprintf("unsupported binary operator: %s", op.String()))
	}
	bn := n.(node.BoolNode)
	if bn.True {
		bn.Expr = AstExprToString(expr)
		bn.FieldMap[filed] = bn.Expr
	}
	return bn
}

func (b BinaryFloatExpr) Invoke(x, y node.ValueNode, op token.Token, expr ast.Expr) node.ValueNode {
	xs, xok := x.(node.FloatNode)
	ys, yok := y.(node.FloatNode)
	if !xok || !yok {
		return node.NewBadNode(x.GetTextValue() + y.GetTextValue())
	}

	var filed = x.GetField()
	var n node.ValueNode
	switch op {
	case token.EQL: // ==
		n = node.NewBoolNode(decimal.NewFromFloat(xs.Value).Equal(decimal.NewFromFloat(ys.Value)), filed)
	case token.LSS: // <
		n = node.NewBoolNode(decimal.NewFromFloat(xs.Value).LessThan(decimal.NewFromFloat(ys.Value)), filed)
	case token.GTR: // >
		n = node.NewBoolNode(decimal.NewFromFloat(xs.Value).GreaterThan(decimal.NewFromFloat(ys.Value)), filed)
	case token.GEQ: // >=
		n = node.NewBoolNode(decimal.NewFromFloat(xs.Value).GreaterThanOrEqual(decimal.NewFromFloat(ys.Value)), filed)
	case token.LEQ: // <=
		n = node.NewBoolNode(decimal.NewFromFloat(xs.Value).LessThanOrEqual(decimal.NewFromFloat(ys.Value)), filed)
	default:
		return node.NewBadNode(fmt.Sprintf("unsupported binary operator: %s", op.String()))
	}

	bn := n.(node.BoolNode)
	if bn.True {
		bn.Expr = AstExprToString(expr)
		bn.FieldMap[filed] = bn.Expr
	}
	return bn
}

func (b BinaryBoolExpr) Invoke(x, y node.ValueNode, op token.Token, expr ast.Expr) node.ValueNode {
	xb, xok := x.(node.BoolNode)
	yb, yok := y.(node.BoolNode)

	if !xok || !yok {
		return node.NewBadNode(x.GetTextValue() + y.GetTextValue())
	}
	var n node.ValueNode
	var condition string
	switch op {
	case token.LAND:
		n = node.NewBoolNode(xb.True && yb.True, "")
		condition = " && "
	case token.LOR:
		n = node.NewBoolNode(xb.True || yb.True, "")
		condition = " || "
	default:
		return node.NewBadNode(fmt.Sprintf("unsupported binary operator: %s", op.String()))
	}
	//主要是为了备忘
	//1. bool值的比较主要是两个表达式比较后的结果属于是汇总
	//例如 a > 1 || b < 2 只有在a>1和b<2都比较完之后才可能出现bool和bool值的比较
	//我这里没有考虑 == true 和 非 抑或之类的情况
	//2. 所以在bool的节点来汇总字段和表达式

	bn := n.(node.BoolNode)
	//如果当前结果为false 说明当前表达式不成立 直接返回
	if bn.True == false {
		return bn
	}
	// 命中的字段
	var fields []string
	// 表达式集合
	var exprCollect string

	//两个表达式相比较 如果左边的为true 说明左边的字段都是命中的，将左边的表达式拼接起来
	if xb.True {
		fields = append(fields, xb.Fields...)
		exprCollect = fmt.Sprintf("%s", xb.Expr)
		if len(xb.FieldMap) > 0 {
			bn.FieldMap = mergeFieldMap(xb.FieldMap, bn.FieldMap)
		}
	}
	//两个表达式相比较 如果右边的为true 说明右边的字段都是命中的，将右边表达式拼接起来
	if yb.True {
		//因为可能存在||的情况 x为false y为true 这样的情况下不需要拼接条件
		if xb.True {
			exprCollect += fmt.Sprintf("%s%s", condition, yb.Expr)
		} else {
			exprCollect += fmt.Sprintf("%s", yb.Expr)
		}
		fields = append(fields, yb.Fields...)
		if len(yb.FieldMap) > 0 {
			bn.FieldMap = mergeFieldMap(yb.FieldMap, bn.FieldMap)
		}
	}
	bn.Fields = fields
	bn.Expr = exprCollect
	return bn
}

func AstExprToString(expr ast.Expr) string {
	buf := bytes.NewBuffer([]byte{})
	_ = printer.Fprint(buf, token.NewFileSet(), expr)
	return buf.String()
}

func mergeFieldMap(m1 map[string]string, m2 map[string]string) map[string]string {
	for k, v := range m2 {
		t, ok := m1[k]
		if ok {
			m1[k] = fmt.Sprintf("%s && %s", t, v)
		} else {
			m1[k] = v
		}
	}
	return m1
}
