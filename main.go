package main

import (
	"fmt"
	"math/big"

	"github.com/JoseCarlosGarcia95/FreeWolf/core/vtypes"
)

func main() {

	var smallnum, _ = new(big.Int).SetString("1", 10)
	var smallnum2, _ = new(big.Int).SetString("2", 10)

	a := vtypes.IntegerExpression{Value: smallnum}
	b := vtypes.IntegerExpression{Value: smallnum2}
	c := vtypes.FracExpression{Numerator: a, Denominator: b}

	d, _ := a.Sum(b).Sum(c).Inverse()

	fmt.Println(d)

	m := vtypes.MathExpression{}

	expr := m.Sum(a)
	expr = expr.Sum(b)
	expr = expr.Sum(c)
	expr = expr.Multiply(d)

	fmt.Println(expr.(vtypes.MathExpression).Evaluate())
}
