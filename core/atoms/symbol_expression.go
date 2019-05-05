package atoms

import (
	"errors"
	"fmt"
)

// SymbolExpression represent an algebraic symbol
type SymbolExpression struct {
	Exponent    IMathExpression
	Coefficient IMathExpression
	Symbol      string
}

// Simplify the current expression
func (expr SymbolExpression) Simplify() (IMathExpression, error) {
	return expr, nil
}

// String return a string from given symbol
func (expr SymbolExpression) String() string {
	cmp, _ := NewIntegerFromInteger(1).Compare(expr.Exponent)

	if cmp == 0 {

		cmp2, _ := NewIntegerFromInteger(1).Compare(expr.Coefficient)

		if cmp2 == 0 {
			return fmt.Sprintf("%s", expr.Symbol)
		}
		return fmt.Sprintf("%s %s", expr.Coefficient, expr.Symbol)
	}
	cmp2, _ := NewIntegerFromInteger(1).Compare(expr.Coefficient)

	if cmp2 == 0 {
		return fmt.Sprintf("%s^%s", expr.Symbol, expr.Exponent)
	}
	return fmt.Sprintf("%s %s ^ %s", expr.Coefficient, expr.Symbol, expr.Exponent)
}

// ToLaTeX return a LaTeX version of symbol
func (expr SymbolExpression) ToLaTeX() string {
	return expr.String()
}

// Sum symbol + other type.
func (expr SymbolExpression) Sum(a IMathExpression) IMathExpression {
	if a.TypeID() == expr.TypeID() {
		b := a.(SymbolExpression)

		cmp, err := b.Exponent.Compare(expr.Exponent)

		if err == nil && cmp == 0 && b.Symbol == expr.Symbol {
			new := SymbolExpression{
				Exponent:    expr.Exponent,
				Coefficient: expr.Coefficient.Sum(b.Coefficient),
				Symbol:      expr.Symbol}
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
func (expr SymbolExpression) Substract(a IMathExpression) IMathExpression {
	if a.TypeID() == expr.TypeID() {
		b := a.(SymbolExpression)

		cmp, err := b.Exponent.Compare(expr.Exponent)

		if err == nil && cmp == 0 && b.Symbol == expr.Symbol {
			new := SymbolExpression{
				Exponent:    expr.Exponent,
				Coefficient: expr.Coefficient.Substract(b.Coefficient),
				Symbol:      expr.Symbol}
			return new
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
			new := SymbolExpression{
				Exponent:    expr.Exponent.Sum(b.Exponent),
				Coefficient: expr.Coefficient.Multiply(b.Coefficient),
				Symbol:      expr.Symbol}

			return new
		}

		new1 := SymbolExpression{
			Exponent:    expr.Exponent,
			Coefficient: expr.Coefficient.Multiply(b.Coefficient),
			Symbol:      expr.Symbol}

		new2 := SymbolExpression{
			Exponent:    b.Exponent,
			Coefficient: NewIntegerFromInteger(1),
			Symbol:      b.Symbol}

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
func (expr SymbolExpression) Divide(a IMathExpression) (IMathExpression, error) {
	if a.TypeID() == expr.TypeID() {
		b := a.(SymbolExpression)

		if b.Symbol == expr.Symbol {
			c, err := expr.Exponent.Divide(b.Exponent)

			new := SymbolExpression{
				Exponent:    c,
				Coefficient: expr.Coefficient.Substract(b.Coefficient),
				Symbol:      expr.Symbol}

			return new, err
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
	return SymbolExpression{
		Exponent:    expr.Exponent.N(),
		Coefficient: expr.Coefficient.N(),
		Symbol:      expr.Symbol}

}

// TypeID is util to detect different types
func (expr SymbolExpression) TypeID() TypeExpression {
	return TypeExpressionSymbol
}
