package atoms

import (
	"errors"
	"fmt"
)

// ExponentExpression represent an algebraic symbol
type ExponentExpression struct {
	Exponent IMathExpression
	Base     IMathExpression
}

// Evaluate evaluate the current IMathExpression
func (expr ExponentExpression) Evaluate() (IMathExpression, error) {
	return expr, nil
}

// Simplify the current expression
func (expr ExponentExpression) Simplify() (IMathExpression, error) {
	return expr, nil
}

// String return a string from given symbol
func (expr ExponentExpression) String() string {
	return fmt.Sprintf("%s ^ %s", expr.Base, expr.Exponent)
}

// ToLaTeX return a LaTeX version of symbol
func (expr ExponentExpression) ToLaTeX() string {
	return expr.String()
}

// Sum symbol + other type.
func (expr ExponentExpression) Sum(a IMathExpression) IMathExpression {
	if a.TypeID() == expr.TypeID() {
		b := a.(ExponentExpression)

		cmp, err := b.Exponent.Compare(expr.Exponent)

		if err == nil && cmp == 0 && b.Base == expr.Base {
			return CoefficientExpression{
				Base:        expr,
				Coefficient: NewIntegerFromInteger(2)}
		}

		if cmp > 0 && err == nil {
			c := MathExpression{}
			new := c.Sum(a)
			new = new.Sum(expr)

			return new
		}
	}

	c := MathExpression{}
	new := c.Sum(expr)
	new = new.Sum(a)

	return new
}

// Substract return the substract of two math expression.
func (expr ExponentExpression) Substract(a IMathExpression) IMathExpression {
	if a.TypeID() == expr.TypeID() {
		b := a.(ExponentExpression)

		cmp, err := b.Exponent.Compare(expr.Exponent)

		if err == nil && cmp == 0 && b.Base == expr.Base {
			return NewIntegerFromInteger(0)
		}
	}

	c := MathExpression{}
	new := c.Sum(expr)
	new = new.Substract(a)

	return new
}

// Multiply return the sum of two math expression.
func (expr ExponentExpression) Multiply(a IMathExpression) IMathExpression {
	if a.TypeID() == expr.TypeID() {
		b := a.(ExponentExpression)

		if b.Base == expr.Base {
			return ExponentExpression{
				Exponent: expr.Exponent.Sum(b.Exponent),
				Base:     expr.Base}
		}

		new1 := ExponentExpression{
			Exponent: expr.Exponent,
			Base:     expr.Base}

		new2 := ExponentExpression{
			Exponent: b.Exponent,
			Base:     b.Base}

		c := MathExpression{}

		return c.Sum(new1).Multiply(new2)

	} else if IsNumber(a) {
		return CoefficientExpression{
			Base:        expr,
			Coefficient: a}
	} else if a.TypeID() == expr.Base.TypeID() && a == expr.Base {
		return ExponentExpression{
			Exponent: expr.Exponent.Sum(NewIntegerFromInteger(1)),
			Base:     expr.Base}
	}

	c := MathExpression{}
	return c.Sum(expr).Multiply(a)
}

// Divide return the sum of two math expression.
func (expr ExponentExpression) Divide(a IMathExpression) (IMathExpression, error) {
	if a.TypeID() == expr.TypeID() {
		b := a.(ExponentExpression)

		if b.Base == expr.Base {
			new := ExponentExpression{
				Exponent: expr.Exponent.Substract(b.Exponent),
				Base:     expr.Base}

			return new, nil
		}
	} else if IsNumber(a) {
		a, err := a.Inverse()

		if err != nil {
			return nil, err
		}

		return CoefficientExpression{
			Base:        expr,
			Coefficient: a}, nil
	} else if a.TypeID() == expr.Base.TypeID() && a == expr.Base {
		newExponent := expr.Exponent.Substract(NewIntegerFromInteger(1))

		if newExponent == NewIntegerFromInteger(0) {
			return NewIntegerFromInteger(1), nil
		}

		return ExponentExpression{
			Exponent: newExponent,
			Base:     expr.Base}, nil
	}

	c := MathExpression{}
	new := c.Sum(expr)
	new, err := new.Divide(a)
	return new, err
}

// Compare two IMathExpression return 0, if equal and + if a<
func (expr ExponentExpression) Compare(a IMathExpression) (int, error) {
	return 0, errors.New("symbols couldn't be compared")
}

// Inverse expression by another expression.
func (expr ExponentExpression) Inverse() (IMathExpression, error) {
	a := NewIntegerFromInteger(1)
	return a.Divide(expr)
}

// N return the current numeric value.
func (expr ExponentExpression) N() IMathExpression {
	return ExponentExpression{
		Exponent: expr.Exponent.N(),
		Base:     expr.Base}

}

// TypeID is util to detect different types
func (expr ExponentExpression) TypeID() TypeExpression {
	return TypeExpressionExponent
}
