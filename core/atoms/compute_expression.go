package atoms

import (
	"errors"
	"fmt"
)

// ComputeOp calculate the given operation.
func ComputeOp(op OperatorsBetweenExpressions, a IMathExpression, b IMathExpression) (IMathExpression, error) {
	switch op {
	case OperatorSum:
		return a.Sum(b), nil
	case OperatorSubstract:
		return a.Substract(b), nil
	case OperatorMult:
		return a.Multiply(b), nil
	case OperatorDivide:
		return a.Divide(b)
	}

	return nil, errors.New("undefined operation")
}

// ComputeExpression calculate the current expression.
// Source: https://eli.thegreenplace.net/2012/08/02/parsing-expressions-by-precedence-climbing
func (expr MathExpression) ComputeExpression(index int, minPrecedence int) IMathExpression {
	var err error
	atomLHS := expr.Parts[index]

	for index+1 < len(expr.Parts) {
		operator := expr.Operators[index+1]

		precedence, right := operator.GetPrecedenceLevel()

		if precedence < minPrecedence {
			break
		}

		nextMintPrec := precedence

		if !right {
			nextMintPrec++
		}

		index++

		atomRHS := expr.ComputeExpression(index, nextMintPrec)
		atomLHS, err = ComputeOp(operator, atomLHS, atomRHS)

		if err != nil {
			fmt.Println(err)
			return nil
		}
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
