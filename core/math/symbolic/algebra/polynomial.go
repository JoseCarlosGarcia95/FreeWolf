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
