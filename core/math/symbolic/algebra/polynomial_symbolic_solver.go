package algebra

import (
	"github.com/JoseCarlosGarcia95/FreeWolf/core/math/symbolic/atoms"
)

// PolynomialSymbolicSolver1th solve 1th degree polynomials.
func PolynomialSymbolicSolver1th(expression atoms.IMathExpression, symbol atoms.SymbolExpression) ([]atoms.IMathExpression, error) {
	var solutions []atoms.IMathExpression

	coeff0 := GetCoefficientByDegree(expression, 0, symbol)
	coeff1 := GetCoefficientByDegree(expression, 1, symbol)

	solution, err := atoms.NewIntegerFromInteger(-1).Multiply(coeff0).Divide(coeff1)

	if err != nil {
		return nil, err
	}

	solution, err = solution.Evaluate()
	solutions = append(solutions, solution)

	return solutions, err
}

// PolynomialSymbolicSolver1th solve 2th degree polynomials.
func PolynomialSymbolicSolver2th(expression atoms.IMathExpression, symbol atoms.SymbolExpression) ([]atoms.IMathExpression, error) {
	var solutions []atoms.IMathExpression

	c := GetCoefficientByDegree(expression, 0, symbol)
	b := GetCoefficientByDegree(expression, 1, symbol)
	a := GetCoefficientByDegree(expression, 2, symbol)

	divisor := atoms.NewIntegerFromInteger(2).Multiply(a)

	discriminant, _ := atoms.ExponentExpression{Base: b, Exponent: atoms.NewIntegerFromInteger(2)}.
		Substract(atoms.NewIntegerFromInteger(4).Multiply(a).Multiply(c)).Evaluate()

	discriminant = atoms.ExponentExpression{Base: discriminant, Exponent: atoms.NewFracFromIntegers(1, 2)}

	solution1 := atoms.NewIntegerFromInteger(-1).Multiply(b).Substract(discriminant)
	solution1, _ = solution1.Divide(divisor)
	solution1, _ = solution1.Evaluate()

	solution2 := atoms.NewIntegerFromInteger(-1).Multiply(b).Sum(discriminant)
	solution2, _ = solution1.Divide(divisor)
	solution2, _ = solution1.Evaluate()

	solutions = append(solutions, solution1)
	solutions = append(solutions, solution2)

	return solutions, nil
}

// PolynomialSymbolicSolver try to solve a polynomial equation in symbolic way.
func PolynomialSymbolicSolver(expression atoms.IMathExpression, symbol atoms.SymbolExpression) ([]atoms.IMathExpression, error) {
	// TODO: Expand expression

	degree := CalculateDegree(expression, symbol)

	if degree == 1 {
		return PolynomialSymbolicSolver1th(expression, symbol)
	} else if degree == 2 {
		return PolynomialSymbolicSolver2th(expression, symbol)
	}

	return nil, nil
}
