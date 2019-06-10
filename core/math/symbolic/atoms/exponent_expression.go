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
func (expression ExponentExpression) Evaluate() (IMathExpression, error) {
	return expression, nil
}

// Simplify the current expression
func (expression ExponentExpression) Simplify() (IMathExpression, error) {
	return expression, nil
}

// String return a string from given symbol
func (expression ExponentExpression) String() string {
	return fmt.Sprintf("%s ^ %s", expression.Base, expression.Exponent)
}

// ToLaTeX return a LaTeX version of symbol
func (expression ExponentExpression) ToLaTeX() string {
	return expression.String()
}

// Sum symbol + other type.
func (expression ExponentExpression) Sum(add IMathExpression) IMathExpression {
	if add.TypeID() == expression.TypeID() {
		addExponent := add.(ExponentExpression)

		if addExponent.Exponent == expression.Exponent &&
			addExponent.Base == expression.Base {
			return CoefficientExpression{
				Base:        expression,
				Coefficient: NewIntegerFromInteger(2)}
		}
	}

	new := MathExpression{}.Sum(expression)
	new = new.Sum(add)

	return new
}

// Substract return the substract of two math expression.
func (expression ExponentExpression) Substract(substract IMathExpression) IMathExpression {
	return expression.Sum(NewIntegerFromInteger(-1).Multiply(substract))
}

// Multiply return the sum of two math expression.
func (expression ExponentExpression) Multiply(factor IMathExpression) IMathExpression {
	if factor.TypeID() == expression.TypeID() {
		factorExponent := factor.(ExponentExpression)

		if factorExponent.Base == expression.Base {
			return ExponentExpression{
				Exponent: expression.Exponent.Sum(factorExponent.Exponent),
				Base:     expression.Base}
		}

		return MathExpression{}.
			Sum(ExponentExpression{
				Exponent: expression.Exponent,
				Base:     expression.Base}).
			Multiply(ExponentExpression{
				Exponent: factorExponent.Exponent,
				Base:     factorExponent.Base})

	} else if IsNumber(factor) {
		return CoefficientExpression{
			Base:        expression,
			Coefficient: factor}
	} else if factor.TypeID() == expression.Base.TypeID() &&
		factor == expression.Base {
		return ExponentExpression{
			Exponent: expression.Exponent.Sum(NewIntegerFromInteger(1)),
			Base:     expression.Base}
	}

	return MathExpression{}.Sum(expression).Multiply(factor)
}

// Divide return the sum of two math expression.
func (expression ExponentExpression) Divide(divisor IMathExpression) (IMathExpression, error) {
	inverse, err := divisor.Inverse()

	if err != nil {
		return nil, err
	}

	return expression.Multiply(inverse), nil
}

// Compare two IMathExpression return 0, if equal and + if a<
func (expression ExponentExpression) Compare(compare IMathExpression) (int, error) {
	return 0, errors.New("symbols couldn't be compared")
}

// Inverse expression by another expression.
func (expression ExponentExpression) Inverse() (IMathExpression, error) {
	return NewIntegerFromInteger(1).Divide(expression)
}

// N return the current numeric value.
func (expression ExponentExpression) N() IMathExpression {
	return ExponentExpression{
		Exponent: expression.Exponent.N(),
		Base:     expression.Base}

}

// TypeID is util to detect different types
func (expression ExponentExpression) TypeID() TypeExpression {
	return TypeExpressionExponent
}
