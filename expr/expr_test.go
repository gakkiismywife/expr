package expr

import (
	"expr/expr/node"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestRun(t *testing.T) {
	t.Run("case1", func(t *testing.T) {
		expr := "(a > 5 || c < 10) && b == 20"
		e, _ := NewEngine(expr)

		n := e.Run(map[string]interface{}{"a": 15, "b": 20, "c": 11})
		bn, ok := n.(node.BoolNode)
		require.Equal(t, true, ok)
		require.Equal(t, true, bn.True)

		require.Equal(t, "a > 5 && b == 20", bn.Expr)
	})

	t.Run("case2", func(t *testing.T) {
		expr := "a == 15 && b == 11 && b == 20"
		e, _ := NewEngine(expr)

		n := e.Run(map[string]interface{}{"a": 15, "b": 20, "c": 11})
		bn, ok := n.(node.BoolNode)
		require.Equal(t, true, ok)
		require.Equal(t, false, bn.True)

		require.Equal(t, "", bn.Expr)
	})

	t.Run("case3", func(t *testing.T) {
		expr := "a == 15 && c == 11 && b == 20"
		e, _ := NewEngine(expr)

		n := e.Run(map[string]interface{}{"a": 15, "b": 20, "c": 11})
		bn, ok := n.(node.BoolNode)
		require.Equal(t, true, ok)
		require.Equal(t, true, bn.True)

		require.Equal(t, "a == 15 && c == 11 && b == 20", bn.Expr)
	})

	t.Run("case4", func(t *testing.T) {
		expr := "a == 15 || c == 11 || b == 20"
		e, _ := NewEngine(expr)

		n := e.Run(map[string]interface{}{"a": 15, "b": 20, "c": 11})
		bn, ok := n.(node.BoolNode)
		require.Equal(t, true, ok)
		require.Equal(t, true, bn.True)

		require.Equal(t, "a == 15 || c == 11 || b == 20", bn.Expr)
	})

	t.Run("case4", func(t *testing.T) {
		expr := "(a == 15 && c == 11) || b == 20"
		e, _ := NewEngine(expr)

		n := e.Run(map[string]interface{}{"a": 15, "b": 2, "c": 11})
		bn, ok := n.(node.BoolNode)
		require.Equal(t, true, ok)
		require.Equal(t, true, bn.True)

		require.Equal(t, "a == 15 && c == 11", bn.Expr)
	})

}
