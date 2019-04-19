package atoms

import (
	"errors"
	"fmt"
	"math/big"
)

// FracExpression represent two values separated by a bar.
type FracExpression struct {
	Numerator   IMathExpression
	Denominator IMathExpression
}

// Simplify could return a new FracExpression or a new Integer expression.
func (expr FracExpression) Simplify() (IMathExpression, error) {
	if expr.Numerator.TypeID() == expr.Denominator.TypeID() && expr.Numerator.TypeID() == TypeIntegerExpression {
		a := expr.Numerator.(IntegerExpression).Value
		b := expr.Denominator.(IntegerExpression).Value

		if b.Cmp(big.NewInt(0)) == 0 {
			return nil, errors.New("denominator should be different of zero")
		}

		c := big.NewInt(1).GCD(nil, nil, big.NewInt(1).Abs(a), &(*b))

		a.Div(a, c)
		b.Div(b, c)

		if b.Cmp(big.NewInt(1)) == 0 {
			return IntegerExpression{Value: a}, nil
		}
	}
	return expr, nil
}

// Sum return a sum.
func (expr FracExpression) Sum(a IMathExpression) IMathExpression {
	if a.TypeID() == TypeIntegerExpression {
		return a.Sum(expr)
	}

	if expr.TypeID() == a.TypeID() {
		simplified, _ := FracExpression{
			Numerator:   expr.Numerator.Multiply(a.(FracExpression).Denominator).Sum(a.(FracExpression).Numerator.Multiply(expr.Denominator)),
			Denominator: expr.Denominator.Multiply(a.(FracExpression).Denominator)}.Simplify()

		return simplified
	}
	return expr
}

// Substract return a sum.
func (expr FracExpression) Substract(a IMathExpression) IMathExpression {
	if a.TypeID() == TypeIntegerExpression {
		return a.Substract(expr)
	}

	if expr.TypeID() == a.TypeID() {
		simplified, _ := FracExpression{
			Numerator:   expr.Numerator.Multiply(a.(FracExpression).Denominator).Substract(a.(FracExpression).Numerator.Multiply(expr.Denominator)),
			Denominator: expr.Denominator.Multiply(a.(FracExpression).Denominator)}.Simplify()

		return simplified
	}
	return expr
}

// Multiply return a simplified version of expression.
func (expr FracExpression) Multiply(a IMathExpression) IMathExpression {
	if a.TypeID() == TypeIntegerExpression {
		return a.Multiply(expr)
	}

	if expr.TypeID() == a.TypeID() {
		simplified, _ := FracExpression{
			Numerator:   expr.Numerator.Multiply(a.(FracExpression).Numerator),
			Denominator: expr.Denominator.Multiply(a.(FracExpression).Denominator)}.Simplify()

		return simplified
	}
	return expr
}

// Divide return a simplified version of expression.
func (expr FracExpression) Divide(a IMathExpression) (IMathExpression, error) {
	a, err := a.Inverse()

	if err != nil {
		return nil, err
	}

	return expr.Multiply(a), nil
}

// Derivative of current expression.
func (expr FracExpression) Derivative() (IMathExpression, error) {
	return IntegerExpression{Value: big.NewInt(0)}, nil
}

// Inverse expression by another expression.
func (expr FracExpression) Inverse() (IMathExpression, error) {
	return FracExpression{
		Numerator:   expr.Denominator,
		Denominator: expr.Numerator}.Simplify()
}

// ToLaTeX return a representation in LaTeX
func (expr FracExpression) ToLaTeX() string {
	return fmt.Sprintf("\\frac{%s}{%s}", expr.Numerator, expr.Denominator)
}

// ToString convert current expression to string.
func (expr FracExpression) String() string {
	return fmt.Sprintf("%s/%s", expr.Numerator, expr.Denominator)
}

// TypeID return an identificator for this type.
func (expr FracExpression) TypeID() TypeExpression {
	return TypeFracExpression
}
