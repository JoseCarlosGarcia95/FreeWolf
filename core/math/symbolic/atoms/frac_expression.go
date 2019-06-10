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

// NewFracFromIntegers return a new frac object from two integers.
func NewFracFromIntegers(x int64, y int64) FracExpression {
	a := NewIntegerFromInteger(x)
	b := NewIntegerFromInteger(y)

	return FracExpression{Numerator: a, Denominator: b}
}

// Simplify could return a new FracExpression or a new Integer expression.
func (expr FracExpression) Simplify() (IMathExpression, error) {
	if expr.Numerator.TypeID() == expr.Denominator.TypeID() && expr.Numerator.TypeID() == TypeIntegerExpression {
		a := expr.Numerator.(IntegerExpression).Value
		b := expr.Denominator.(IntegerExpression).Value

		if b.Cmp(big.NewInt(0)) == 0 {
			return nil, errors.New("denominator should be different of zero")
		}

		if a.Cmp(big.NewInt(0)) == 0 {
			return NewIntegerFromInteger(0), nil
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
		numerator := expr.Numerator.Multiply(a.(FracExpression).Denominator)
		numerator = numerator.Sum(a.(FracExpression).Numerator.Multiply(expr.Denominator))

		denominator := expr.Denominator.Multiply(a.(FracExpression).Denominator)
		simplified, _ := FracExpression{Numerator: numerator, Denominator: denominator}.Simplify()

		return simplified
	} else if a.TypeID() == TypeExpressionReal {
		return a.Sum(expr)
	}

	c := MathExpression{}
	new := c.Sum(expr)
	new = new.Sum(a)

	return new
}

// Substract return a sum.
func (expr FracExpression) Substract(a IMathExpression) IMathExpression {
	if a.TypeID() == TypeIntegerExpression {
		return a.Substract(expr).Multiply(NewIntegerFromInteger(-1))
	}

	if expr.TypeID() == a.TypeID() {
		simplified, _ := FracExpression{
			Numerator:   expr.Numerator.Multiply(a.(FracExpression).Denominator).Substract(a.(FracExpression).Numerator.Multiply(expr.Denominator)),
			Denominator: expr.Denominator.Multiply(a.(FracExpression).Denominator)}.Simplify()

		return simplified
	} else if a.TypeID() == TypeExpressionReal {
		return a.Substract(expr).Multiply(NewIntegerFromInteger(-1))
	}

	c := MathExpression{}
	new := c.Sum(expr)
	new = new.Substract(a)

	return new
}

// Evaluate evaluate the current IMathExpression
func (expr FracExpression) Evaluate() (IMathExpression, error) {
	return expr.Simplify()
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
	} else if a.TypeID() == TypeExpressionReal || a.TypeID() == TypeExpressionSymbol {
		return a.Multiply(expr)
	}
	c := MathExpression{}
	new := c.Sum(expr)
	new = new.Multiply(a)

	return new
}

// Divide return a simplified version of expression.
func (expr FracExpression) Divide(a IMathExpression) (IMathExpression, error) {
	a, err := a.Inverse()

	if err != nil {
		return nil, err
	}

	return expr.Multiply(a), nil
}

// Compare two IMathExpression return 0, if equal and + if a<
func (expr FracExpression) Compare(a IMathExpression) (int, error) {
	a, _ = a.Simplify()
	result := 0
	var err error

	if expr.Numerator.TypeID() == TypeIntegerExpression &&
		expr.Denominator.TypeID() == TypeIntegerExpression {

		if a.TypeID() == TypeIntegerExpression ||
			(a.TypeID() == TypeFracExpression &&
				a.(FracExpression).Numerator.TypeID() == TypeIntegerExpression &&
				a.(FracExpression).Denominator.TypeID() == TypeIntegerExpression) {
			c := expr.Substract(a)

			if c.TypeID() == TypeIntegerExpression {
				result = c.(IntegerExpression).Value.Cmp(big.NewInt(0))
			} else {
				result = c.(FracExpression).Numerator.(IntegerExpression).Value.Cmp(big.NewInt(0))
			}
		} else if a.TypeID() == TypeExpressionReal {
			result, err = a.Compare(expr)
			result *= -1
		} else {
			err = errors.New("unable to compare given types")
		}
	} else {
		err = errors.New("unable to compare given types")
	}
	return result, err
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

// N return a numeric value.
func (expr FracExpression) N() IMathExpression {
	if expr.Numerator.TypeID() == TypeIntegerExpression && expr.Denominator.TypeID() == TypeIntegerExpression {
		a := expr.Numerator.N()
		b := expr.Denominator.N()

		result, _ := a.Divide(b)

		return result
	}
	return expr
}

// ToString convert current expression to string.
func (expr FracExpression) String() string {
	return fmt.Sprintf("%s/%s", expr.Numerator, expr.Denominator)
}

// TypeID return an identificator for this type.
func (expr FracExpression) TypeID() TypeExpression {
	return TypeFracExpression
}
