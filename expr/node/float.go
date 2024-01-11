package node

import "fmt"

type FloatNode struct {
	Value float64
	Field string
}

func NewFloat64(Value float64, field string) FloatNode {
	return FloatNode{
		Value: Value,
		Field: field,
	}
}

func (f FloatNode) GetType() Type {
	return TypeFloat
}

func (f FloatNode) GetTextValue() string {
	return fmt.Sprintf("%v", f.Value)
}

func (f FloatNode) GetField() string {
	return f.Field
}
