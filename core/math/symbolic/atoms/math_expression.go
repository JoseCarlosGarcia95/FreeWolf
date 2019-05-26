package atoms

import (
	"bytes"
	"sort"
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
	// TypeExpressionReal represent a real number.
	TypeExpressionReal TypeExpression = 3
	// TypeExpressionSymbol represent a symbol.
	TypeExpressionSymbol TypeExpression = 4
	// TypeExpressionExponent represent an exponent.
	TypeExpressionExponent TypeExpression = 5
)

// OperatorsBetweenExpressions represent abstract type of operation.
type OperatorsBetweenExpressions int

const (
	// OperatorSum represent abstract sum.
	OperatorSum OperatorsBetweenExpressions = 0
	// OperatorMult represent abstract multiplication.
	OperatorMult OperatorsBetweenExpressions = 1
	// OperatorSubstract represent abstract substraction.
	OperatorSubstract OperatorsBetweenExpressions = 2
	// OperatorDivide represent abstract division.
	OperatorDivide OperatorsBetweenExpressions = 3
)

// IMathExpression abstract the idea of MathExpression.
type IMathExpression interface {
	Simplify() (IMathExpression, error)
	String() string
	ToLaTeX() string
	TypeID() TypeExpression
	Sum(IMathExpression) IMathExpression
	Substract(IMathExpression) IMathExpression
	Multiply(IMathExpression) IMathExpression
	Divide(IMathExpression) (IMathExpression, error)
	Inverse() (IMathExpression, error)
	Compare(IMathExpression) (int, error)
	N() IMathExpression
	Evaluate() (IMathExpression, error)
}

// MathExpression abstract the idea of MathExpression.
type MathExpression struct {
	Parts     []IMathExpression
	Operators []OperatorsBetweenExpressions
}

// Sum a new expression to current expression
func (expr MathExpression) Sum(a IMathExpression) IMathExpression {
	expr.Operators = append(expr.Operators, OperatorSum)
	expr.Parts = append(expr.Parts, a)
	return expr
}

// Compare two IMathExpression return 0, if equal and + if a<
func (expr MathExpression) Compare(a IMathExpression) (int, error) {
	return 0, nil
}

// Substract a new expression to current expression
func (expr MathExpression) Substract(a IMathExpression) IMathExpression {
	if a.TypeID() == TypeExpressionGroup {
		b := a.(MathExpression)
		c := NewIntegerFromInteger(-1)

		for i := 0; i < len(b.Operators); i++ {
			expr.Operators = append(expr.Operators, b.Operators[i])
			expr.Parts = append(expr.Parts, b.Parts[i].Multiply(c))
		}
	} else {
		expr.Operators = append(expr.Operators, OperatorSubstract)
		expr.Parts = append(expr.Parts, a)
	}

	return expr
}

// Multiply expression by another expression.
func (expr MathExpression) Multiply(a IMathExpression) IMathExpression {
	expr.Parts = append(expr.Parts, a)
	expr.Operators = append(expr.Operators, OperatorMult)

	return expr
}

// Divide expression by another expression.
func (expr MathExpression) Divide(a IMathExpression) (IMathExpression, error) {
	expr.Parts = append(expr.Parts, a)
	expr.Operators = append(expr.Operators, OperatorDivide)

	return expr, nil
}

// Inverse expression by another expression.
func (expr MathExpression) Inverse() (IMathExpression, error) {

	return expr, nil
}

// SortNumbersInExpression put numbers at the end of expression
func (expr MathExpression) SortNumbersInExpression() MathExpression {
	partsLen := len(expr.Parts)

	if !IsNumber(expr.Parts[partsLen-1]) {
		numIndex := -1

		for i := 0; i < partsLen; i++ {
			if IsNumber(expr.Parts[i]) {
				numIndex = i
			}
		}

		tmp := expr.Parts[partsLen-1]
		expr.Parts[partsLen-1] = expr.Parts[numIndex]
		expr.Parts[numIndex] = tmp
	}

	return expr
}

