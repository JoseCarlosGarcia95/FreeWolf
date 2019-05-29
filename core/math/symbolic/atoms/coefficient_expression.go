package atoms

import (
	"errors"
	"fmt"
)

// CoefficientExpression represent an algebraic symbol
type CoefficientExpression struct {
	Coefficient IMathExpression
	Base        IMathExpression
}

// Evaluate evaluate the current IMathExpression
func (expr CoefficientExpression) Evaluate() (IMathExpression, error) {
	return expr, nil
}

// Simplify the current expression
func (expr CoefficientExpression) Simplify() (IMathExpression, error) {
	return expr, nil
}

// String return a string from given symbol
func (expr CoefficientExpression) String() string {
	return fmt.Sprintf("%s %s", expr.Coefficient, expr.Base)
}

// ToLaTeX return a LaTeX version of symbol
func (expr CoefficientExpression) ToLaTeX() string {
	return expr.String()
}

// Sum symbol + other type.
func (expr CoefficientExpression) Sum(a IMathExpression) IMathExpression {
	if a.TypeID() == expr.TypeID() {
		b := a.(CoefficientExpression)

		if b.Base == expr.Base {
			return CoefficientExpression{
				Coefficient: expr.Coefficient.Sum(b.Coefficient),
				Base:        expr.Base}
		}
	} else if a == expr.Base {
		return CoefficientExpression{
			Coefficient: expr.Coefficient.Sum(NewIntegerFromInteger(1)),
			Base:        expr.Base}
	}

	c := MathExpression{}
	new := c.Sum(expr)
	new = new.Sum(a)

	return new
}

// Substract return the substract of two math expression.
func (expr CoefficientExpression) Substract(a IMathExpression) IMathExpression {
	if a.TypeID() == expr.TypeID() {
		b := a.(CoefficientExpression)

		if b.Base == expr.Base {
			return CoefficientExpression{
				Coefficient: expr.Coefficient.Substract(b.Coefficient),
				Base:        expr.Base}
		}
	} else if expr.Base.TypeID() == a.TypeID() && expr.Base == a {
		return CoefficientExpression{
			Coefficient: expr.Coefficient.Substract(NewIntegerFromInteger(1)),
			Base:        expr.Base}
	}

	c := MathExpression{}
	new := c.Sum(expr)
	new = new.Substract(a)

	return new
}

// Multiply return the sum of two math expression.
func (expr CoefficientExpression) Multiply(a IMathExpression) IMathExpression {
	if a.TypeID() == expr.TypeID() {
		b := a.(CoefficientExpression)

		new := CoefficientExpression{
			Coefficient: expr.Coefficient.Multiply(b.Coefficient),
			Base:        expr.Base.Multiply(b.Base)}

		return new

	} else if IsNumber(a) {
		expr.Coefficient = expr.Coefficient.Multiply(a)
		return expr
	} else if a.TypeID() == expr.Base.TypeID() || expr.Base.TypeID() == TypeExpressionExponent {
		expr.Base = expr.Base.Multiply(a)
		return expr
	}

	c := MathExpression{}
	return c.Sum(expr).Multiply(a)
}

// Divide return the sum of two math expression.
func (expr CoefficientExpression) Divide(a IMathExpression) (IMathExpression, error) {
	var err error

	if a.TypeID() == expr.TypeID() {
		b := a.(CoefficientExpression)

		c, err := expr.Base.Divide(b.Base)

		if err != nil {
			return nil, err
		}

		d, err := expr.Coefficient.Divide(b.Base)

		if err != nil {
			return nil, err
		}

		new := CoefficientExpression{
			Coefficient: d,
			Base:        c}

		return new, nil

	} else if IsNumber(a) {
		expr.Coefficient, err = expr.Coefficient.Divide(a)

		cmp, err := expr.Coefficient.Compare(NewIntegerFromInteger(1))

		if cmp == 0 {
			return expr.Base, nil
		}

		return expr, err
	} else if a.TypeID() == expr.Base.TypeID() || expr.Base.TypeID() == TypeExpressionExponent {
		expr.Base, err = expr.Base.Divide(a)

		cmp, err := expr.Base.Compare(NewIntegerFromInteger(1))

		if cmp == 0 {
			return expr.Coefficient, nil
		}

		return expr, err
	}

	c := MathExpression{}
	new := c.Sum(expr)
	new, err = new.Divide(a)
	return new, err
}

// Compare two IMathExpression return 0, if equal and + if a<
func (expr CoefficientExpression) Compare(a IMathExpression) (int, error) {
	return 0, errors.New("symbols couldn't be compared")
}

// Inverse expression by another expression.
func (expr CoefficientExpression) Inverse() (IMathExpression, error) {
	a := NewIntegerFromInteger(1)
	return a.Divide(expr)
}

// N return the current numeric value.
func (expr CoefficientExpression) N() IMathExpression {
	return CoefficientExpression{
		Coefficient: expr.Coefficient.N(),
		Base:        expr.Base.N()}

}

// TypeID is util to detect different types
func (expr CoefficientExpression) TypeID() TypeExpression {
	return TypeExpressionCoefficient
}
