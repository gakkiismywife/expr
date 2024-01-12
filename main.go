package main

import (
	"expr/expr"
	"fmt"
)

func main() {
	//exprStr := "圆度 <=0.4 && 扁平度 <= 0.4 && 灰度方差 >= 30 && 灰度方差 <= 40"
	exprStr := "(PhyArea >= 500 && Flatness >= 0.25 && GrayMin <= 80 && GrayMax <= 230)"
	//exprStr := "a == 15 && b == 20 && c == 11"
	//exprStr := "a == 15 && b == 11 && b == 20"
	m := map[string]interface{}{"Flatness": 0.7142857313156128, "GrayMax": 203, "GrayMin": 49, "PhyArea": 31508.06281024}

	engine, err := expr.NewEngine(exprStr)
	if err != nil {
		fmt.Println(err)
		return
	}

	engine.Run(m)
}
