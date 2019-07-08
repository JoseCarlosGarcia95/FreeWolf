package atoms

// ZERO Represent zero Integer.
var ZERO = NewIntegerFromInt(0)

// ONE Represent one Integer.
var ONE = NewIntegerFromInt(1)

// RationalZero represent zero rational number.
var RationalZero = NewRationalFromInteger(0)

// ComplexZero represent zero
var ComplexZero = NewComplex()

// ComplexOne represent one
var ComplexOne = NewComplexFromInteger(1)

// Number define default number class.
type Number = Complex
