package node

import "fmt"

type IntNode struct {
	Value int64
	Field string
}

func NewInt(Value int64, field string) IntNode {
	return IntNode{
		Value: Value,
		Field: field,
	}
}

func (f IntNode) GetType() Type {
	return TypeInt
}

func (f IntNode) GetTextValue() string {
	return fmt.Sprintf("%v", f.Value)
}

func (f IntNode) GetField() string {
	return f.Field
}
