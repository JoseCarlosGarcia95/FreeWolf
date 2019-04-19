package vtypes

import (
	"fmt"
	"math/big"
)

// IntegerExpression is a simple implementation of integer in freewolf language.
type IntegerExpression struct {
	Value *big.Int
}

// Simplify return a simplified version of expression.
func (expr IntegerExpression) Simplify() (IMathExpression, error) {
	return expr, nil
}

// Sum return a simplified version of expression.
func (expr IntegerExpression) Sum(a IMathExpression) IMathExpression {

	if a.TypeID() == expr.TypeID() {
		return IntegerExpression{Value: big.NewInt(0).Add(expr.Value, &(*a.(IntegerExpression).Value))}
	}

	if a.TypeID() == TypeFracExpression {
		b := IntegerExpression{Value: big.NewInt(1)}
		return FracExpression{Numerator: expr, Denominator: b}.Sum(a)
	}

	return expr
}

// Multiply the product of two elements.
func (expr IntegerExpression) Multiply(a IMathExpression) IMathExpression {

	if a.TypeID() == expr.TypeID() {
		return IntegerExpression{Value: big.NewInt(1).Mul(expr.Value, &(*a.(IntegerExpression).Value))}
	}

	if a.TypeID() == TypeFracExpression {
		return FracExpression{
			Numerator:   expr,
			Denominator: IntegerExpression{Value: big.NewInt(1)}}.Multiply(a)
	}

	return expr
}

// Inverse expression by another expression.
func (expr IntegerExpression) Inverse() (IMathExpression, error) {
	return FracExpression{
		Numerator:   IntegerExpression{Value: big.NewInt(1)},
		Denominator: expr}.Simplify()
}

// Derivative of current expression.
func (expr IntegerExpression) Derivative() (IMathExpression, error) {
	return IntegerExpression{Value: big.NewInt(0)}, nil
}

// ToLaTeX return a representation in LaTeX
func (expr IntegerExpression) ToLaTeX() string {
	return expr.String()
}

// ToString convert current expression to string.
func (expr IntegerExpression) String() string {
	return fmt.Sprintf("%v", expr.Value)
}

// TypeID return an identificator for this type.
func (expr IntegerExpression) TypeID() TypeExpression {
	return TypeIntegerExpression
}
