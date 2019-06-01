package tests

import (
	"testing"

	"github.com/JoseCarlosGarcia95/FreeWolf/core/math/symbolic/algebra"
	"github.com/JoseCarlosGarcia95/FreeWolf/core/math/symbolic/atoms"
)

func TestCheckPolynomial(t *testing.T) {
	a := atoms.NewIntegerFromInteger(3)
	b := atoms.NewIntegerFromInteger(2)
	c := atoms.NewIntegerFromInteger(10)

	sym1 := atoms.SymbolExpression{Symbol: "x"}

	exp1 := atoms.ExponentExpression{Exponent: b, Base: sym1}
	exp2 := atoms.ExponentExpression{Exponent: a, Base: sym1}
	exp3 := atoms.ExponentExpression{Exponent: c, Base: sym1}

	coeff1 := atoms.CoefficientExpression{Coefficient: a, Base: exp1}
	coeff2 := atoms.CoefficientExpression{Coefficient: b, Base: exp2}
	coeff3 := atoms.CoefficientExpression{Coefficient: b, Base: exp3}

	result := a.Sum(coeff1).Sum(b).Sum(coeff2).Sum(coeff3)

	if !algebra.IsPolynomial(result) {
		t.Errorf("%s should be a polynomial", result)
	}

	exp4 := atoms.ExponentExpression{Exponent: sym1, Base: sym1}

	if algebra.IsPolynomial(exp4) {
		t.Errorf("%s should be a polynomial", exp4)
	}

	if !algebra.IsPolynomial(result.Multiply(result)) {
		t.Errorf("%s should be a polynomial", result.Multiply(result))
	}

}
