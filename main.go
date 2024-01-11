package main

import (
	"expr/expr"
	"fmt"
)

func main() {
	//exprStr := "(a > 5 || c < 10) && b == 20"
	//exprStr := "a == 15 && b == 20 && c == 11"
	exprStr := "a == 15 && b == 11 && b == 20"
	m := map[string]interface{}{"a": 15, "b": 20, "c": 11}

	engine, err := expr.NewEngine(exprStr)
	if err != nil {
		fmt.Println(err)
		return
	}

	engine.Run(m)
}
