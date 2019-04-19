package vtypes

import (
	"bytes"
)

// TypeExpression represent expression type for easy identify expression types.
type TypeExpression int

const (
	// TypeIntegerExpression represent integers.
	TypeIntegerExpression TypeExpression = 0
	// TypeFracExpression represent fractions.
	TypeFracExpression TypeExpression = 1
	// TypeExpressionGroup represent a group of expression.
	TypeExpressionGroup TypeExpression = 2
)

// OperatorsBetweenExpressions represent abstract type of operation.
type OperatorsBetweenExpressions int

const (
	// OperatorSum represent abstract sum.
	OperatorSum OperatorsBetweenExpressions = 0
	// OperatorMult represent abstract multiplication.
	OperatorMult OperatorsBetweenExpressions = 1
)

// IMathExpression abstract the idea of MathExpression.
type IMathExpression interface {
	Simplify() (IMathExpression, error)
	String() string
	ToLaTeX() string
	TypeID() TypeExpression
	Sum(IMathExpression) IMathExpression
	Multiply(IMathExpression) IMathExpression
	Inverse() (IMathExpression, error)
	Derivative() (IMathExpression, error)
}

// MathExpression abstract the idea of MathExpression.
type MathExpression struct {
	Parts     []IMathExpression
	Operators []OperatorsBetweenExpressions
}

// Sum a new expression to current expression
func (expr MathExpression) Sum(a IMathExpression) IMathExpression {
	expr.Parts = append(expr.Parts, a)
	expr.Operators = append(expr.Operators, OperatorSum)

	return expr
}

// Multiply expression by another expression.
func (expr MathExpression) Multiply(a IMathExpression) IMathExpression {
	expr.Parts = append(expr.Parts, a)
	expr.Operators = append(expr.Operators, OperatorMult)

	return expr
}

// Derivative of the current expression
func (expr MathExpression) Derivative() (IMathExpression, error) {
	return expr, nil
}

// Inverse expression by another expression.
func (expr MathExpression) Inverse() (IMathExpression, error) {
	return expr, nil
}

// Simplify try to Simplify the current expression.
func (expr MathExpression) Simplify() (IMathExpression, error) {
	if len(expr.Parts) == 1 {
		return expr.Parts[0].Simplify()
	}
	return expr, nil
}

// ToString convert current expression to string.
func (expr MathExpression) String() string {
	var buffer bytes.Buffer
	operatorLen := len(expr.Operators)

	for i := 0; i < operatorLen; i++ {
		buffer.WriteString(expr.Operators[i].String())
		buffer.WriteString(expr.Parts[i].String())
	}

	return buffer.String()
}

// ToString convert operator to string.
func (op OperatorsBetweenExpressions) String() string {
	switch op {
	case OperatorSum:
		return "+"
	case OperatorMult:
		return "*"
	}

	return ""
}

// ToLaTeX convert current expression to string.
func (expr MathExpression) ToLaTeX() string {
	var buffer bytes.Buffer
	operatorLen := len(expr.Operators)

	for i := 0; i < operatorLen; i++ {
		buffer.WriteString(expr.Operators[i].String())
		buffer.WriteString(expr.Parts[i].ToLaTeX())
	}

	return buffer.String()
}

// TypeID return type of ID of current MathExpression
func (expr MathExpression) TypeID() TypeExpression {
	return TypeExpressionGroup
}

// Evaluate the current MathExpression
func (expr MathExpression) Evaluate() IMathExpression {

	return expr
}
