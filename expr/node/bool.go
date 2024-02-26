package node

import "fmt"

type BoolNode struct {
	textValue string
	True      bool
	Fields    []string
	Expr      string
	FieldMap  map[string]string
}

func (bNode BoolNode) GetTextValue() string {
	return bNode.textValue
}

func (bNode BoolNode) GetValue() bool {
	return bNode.True
}

func (bNode BoolNode) GetType() Type {
	return TypeBool
}

func (bNode BoolNode) GetField() string {
	return ""
}

func (bNode BoolNode) GetFields() []string {
	return bNode.Fields
}

func NewBoolNode(b bool, filed string) ValueNode {
	return BoolNode{
		True:      b,
		textValue: fmt.Sprintf("%t", b),
		Fields:    []string{filed},
		FieldMap:  make(map[string]string),
	}
}
