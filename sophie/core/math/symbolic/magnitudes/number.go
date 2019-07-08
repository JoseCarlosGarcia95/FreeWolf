package magnitudes

import (
	"github.com/JoseCarlosGarcia95/FreeWolf/sophie/core/math/symbolic"
	"github.com/JoseCarlosGarcia95/FreeWolf/sophie/core/math/symbolic/atoms"
)

// Number represent variables.
type Number struct {
	value atoms.Number
}

// Val return numeric value of variable
func (num *Number) Val() atoms.Number {
	return num.value
}

// Set numeric value of variable
func (num *Number) Set(value atoms.Number) {
	num.value = value
}

// Type return terms type.
func (num Number) Type() symbolic.TermsType {
	return symbolic.TermsTypeNumber
}

// String return string representation.
func (num Number) String() string {
	return num.String()
}
