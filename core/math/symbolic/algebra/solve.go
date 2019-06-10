package algebra

import (
	"errors"

	"github.com/JoseCarlosGarcia95/FreeWolf/core/math/symbolic/atoms"
)

// Solve find roots using symbolic methods.
func Solve(expression atoms.IMathExpression, symbol atoms.SymbolExpression) ([]atoms.IMathExpression, error) {
	//if IsPolynomial(expression) {
	return PolynomialSymbolicSolver(expression, symbol)
	//	}

	return nil, errors.New("This expression cannot be solved with methods available to Solve")
}
