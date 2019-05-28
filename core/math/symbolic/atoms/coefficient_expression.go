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
			new := CoefficientExpression{
				Coefficient: expr.Coefficient.Substract(b.Coefficient),
				Base:        expr.Base}
			return new
		}
	}

	c := MathExpression{}
	new := c.Sum(expr)
	new = new.Sum(a)

	return new
}

// Multiply return the sum of two math expression.
func (expr CoefficientExpression) Multiply(a IMathExpression) IMathExpression {
	if a.TypeID() == expr.TypeID() {
		b := a.(CoefficientExpression)

		new := CoefficientExpression{
			Coefficient: expr.Coefficient.Multiply(b.Coefficient),
			Base:        expr.Base.Multiply(expr.Base)}

		return new

	} else if IsNumber(a) {
		expr.Coefficient = expr.Coefficient.Multiply(a)
		return expr
	}

	c := MathExpression{}
	return c.Sum(expr).Multiply(a)
}

// Divide return the sum of two math expression.
func (expr CoefficientExpression) Divide(a IMathExpression) (IMathExpression, error) {
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

		return new, err

	}

	c := MathExpression{}
	new := c.Sum(expr)
	new, err := new.Divide(a)
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
		Base:        expr.Base}

}

// TypeID is util to detect different types
func (expr CoefficientExpression) TypeID() TypeExpression {
	return TypeExpressionCoefficient
}
