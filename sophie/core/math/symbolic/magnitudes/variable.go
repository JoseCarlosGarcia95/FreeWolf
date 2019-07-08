package magnitudes

import (
	"github.com/JoseCarlosGarcia95/FreeWolf/sophie/core/math/symbolic"
	"github.com/JoseCarlosGarcia95/FreeWolf/sophie/core/math/symbolic/atoms"
)

// Variable represent variables.
type Variable struct {
	name  string
	value atoms.Number
}

// Val return numeric value of variable
func (variable Variable) Val() atoms.Number {
	return variable.value
}

// Set numeric value of variable
func (variable Variable) Set(value atoms.Number) {
	variable.value = value
}

// Type return terms type.
func (variable Variable) Type() symbolic.SymbolType {
	return symbolic.SymbolTypeVariable
}

// String return string representation.
func (variable Variable) String() string {
	return variable.name
}
