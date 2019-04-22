package tests

import (
	"math/big"
	"testing"

	"github.com/JoseCarlosGarcia95/FreeWolf/core/atoms"
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
