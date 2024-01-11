package expr

import (
	"expr/expr/node"
	"fmt"
	"github.com/shopspring/decimal"
	"go/token"
)

type BinaryIntExpr struct{}

type BinaryFloatExpr struct{}

type BinaryBoolExpr struct{}

func (b BinaryIntExpr) Invoke(x, y node.ValueNode, op token.Token) node.ValueNode {
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

	}
	return bn
}

func (b BinaryFloatExpr) Invoke(x, y node.ValueNode, op token.Token) node.ValueNode {
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
	}
	return bn
}

func (b BinaryBoolExpr) Invoke(x, y node.ValueNode, op token.Token) node.ValueNode {
	xb, xok := x.(node.BoolNode)
	yb, yok := y.(node.BoolNode)

	if !xok || !yok {
		return node.NewBadNode(x.GetTextValue() + y.GetTextValue())
	}
	var field = xb.GetField()
	var n node.ValueNode
	switch op {
	case token.LAND:
		n = node.NewBoolNode(xb.True && yb.True, field)
	case token.LOR:
		n = node.NewBoolNode(xb.True || yb.True, field)
	default:
		return node.NewBadNode(fmt.Sprintf("unsupported binary operator: %s", op.String()))
	}
	n = node.NewBoolNode(xb.True || yb.True, field)
	bn := n.(node.BoolNode)
	var fields []string
	if xb.True {
		fields = append(fields, xb.Fields...)
	}
	if yb.True {
		fields = append(fields, yb.Fields...)
	}
	bn.Fields = fields
	return bn
}
