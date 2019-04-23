package atoms

import (
	"errors"
	"fmt"
	"math/big"
)

// RealExpression is a simple implementation of real in freewolf language.
type RealExpression struct {
	Value *big.Float
}

// NewRealFromFloat return a new real object from one float
func NewRealFromFloat(x float64) RealExpression {
	return RealExpression{Value: big.NewFloat(x)}
}

// Simplify return a simplified version of expression.
func (expr RealExpression) Simplify() (IMathExpression, error) {
	return expr, nil
}

// String convert current expression to string.
func (expr RealExpression) String() string {
	return fmt.Sprintf("%v", expr.Value)
}

// ToLaTeX return a representation in LaTeX
func (expr RealExpression) ToLaTeX() string {
	return expr.String()
}

// Sum return the sum of two math expression.
func (expr RealExpression) Sum(a IMathExpression) IMathExpression {
	b := a.N()

	if b.TypeID() == expr.TypeID() {
		return RealExpression{Value: big.NewFloat(0).Add(expr.Value, &(*b.(RealExpression).Value))}
	}

	return expr
}

// Substract return the substract of two math expression.
func (expr RealExpression) Substract(a IMathExpression) IMathExpression {
	b := a.N()

	if b.TypeID() == expr.TypeID() {
		return RealExpression{Value: big.NewFloat(0).Sub(expr.Value, &(*b.(RealExpression).Value))}
	}

	return expr
}

// Multiply return the sum of two math expression.
func (expr RealExpression) Multiply(a IMathExpression) IMathExpression {
	b := a.N()

	if b.TypeID() == expr.TypeID() {
		return RealExpression{Value: big.NewFloat(0).Mul(expr.Value, &(*b.(RealExpression).Value))}
	}

	return expr
}

// Divide return the sum of two math expression.
func (expr RealExpression) Divide(a IMathExpression) (IMathExpression, error) {
	b := a.N()

	if b.TypeID() == expr.TypeID() {
		return RealExpression{Value: big.NewFloat(0).Quo(expr.Value, &(*b.(RealExpression).Value))}, nil
	}

	return expr, nil
}

// Inverse return the inverse of a IMathExpression
func (expr RealExpression) Inverse() (IMathExpression, error) {
	return RealExpression{Value: big.NewFloat(0).Quo(big.NewFloat(1), &(*expr.Value))}, nil
}

// TypeID is util to detect different types
func (expr RealExpression) TypeID() TypeExpression {
	return TypeExpressionReal
}

// Derivative of current expression.
func (expr RealExpression) Derivative() (IMathExpression, error) {
	return RealExpression{Value: big.NewFloat(0)}, nil
}

// N return the current numeric value.
func (expr RealExpression) N() IMathExpression {
	return expr
}

// Compare two given expressions
func (expr RealExpression) Compare(a IMathExpression) (int, error) {
	b := a.N()

	if expr.TypeID() == b.TypeID() {
		c := b.(RealExpression).Value

		return expr.Value.Cmp(c), nil
	}

	return 0, errors.New("unable to compare given types")
}
