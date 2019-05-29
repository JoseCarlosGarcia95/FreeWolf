package atoms

import (
	"errors"
	"fmt"
)

// SymbolExpression represent an algebraic symbol
type SymbolExpression struct {
	Symbol string
}

// Evaluate evaluate the current IMathExpression
func (expr SymbolExpression) Evaluate() (IMathExpression, error) {
	return expr, nil
}

// Simplify the current expression
func (expr SymbolExpression) Simplify() (IMathExpression, error) {
	return expr, nil
}

// String return a string from given symbol
func (expr SymbolExpression) String() string {
	return fmt.Sprintf("%s", expr.Symbol)
}

// ToLaTeX return a LaTeX version of symbol
func (expr SymbolExpression) ToLaTeX() string {
	return expr.String()
}

// Sum symbol + other type.
func (expr SymbolExpression) Sum(a IMathExpression) IMathExpression {
	if a.TypeID() == expr.TypeID() {
		b := a.(SymbolExpression)

		if b.Symbol == expr.Symbol {
			return CoefficientExpression{
				Coefficient: NewIntegerFromInteger(2),
				Base:        expr}
		}
	}

	c := MathExpression{}
	new := c.Sum(expr)
	new = new.Sum(a)

	return new
}

// Substract return the substract of two math expression.
func (expr SymbolExpression) Substract(a IMathExpression) IMathExpression {
	if a.TypeID() == expr.TypeID() {
		b := a.(SymbolExpression)

		if b.Symbol == expr.Symbol {
			return NewIntegerFromInteger(0)
		}
	}

	c := MathExpression{}
	new := c.Sum(expr)
	new = new.Substract(a)

	return new
}

// Multiply return the sum of two math expression.
func (expr SymbolExpression) Multiply(a IMathExpression) IMathExpression {
	if a.TypeID() == expr.TypeID() {
		b := a.(SymbolExpression)

		if b.Symbol == expr.Symbol {
			return ExponentExpression{
				Base:     expr,
				Exponent: NewIntegerFromInteger(2)}
		}

	} else if IsNumber(a) {
		return CoefficientExpression{
			Base:        expr,
			Coefficient: a}
	}

	c := MathExpression{}
	return c.Sum(expr).Multiply(a)
}

// Divide return the sum of two math expression.
func (expr SymbolExpression) Divide(a IMathExpression) (IMathExpression, error) {
	if a.TypeID() == expr.TypeID() {
		b := a.(SymbolExpression)

		if b.Symbol == expr.Symbol {
			return NewIntegerFromInteger(1), nil
		}
	}

	c := MathExpression{}
	new := c.Sum(expr)
	new, err := new.Divide(a)
	return new, err
}

// Compare two IMathExpression return 0, if equal and + if a<
func (expr SymbolExpression) Compare(a IMathExpression) (int, error) {
	return 0, errors.New("symbols couldn't be compared")
}

// Inverse expression by another expression.
func (expr SymbolExpression) Inverse() (IMathExpression, error) {
	a := NewIntegerFromInteger(1)
	return a.Divide(expr)
}

// N return the current numeric value.
func (expr SymbolExpression) N() IMathExpression {
	return expr

}

// TypeID is util to detect different types
func (expr SymbolExpression) TypeID() TypeExpression {
	return TypeExpressionSymbol
}
