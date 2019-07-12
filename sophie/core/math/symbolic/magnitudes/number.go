package magnitudes

import (
	"github.com/JoseCarlosGarcia95/FreeWolf/sophie/core/math/symbolic"
	"github.com/JoseCarlosGarcia95/FreeWolf/sophie/core/math/symbolic/atoms"
)

// Number represent variables.
type Number struct {
	value atoms.Number
}

// NewSymbolicNumber return a new number.
func NewSymbolicNumber() Number {
	return NewSymbolicNumberFromInteger(0)
}

// NewSymbolicNumberFromInteger return a new number initialize by integer.
func NewSymbolicNumberFromInteger(num int64) Number {
	return NewSymbolicNumberFromNumber(atoms.NewComplexFromInteger(num))
}

// NewSymbolicNumberFromNumber return a symbolic number.
func NewSymbolicNumberFromNumber(num atoms.Number) Number {
	newNumber := Number{}
	newNumber.value = num
	return newNumber
}

// Val return numeric value of variable
func (num Number) Val() interface{} {
	return num.value
}

// Set numeric value of variable
func (num Number) Set(value interface{}) symbolic.Magnitude {
	num.value = value.(atoms.Number)
	return num
}

// Type return terms type.
func (num Number) Type() symbolic.SymbolType {
	return symbolic.SymbolTypeNumber
}

// String return string representation.
func (num Number) String() string {
	return num.value.String()
}
