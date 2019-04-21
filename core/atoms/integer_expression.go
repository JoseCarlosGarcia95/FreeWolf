package atoms

import (
	"errors"
	"fmt"
	"math/big"
)

// IntegerExpression is a simple implementation of integer in freewolf language.
type IntegerExpression struct {
	Value *big.Int
}

// NewIntegerFromInteger return a IntegerExpression from language integer.
func NewIntegerFromInteger(a int64) IntegerExpression {
	nint := IntegerExpression{Value: big.NewInt(a)}
	return nint
}

// Simplify return a simplified version of expression.
func (expr IntegerExpression) Simplify() (IMathExpression, error) {
	return expr, nil
}

// Compare two IMathExpression return 0, if equal and + if a<
func (expr IntegerExpression) Compare(a IMathExpression) (int, error) {
	a, _ = a.Simplify()
	result := 0
	var err error

	if a.TypeID() == TypeFracExpression {
		result, err = a.Compare(expr)
		result = result * -1
	} else if a.TypeID() == TypeIntegerExpression {
		c := expr.Substract(a).(IntegerExpression)
		result = c.Value.Cmp(big.NewInt(0))
	} else {
		err = errors.New("unable to compare given expression")
	}
	return result, err
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

// Substract return a simplified version of expression.
func (expr IntegerExpression) Substract(a IMathExpression) IMathExpression {

	if a.TypeID() == expr.TypeID() {
		return IntegerExpression{Value: big.NewInt(0).Sub(expr.Value, &(*a.(IntegerExpression).Value))}
	}

	if a.TypeID() == TypeFracExpression {
		b := IntegerExpression{Value: big.NewInt(1)}
		return FracExpression{Numerator: expr, Denominator: b}.Substract(a)
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

// Divide the product of two elements.
func (expr IntegerExpression) Divide(a IMathExpression) (IMathExpression, error) {
	a, err := a.Inverse()

	if err != nil {
		return nil, err
	}

	return expr.Multiply(a), nil
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
