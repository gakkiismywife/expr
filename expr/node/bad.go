package node

type BadNode struct {
	ErrorMessage string
}

func (bNode BadNode) GetTextValue() string {
	return bNode.ErrorMessage
}

func (bNode BadNode) GetType() Type {
	return TypeBad
}
func (bNode BadNode) GetField() string { return "" }

func NewBadNode(str string) ValueNode {
	return BadNode{
		ErrorMessage: str,
	}
}
