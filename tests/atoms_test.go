package tests

import (
	"fmt"
	"math/big"
	"testing"

	"github.com/JoseCarlosGarcia95/FreeWolf/core/math/symbolic/atoms"
)

func TestIntegersSum(t *testing.T) {
	a := atoms.NewIntegerFromInteger(1)
	b := atoms.NewIntegerFromInteger(2)
	c := a.Sum(b)

	if c.(atoms.IntegerExpression).Value.Cmp(big.NewInt(3)) != 0 {
		t.Errorf("Sum should be 3 but %s", c)
	}
}

func TestIntegersSubs(t *testing.T) {
	a := atoms.NewIntegerFromInteger(1)
	b := atoms.NewIntegerFromInteger(2)
	c := b.Substract(a)

	if c.(atoms.IntegerExpression).Value.Cmp(big.NewInt(1)) != 0 {
		t.Errorf("Substraction should be 1 but %s", c)
	}
}

func TestIntegersMultiply(t *testing.T) {
	a := atoms.NewIntegerFromInteger(3)
	b := atoms.NewIntegerFromInteger(5)
	c := a.Multiply(b)

	if c.(atoms.IntegerExpression).Value.Cmp(big.NewInt(15)) != 0 {
		t.Errorf("Product should be 15, but %s", c)
	}
}

func TestIntegersInverse(t *testing.T) {
	a := atoms.NewIntegerFromInteger(2)
	b, _ := a.Inverse()

	c := atoms.NewFracFromIntegers(1, 2)

	cmp, err := b.Compare(c)

	if err != nil {
		t.Errorf(err.Error())
	}
	if cmp != 0 {
		t.Errorf("Inverse should be 1/2, but %s", b)
	}
}

func TestIntegersDivison(t *testing.T) {
	a := atoms.NewIntegerFromInteger(2)
	b := atoms.NewIntegerFromInteger(3)

	c, err := a.Divide(b)
	if err != nil {
		t.Errorf(err.Error())
	}

	d := atoms.NewFracFromIntegers(2, 3)

	cmp, err := c.Compare(d)

	if err != nil {
		t.Errorf(err.Error())
	}
	if cmp != 0 {
		t.Errorf("Division should be 2/3, but %s", b)
	}
}

func TestComplexMathExpression(t *testing.T) {
	a := atoms.NewIntegerFromInteger(2)
	b := atoms.NewIntegerFromInteger(3)
	c := atoms.NewFracFromIntegers(1, 2)

	m := atoms.MathExpression{}

	expr := m.Sum(b)
	expr = expr.Sum(a)
	expr = expr.Multiply(c)

	result, err := expr.Simplify()

	if err != nil {
		t.Errorf(err.Error())
	}
	expected := atoms.NewIntegerFromInteger(4)

	cmp, err := result.Compare(expected)
	if err != nil {
		t.Errorf(err.Error())
	}
	if cmp != 0 {
		t.Errorf("Expression should be 4, but %s", result)
	}
}

func TestRealSum(t *testing.T) {
	a := atoms.NewFracFromIntegers(1, 2)
	b := atoms.NewRealFromFloat(1.2)
	c := b.Sum(a)

	expected := atoms.NewRealFromFloat(1.7)

	cmp, err := expected.Compare(c)

	if err != nil {
		t.Errorf(err.Error())
	}
	if cmp != 0 {
		t.Errorf("Expression should be 1.7, but %s", c)
	}
}

func TestRealExpression(t *testing.T) {
	a := atoms.NewFracFromIntegers(1, 2)
	b := atoms.NewRealFromFloat(1.2)

	m := atoms.MathExpression{}
	expr := m.Sum(a)
	expr = expr.Sum(b)

	c, _ := expr.Evaluate()
	expected := atoms.NewRealFromFloat(1.7)

	cmp, err := expected.Compare(c)

	if err != nil {
		t.Errorf(err.Error())
	}
	if cmp != 0 {
		t.Errorf("Expression should be 1.7, but %s", c)
	}
}

func benchmarkIntegerExpression(a int64, b *testing.B) {
	one := atoms.NewIntegerFromInteger(a)
	expr := atoms.MathExpression{}

	new := expr.Sum(atoms.NewIntegerFromInteger(0))
	for n := 0; n < b.N; n++ {
		new = new.Sum(one)
	}

	new.Evaluate()
}

func BenchmarkIntegerExpression1(b *testing.B) {
	benchmarkIntegerExpression(1, b)
}

func BenchmarkIntegerExpression10(b *testing.B) {
	benchmarkIntegerExpression(10, b)
}

func BenchmarkIntegerExpression100000000(b *testing.B) {
	benchmarkIntegerExpression(100000000, b)
}

func benchmarkFracExpression(a int64, c int64, b *testing.B) {
	frac := atoms.NewFracFromIntegers(a, c)
	expr := atoms.MathExpression{}

	new := expr.Sum(atoms.NewIntegerFromInteger(0))
	for n := 0; n < b.N; n++ {
		new = new.Sum(frac)
	}

	new.Evaluate()
}

func BenchmarkFracExpression1x2(b *testing.B) {
	benchmarkFracExpression(1, 2, b)
}

func BenchmarkFracExpression2x3(b *testing.B) {
	benchmarkFracExpression(2, 3, b)
}

func BenchmarkFracExpression13x23(b *testing.B) {
	benchmarkFracExpression(13, 23, b)
}

func BenchmarkRealExpression(b *testing.B) {
	real := atoms.NewRealFromFloat(1)
	expr := atoms.MathExpression{}

	new := expr.Sum(atoms.NewIntegerFromInteger(0))
	for n := 0; n < b.N; n++ {
		new = new.Sum(real)
	}
}

func TestSymbolsBasic(t *testing.T) {
	a := atoms.NewIntegerFromInteger(3)
	b := atoms.NewIntegerFromInteger(2)

	sym1 := atoms.SymbolExpression{Symbol: "x", Exponent: b, Coefficient: a}
	sym2 := atoms.SymbolExpression{Symbol: "y", Exponent: b, Coefficient: a}

	result := a.Multiply(sym1).Multiply(sym2)
	if result.String() != "+27 x ^ 2*y^2" {
		t.Errorf("Expression should be +27 x ^ 2*y^2, but %s", result)
	}
}

func TestSymbolSort(t *testing.T) {
	a := atoms.NewIntegerFromInteger(3)
	b := atoms.NewIntegerFromInteger(2)
	c := atoms.NewIntegerFromInteger(10)

	sym1 := atoms.SymbolExpression{Symbol: "x", Exponent: b, Coefficient: a}
	sym2 := atoms.SymbolExpression{Symbol: "x", Exponent: a, Coefficient: b}
	sym3 := atoms.SymbolExpression{Symbol: "x", Exponent: c, Coefficient: b}

	fmt.Println(a.Sum(sym1).Sum(b).Sum(sym2).Sum(sym3).Sum(sym1).Sum(sym3).Sum(sym2))
	result, err := a.Sum(sym1).Sum(b).Sum(sym2).Sum(sym3).Evaluate()

	if err != nil {
		t.Errorf(err.Error())
	}

	if result.String() != "+2 x ^ 10+2 x ^ 3+3 x ^ 2+5" {
		t.Errorf("Expression should be +2 x ^ 10+2 x ^ 3+3 x ^ 2+5, but %s", result)
	}
}