// SortSymbolsInExpression work for sorting symbols on math expression
func (expr MathExpression) SortSymbolsInExpression() MathExpression {
	partsLen := len(expr.Parts)
	symbolIndex := 0

	symbols := make(map[int]int)
	var exponents []int

	for i := 0; i < partsLen; i++ {
		if expr.Parts[i].TypeID() == TypeExpressionSymbol &&
			expr.Parts[i].(SymbolExpression).Exponent.TypeID() == TypeIntegerExpression {
			expInt := int(expr.Parts[i].(SymbolExpression).Exponent.(IntegerExpression).Value.Int64())
			symbols[expInt] = i
			exponents = append(exponents, expInt)
		}
	}

	sort.Ints(exponents)

	for i := range exponents {
		exp := exponents[i]

		tmp := expr.Parts[symbolIndex]
		expr.Parts[symbolIndex] = expr.Parts[symbols[exp]]
		expr.Parts[symbols[exp]] = tmp

		symbolIndex++

		for j := range exponents {
			if symbols[exponents[j]] == symbolIndex {
				symbols[exponents[j]] = symbols[exp]
				break
			}
		}
	}

	return expr
}

// SortExpression push to the end number expressions
func (expr MathExpression) SortExpression() MathExpression {
	expr = expr.SortNumbersInExpression()
	expr = expr.SortSymbolsInExpression()

	return expr
}

// Evaluate evaluate the current IMathExpression
func (expr MathExpression) Evaluate() (IMathExpression, error) {
	if len(expr.Parts) == 1 {
		return expr.Parts[0].Evaluate()
	}

	index := 0
	newexpr := expr.ComputeExpression(&index, 1)

	if newexpr.TypeID() == TypeExpressionGroup {
		return newexpr.(MathExpression).SortExpression(), nil
	}

	return newexpr, nil
}

// SimplifyParts simplify every part of the expression.
func (expr MathExpression) SimplifyParts() (MathExpression, error) {
	var err error
	partsLen := len(expr.Parts)

	for i := 0; i < partsLen; i++ {
		expr.Parts[i], err = expr.Parts[i].Simplify()

		if err != nil {
			return expr, err
		}
	}

	return expr, nil
}

// Simplify try to Simplify the current expression.
func (expr MathExpression) Simplify() (IMathExpression, error) {
	evaluate, err := expr.Evaluate()

	if err != nil {
		return nil, err
	}

	if evaluate.TypeID() != expr.TypeID() {
		return evaluate.Simplify()
	}

	exprEval := evaluate.(MathExpression)
	exprEval, err = exprEval.SimplifyParts()

	if err != nil {
		return nil, err
	}

	return exprEval, err
}

// ParenthesisNeeded return true if we need a parenthesis for this expression on printing.
func ParenthesisNeeded(expr IMathExpression) bool {
	if expr.TypeID() != TypeExpressionGroup {
		return false
	}

	operatorLen := len(expr.(MathExpression).Operators)

	for i := 1; i < operatorLen; i++ {
		if expr.(MathExpression).Operators[i] == OperatorSum ||
			expr.(MathExpression).Operators[i] == OperatorSubstract {
			return true
		}
	}
	return false
}

// ToString convert current expression to string.
func (expr MathExpression) String() string {
	var buffer bytes.Buffer
	operatorLen := len(expr.Operators)

	for i := 0; i < operatorLen; i++ {
		parenthesisNeeded := ParenthesisNeeded(expr.Parts[i])

		if i != 0 {
			buffer.WriteString(expr.Operators[i].String())
		}

		if parenthesisNeeded {
			buffer.WriteString("(")
		}
		buffer.WriteString(expr.Parts[i].String())

		if parenthesisNeeded {
			buffer.WriteString(")")
		}
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
	case OperatorDivide:
		return "/"
	case OperatorSubstract:
		return "-"
	}

	return ""
}

// ToLaTeX convert current expression to string.
func (expr MathExpression) ToLaTeX() string {
	var buffer bytes.Buffer
	operatorLen := len(expr.Operators)

	for i := 0; i < operatorLen; i++ {
		parenthesisNeeded := ParenthesisNeeded(expr.Parts[i])

		if i != 0 {
			buffer.WriteString(expr.Operators[i].String())
		}

		if parenthesisNeeded {
			buffer.WriteString("\\left(")
		}

		buffer.WriteString(expr.Parts[i].ToLaTeX())

		if parenthesisNeeded {
			buffer.WriteString("\\right)")
		}
	}

	return buffer.String()
}

// N return a numeric value.
func (expr MathExpression) N() IMathExpression {
	return expr
}

// TypeID return type of ID of current MathExpression
func (expr MathExpression) TypeID() TypeExpression {
	return TypeExpressionGroup
}

// IsNumber return true if IMathExpression is a pure number type.
func IsNumber(a IMathExpression) bool {
	return a.TypeID() == TypeExpressionReal ||
		a.TypeID() == TypeFracExpression ||
		a.TypeID() == TypeIntegerExpression
}
