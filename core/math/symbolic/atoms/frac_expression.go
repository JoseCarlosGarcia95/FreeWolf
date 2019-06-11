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
func (expression FracExpression) Simplify() (IMathExpression, error) {
	if expression.Numerator.TypeID() == expression.Denominator.TypeID() &&
		expression.Numerator.TypeID() == TypeIntegerExpression {
		numerator := expression.Numerator.(IntegerExpression).Value
		denominator := expression.Denominator.(IntegerExpression).Value

		if denominator.Cmp(big.NewInt(0)) == 0 {
			return nil, errors.New("denominator should be different of zero")
		}

		if numerator.Cmp(big.NewInt(0)) == 0 {
			return NewIntegerFromInteger(0), nil
		}

		c := big.NewInt(1).GCD(nil, nil, big.NewInt(1).Abs(numerator), &(*denominator))

		numerator.Div(numerator, c)
		denominator.Div(denominator, c)

		if denominator.Cmp(big.NewInt(1)) == 0 {
			return IntegerExpression{Value: numerator}, nil
		}
	}
	return expression, nil
}

// Sum return a sum.
func (expression FracExpression) Sum(add IMathExpression) IMathExpression {
	if add.TypeID() == TypeIntegerExpression {
		return add.Sum(expression)
	}

	if expression.TypeID() == add.TypeID() {
		numerator := expression.Numerator.Multiply(add.(FracExpression).Denominator)
		numerator = numerator.Sum(add.(FracExpression).Numerator.Multiply(expression.Denominator))

		denominator := expression.Denominator.Multiply(add.(FracExpression).Denominator)
		simplified, _ := FracExpression{Numerator: numerator, Denominator: denominator}.Simplify()

		return simplified
	} else if add.TypeID() == TypeExpressionReal {
		return add.Sum(expression)
	}

	return MathExpression{}.Sum(expression).Sum(add)
}

// Substract return a sum.
func (expression FracExpression) Substract(substract IMathExpression) IMathExpression {
	return expression.Sum(NewIntegerFromInteger(-1).Multiply(substract))

}

// Evaluate evaluate the current IMathExpression
func (expression FracExpression) Evaluate() (IMathExpression, error) {
	return expression.Simplify()
}

// Multiply return a simplified version of expression.
func (expression FracExpression) Multiply(factor IMathExpression) IMathExpression {
	if factor.TypeID() == TypeIntegerExpression {
		return factor.Multiply(expression)
	}

	if expression.TypeID() == factor.TypeID() {
		simplified, _ := FracExpression{
			Numerator: expression.Numerator.
				Multiply(factor.(FracExpression).Numerator),
			Denominator: expression.Denominator.
				Multiply(factor.(FracExpression).Denominator)}.Simplify()

		return simplified
	} else if factor.TypeID() == TypeExpressionReal ||
		factor.TypeID() == TypeExpressionSymbol ||
		factor.TypeID() == TypeExpressionCoefficient ||
		factor.TypeID() == TypeExpressionExponent {
		return factor.Multiply(expression)
	}

	return MathExpression{}.Sum(expression).Multiply(factor)
}

// Divide return a simplified version of expression.
func (expression FracExpression) Divide(divisor IMathExpression) (IMathExpression, error) {
	inverse, err := divisor.Inverse()

	if err != nil {
		return nil, err
	}

	return expression.Multiply(inverse), nil
}

// Compare two IMathExpression return 0, if equal and + if a<
func (expression FracExpression) Compare(a IMathExpression) (int, error) {
	a, _ = a.Simplify()
	result := 0
	var err error

	if expression.Numerator.TypeID() == TypeIntegerExpression &&
		expression.Denominator.TypeID() == TypeIntegerExpression {

		if a.TypeID() == TypeIntegerExpression ||
			(a.TypeID() == TypeFracExpression &&
				a.(FracExpression).Numerator.TypeID() == TypeIntegerExpression &&
				a.(FracExpression).Denominator.TypeID() == TypeIntegerExpression) {
			c := expression.Substract(a)

			if c.TypeID() == TypeIntegerExpression {
				result = c.(IntegerExpression).Value.Cmp(big.NewInt(0))
			} else {
				result = c.(FracExpression).Numerator.(IntegerExpression).Value.Cmp(big.NewInt(0))
			}
		} else if a.TypeID() == TypeExpressionReal {
			result, err = a.Compare(expression)
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
func (expression FracExpression) Inverse() (IMathExpression, error) {
	return FracExpression{
		Numerator:   expression.Denominator,
		Denominator: expression.Numerator}.Simplify()
}

// ToLaTeX return a representation in LaTeX
func (expression FracExpression) ToLaTeX() string {
	return fmt.Sprintf("\\frac{%s}{%s}", expression.Numerator, expression.Denominator)
}

// N return a numeric value.
func (expression FracExpression) N() IMathExpression {
	if expression.Numerator.TypeID() == TypeIntegerExpression && expression.Denominator.TypeID() == TypeIntegerExpression {
		a := expression.Numerator.N()
		b := expression.Denominator.N()

		result, _ := a.Divide(b)

		return result
	}
	return expression
}

// ToString convert current expression to string.
func (expression FracExpression) String() string {
	return fmt.Sprintf("%s/%s", expression.Numerator, expression.Denominator)
}

// TypeID return an identificator for this type.
func (expression FracExpression) TypeID() TypeExpression {
	return TypeFracExpression
}
