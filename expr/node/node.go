package node

import (
	"go/ast"
	"go/token"
	"strconv"
)

type Type int

const (
	TypeInt = iota
	TypeFloat
	TypeBad
	TypeBool
)

type ValueNode interface {
	GetType() Type
	GetTextValue() string
	GetField() string
}

func NewNodeByLit(x *ast.BasicLit) ValueNode {
	switch x.Kind {
	case token.INT:
		i, _ := strconv.ParseFloat(x.Value, 64)
		return NewFloat64(i, "")
	default:
		f, _ := strconv.ParseFloat(x.Value, 64)
		return NewFloat64(f, "")
	}
}
