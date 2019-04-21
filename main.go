package main

import (
	"fmt"
	"math/big"

	"github.com/JoseCarlosGarcia95/FreeWolf/core/atoms"
)

func main() {

	var smallnum, _ = new(big.Int).SetString("1", 10)
	var smallnum2, _ = new(big.Int).SetString("2", 10)

	a := atoms.IntegerExpression{Value: smallnum}
	b := atoms.IntegerExpression{Value: smallnum2}
	c := atoms.FracExpression{Numerator: a, Denominator: b}

	m := atoms.MathExpression{}

	expr := m.Sum(c)
	expr = expr.Sum(c)
	expr = expr.Multiply(c)
	expr = expr.Sum(a)

	fmt.Println(expr)

	expr, _ = expr.Simplify()

	fmt.Println(expr)
}
