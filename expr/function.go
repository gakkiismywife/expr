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
	bn := n.(node.BoolNode)
	if bn.True == false {
		return bn
	}
	var fields []string
	var exprCollect string
	if xb.True {
		fields = append(fields, xb.Fields...)
		exprCollect = fmt.Sprintf("%s", xb.Expr)
	}
	if yb.True {
		exprCollect += fmt.Sprintf("%s%s", condition, yb.Expr)
		fields = append(fields, yb.Fields...)
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
