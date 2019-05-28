package atoms

import (
	"errors"
	"fmt"
)

// ExponentExpression represent an algebraic symbol
type ExponentExpression struct {
	Exponent    IMathExpression
	Coefficient IMathExpression
	Base        IMathExpression
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
	cmp, _ := NewIntegerFromInteger(1).Compare(expr.Exponent)

	if cmp == 0 {

		cmp2, _ := NewIntegerFromInteger(1).Compare(expr.Coefficient)

		if cmp2 == 0 {
			return fmt.Sprintf("%s", expr.Base)
		}
		return fmt.Sprintf("%s %s", expr.Coefficient, expr.Base)
	}
	cmp2, _ := NewIntegerFromInteger(1).Compare(expr.Coefficient)

	if cmp2 == 0 {
		return fmt.Sprintf("%s^%s", expr.Base, expr.Exponent)
	}
	return fmt.Sprintf("%s %s ^ %s", expr.Coefficient, expr.Base, expr.Exponent)
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
			new := ExponentExpression{
				Exponent:    expr.Exponent,
				Coefficient: expr.Coefficient.Sum(b.Coefficient),
				Base:        expr.Base}
			return new
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
			new := ExponentExpression{
				Exponent:    expr.Exponent,
				Coefficient: expr.Coefficient.Substract(b.Coefficient),
				Base:        expr.Base}
			return new
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
			new := ExponentExpression{
				Exponent:    expr.Exponent.Sum(b.Exponent),
				Coefficient: expr.Coefficient.Multiply(b.Coefficient),
				Base:        expr.Base}

			return new
		}

		new1 := ExponentExpression{
			Exponent:    expr.Exponent,
			Coefficient: expr.Coefficient.Multiply(b.Coefficient),
			Base:        expr.Base}

		new2 := ExponentExpression{
			Exponent:    b.Exponent,
			Coefficient: NewIntegerFromInteger(1),
			Base:        b.Base}

		c := MathExpression{}

		return c.Sum(new1).Multiply(new2)

	} else if IsNumber(a) {
		expr.Coefficient = expr.Coefficient.Multiply(a)
		return expr
	}

	c := MathExpression{}
	return c.Sum(expr).Multiply(a)
}

// Divide return the sum of two math expression.
func (expr ExponentExpression) Divide(a IMathExpression) (IMathExpression, error) {
	if a.TypeID() == expr.TypeID() {
		b := a.(ExponentExpression)

		if b.Base == expr.Base {
			c, err := expr.Exponent.Divide(b.Exponent)

			new := ExponentExpression{
				Exponent:    c,
				Coefficient: expr.Coefficient.Substract(b.Coefficient),
				Base:        expr.Base}

			return new, err
		}
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
		Exponent:    expr.Exponent.N(),
		Coefficient: expr.Coefficient.N(),
		Base:        expr.Base}

}

// TypeID is util to detect different types
func (expr ExponentExpression) TypeID() TypeExpression {
	return TypeExpressionExponent
}
