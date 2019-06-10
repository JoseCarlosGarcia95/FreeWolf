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
func (expression CoefficientExpression) Evaluate() (IMathExpression, error) {
	return expression, nil
}

// Simplify the current expression
func (expression CoefficientExpression) Simplify() (IMathExpression, error) {
	return expression, nil
}

// String return a string from given symbol
func (expression CoefficientExpression) String() string {
	return fmt.Sprintf("%s %s", expression.Coefficient, expression.Base)
}

// ToLaTeX return a LaTeX version of symbol
func (expression CoefficientExpression) ToLaTeX() string {
	return expression.String()
}

// Sum symbol + other type.
func (expression CoefficientExpression) Sum(add IMathExpression) IMathExpression {
	if add.TypeID() == expression.TypeID() {
		addCoefficient := add.(CoefficientExpression)

		if addCoefficient.Base == expression.Base {
			coefficientSum := expression.Coefficient.Sum(addCoefficient.Coefficient)

			if coefficientSum == NewIntegerFromInteger(0) {
				return NewIntegerFromInteger(0)
			} else if coefficientSum == NewIntegerFromInteger(1) {
				return expression.Base
			}

			return CoefficientExpression{
				Coefficient: coefficientSum,
				Base:        expression.Base}
		}
	} else if add == expression.Base {
		return CoefficientExpression{
			Coefficient: expression.Coefficient.Sum(NewIntegerFromInteger(1)),
			Base:        expression.Base}
	}

	new := MathExpression{}.Sum(expression)
	new = new.Sum(add)

	return new
}

// Substract return the substract of two math expression.
func (expression CoefficientExpression) Substract(substract IMathExpression) IMathExpression {
	return expression.Sum(NewIntegerFromInteger(-1).Multiply(substract))
}

// Multiply return the sum of two math expression.
func (expression CoefficientExpression) Multiply(factor IMathExpression) IMathExpression {
	if factor.TypeID() == expression.TypeID() {
		factorCoefficient := factor.(CoefficientExpression)

		return CoefficientExpression{
			Coefficient: expression.Coefficient.Multiply(factorCoefficient.Coefficient),
			Base:        expression.Base.Multiply(factorCoefficient.Base)}

	} else if IsNumber(factor) {
		expression.Coefficient = expression.Coefficient.Multiply(factor)
		return expression
	} else if factor.TypeID() == expression.Base.TypeID() || expression.Base.TypeID() == TypeExpressionExponent {
		expression.Base = expression.Base.Multiply(factor)
		return expression
	}

	return MathExpression{}.Sum(expression).Multiply(factor)
}

// Divide return the sum of two math expression.
func (expression CoefficientExpression) Divide(divisor IMathExpression) (IMathExpression, error) {
	inverse, err := divisor.Inverse()

	if err != nil {
		return nil, err
	}

	return expression.Multiply(inverse), nil
}

// Compare two IMathExpression return 0, if equal and + if a<
func (expression CoefficientExpression) Compare(compare IMathExpression) (int, error) {
	return 0, errors.New("symbols couldn't be compared")
}

// Inverse expression by another expression.
func (expression CoefficientExpression) Inverse() (IMathExpression, error) {
	return NewIntegerFromInteger(1).Divide(expression)
}

// N return the current numeric value.
func (expression CoefficientExpression) N() IMathExpression {
	return CoefficientExpression{
		Coefficient: expression.Coefficient.N(),
		Base:        expression.Base.N()}

}

// TypeID is util to detect different types
func (expression CoefficientExpression) TypeID() TypeExpression {
	return TypeExpressionCoefficient
}
