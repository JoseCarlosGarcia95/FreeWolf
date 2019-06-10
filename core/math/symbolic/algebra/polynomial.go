package algebra

import (
	"github.com/JoseCarlosGarcia95/FreeWolf/core/math/symbolic/atoms"
)

// IsPolynomial return true if a given math expression is a polynomial.
func IsPolynomial(poly atoms.IMathExpression) bool {
	poly, err := poly.Simplify()

	if err != nil {
		return false
	}

	switch poly.TypeID() {
	case atoms.TypeExpressionSymbol:
	case atoms.TypeExpressionReal:
	case atoms.TypeFracExpression:
	case atoms.TypeIntegerExpression:
		return true
	case atoms.TypeExpressionCoefficient:
		coefficient := poly.(atoms.CoefficientExpression)
		return IsPolynomial(coefficient.Coefficient) || IsPolynomial(coefficient.Base)
	case atoms.TypeExpressionExponent:
		exponent := poly.(atoms.ExponentExpression)
		return IsPolynomial(exponent.Base) && atoms.IsNumber(exponent.Exponent)
	case atoms.TypeExpressionGroup:
		group := poly.(atoms.MathExpression)
		partsLen := len(group.Parts)

		for i := 0; i < partsLen; i++ {
			if !IsPolynomial(group.Parts[i]) {
				return false
			}
		}

		return true
	}

	return false
}

func CalculateDegree(poly atoms.IMathExpression, symbol atoms.SymbolExpression) int {
	degree := 0

	if poly.TypeID() == atoms.TypeExpressionGroup {
		polyExpression := poly.(atoms.MathExpression)

		partsLen := len(polyExpression.Parts)

		for i := 0; i < partsLen; i++ {
			if polyExpression.Operators[i] != atoms.OperatorSum {
				continue
			}

			expression := polyExpression.Parts[i]

			if expression.TypeID() == atoms.TypeExpressionSymbol {
				expressionSymbol := expression.(atoms.SymbolExpression)

				if expressionSymbol == symbol && degree < 1 {
					degree = 1
				}
			} else if expression.TypeID() == atoms.TypeExpressionExponent ||
				expression.TypeID() == atoms.TypeExpressionCoefficient {
				var coefficientExpression atoms.CoefficientExpression

				if expression.TypeID() == atoms.TypeExpressionCoefficient {
					coefficientExpression = expression.(atoms.CoefficientExpression)
				} else {
					coefficientExpression = atoms.CoefficientExpression{
						Base:        expression,
						Coefficient: atoms.NewIntegerFromInteger(1)}
				}

				if coefficientExpression.Base.TypeID() == atoms.TypeExpressionSymbol {
					expressionSymbol := coefficientExpression.Base.(atoms.SymbolExpression)

					if expressionSymbol == symbol && degree < 1 {
						degree = 1
					}
				} else if coefficientExpression.Base.TypeID() == atoms.TypeExpressionExponent {
					expressionExponent := coefficientExpression.Base.(atoms.ExponentExpression)
					testDegree := int(expressionExponent.Exponent.(atoms.IntegerExpression).Value.Int64())

					if expressionExponent.Base.TypeID() == atoms.TypeExpressionSymbol {
						expressionSymbol := expressionExponent.Base.(atoms.SymbolExpression)

						if expressionSymbol == symbol && degree < testDegree {
							degree = testDegree
						}
					}
				}

			}
		}
	}

	return degree
}

// GetCoefficientByDegree calculate coefficient by degree and symbol.
func GetCoefficientByDegree(poly atoms.IMathExpression, degree int, symbol atoms.SymbolExpression) atoms.IMathExpression {
	coefficient := atoms.MathExpression{}.Sum(atoms.NewIntegerFromInteger(0))

	if poly.TypeID() == atoms.TypeExpressionGroup {
		polyExpression := poly.(atoms.MathExpression)

		partsLen := len(polyExpression.Parts)

		for i := 0; i < partsLen; i++ {
			if polyExpression.Operators[i] != atoms.OperatorSum {
				continue
			}

			expression := polyExpression.Parts[i]

			if degree == 0 && atoms.IsNumber(expression) {
				coefficient = coefficient.Sum(expression)
			} else if expression.TypeID() == atoms.TypeExpressionSymbol {
				expressionSymbol := expression.(atoms.SymbolExpression)

				if expressionSymbol == symbol && degree == 1 {
					coefficient = coefficient.Sum(atoms.NewIntegerFromInteger(1))
				} else if expressionSymbol != symbol && degree == 0 {
					coefficient = coefficient.Sum(expressionSymbol)
				}
			} else if expression.TypeID() == atoms.TypeExpressionExponent ||
				expression.TypeID() == atoms.TypeExpressionCoefficient {
				var coefficientExpression atoms.CoefficientExpression

				if expression.TypeID() == atoms.TypeExpressionCoefficient {
					coefficientExpression = expression.(atoms.CoefficientExpression)
				} else {
					coefficientExpression = atoms.CoefficientExpression{
						Base:        expression,
						Coefficient: atoms.NewIntegerFromInteger(1)}
				}

				if coefficientExpression.Base.TypeID() == atoms.TypeExpressionSymbol {
					expressionSymbol := coefficientExpression.Base.(atoms.SymbolExpression)

					if expressionSymbol == symbol && degree == 1 {
						coefficient = coefficient.Sum(coefficientExpression.Coefficient)
					} else if expressionSymbol != symbol && degree == 0 {
						coefficient = coefficient.Sum(coefficientExpression)
					}
				} else if coefficientExpression.Base.TypeID() == atoms.TypeExpressionExponent {
					expressionExponent := coefficientExpression.Base.(atoms.ExponentExpression)
					testDegree := int(expressionExponent.Exponent.(atoms.IntegerExpression).Value.Int64())

					if expressionExponent.Base.TypeID() == atoms.TypeExpressionSymbol {
						expressionSymbol := expressionExponent.Base.(atoms.SymbolExpression)

						if expressionSymbol == symbol && degree == testDegree {
							coefficient = coefficient.Sum(coefficientExpression.Coefficient)
						} else if expressionSymbol != symbol && degree == 0 {
							coefficient = coefficient.Sum(coefficientExpression)
						}
					}
				}

			}
		}
	}

	eval, _ := coefficient.Evaluate()

	return eval
}
