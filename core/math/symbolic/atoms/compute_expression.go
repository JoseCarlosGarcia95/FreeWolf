package atoms

import (
	"errors"
)

// SumExpressionGroupSymbol sum one expression group and symbol.
func SumExpressionGroupSymbol(expr MathExpression, other IMathExpression) IMathExpression {
	/*
		searchExp := other.(SymbolExpression).Exponent
		partsLen := len(expr.Parts)

		for i := 0; i < partsLen; i++ {

			if expr.Parts[i].TypeID() != TypeExpressionSymbol ||
				expr.Operators[i] != OperatorSum ||
				(i+1 < partsLen && expr.Operators[i+1] != OperatorSum) {
				continue
			}

			part := expr.Parts[i].(SymbolExpression)

			if part.Symbol == other.(SymbolExpression).Symbol {
				cmp, _ := part.Exponent.Compare(searchExp)

				if cmp == 0 {
					expr.Parts[i] = expr.Parts[i].Sum(other)
					return expr
				}
			}
		}
	*/
	return nil
}

// SumExpressionNumber sum one expression group and symbol.
func SumExpressionNumber(expr MathExpression, other IMathExpression) IMathExpression {
	partsLen := len(expr.Parts)

	for i := 0; i < partsLen; i++ {

		if !IsNumber(expr.Parts[i]) ||
			expr.Operators[i] != OperatorSum ||
			(i+1 < partsLen && expr.Operators[i+1] != OperatorSum) {
			continue
		}

		expr.Parts[i] = expr.Parts[i].Sum(other)
		return expr
	}

	return nil
}

// SumExpression represent a sum of expression group.
func SumExpression(a IMathExpression, b IMathExpression) IMathExpression {
	var expr IMathExpression
	var other IMathExpression

	if a.TypeID() == TypeExpressionGroup {
		expr = a.(MathExpression)
		other = b
	} else {
		other = a
		expr = b.(MathExpression)
	}

	if other.TypeID() == TypeExpressionSymbol {
		result := SumExpressionGroupSymbol(expr.(MathExpression), other)

		if result != nil {
			return result
		}

	} else if IsNumber(other) {
		result := SumExpressionNumber(expr.(MathExpression), other)

		if result != nil {
			return result
		}
	}

	return nil
}

// SumOperator on compute expression
func SumOperator(a IMathExpression, b IMathExpression) IMathExpression {
	if a.TypeID() == TypeExpressionGroup || b.TypeID() == TypeExpressionGroup {
		result := SumExpression(a, b)

		if result != nil {
			return result
		}
	}
	return a.Sum(b)
}

// ComputeOp calculate the given operation.
func ComputeOp(op OperatorsBetweenExpressions, a IMathExpression, b IMathExpression) (IMathExpression, error) {
	switch op {
	case OperatorSum:
		return SumOperator(a, b), nil
	case OperatorSubstract:
		return SumOperator(a, b.Multiply(NewIntegerFromInteger(-1))), nil
	case OperatorMult:
		return a.Multiply(b), nil
	case OperatorDivide:
		return a.Divide(b)
	}

	return nil, errors.New("undefined operation")
}

// ComputeExpression calculate the current expression.
// Credits:
// Based on: https://eli.thegreenplace.net/2012/08/02/parsing-expressions-by-precedence-climbing
func (expr MathExpression) ComputeExpression(index *int, minPrecedence int) IMathExpression {
	atomLHS := expr.Parts[*index]

	if atomLHS.TypeID() == TypeExpressionGroup {
		atomLHS, _ = atomLHS.(MathExpression).Evaluate()
	}

	for {
		if *index >= len(expr.Operators)-1 {
			break
		}

		operator := expr.Operators[*index+1]
		precedence, right := operator.GetPrecedenceLevel()

		if precedence < minPrecedence {
			break
		}

		nextMinPrec := precedence

		if !right {
			nextMinPrec++
		}

		*index++
		atomRHS := expr.ComputeExpression(index, nextMinPrec)
		atomLHS, _ = ComputeOp(operator, atomLHS, atomRHS)
	}

	return atomLHS
}

// GetPrecedenceLevel calculate the precedence level of the given operator.
func (op OperatorsBetweenExpressions) GetPrecedenceLevel() (int, bool) {
	precedence := 0
	right := false
	switch op {
	case OperatorSum, OperatorSubstract:
		precedence = 1
		right = false
		break

	case OperatorMult, OperatorDivide:
		precedence = 2
		right = false
		break
	}

	return precedence, right
}
