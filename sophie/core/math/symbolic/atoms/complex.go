package atoms

import (
	"errors"
	"fmt"
)

// Complex represent complex numbers
type Complex struct {
	real Rational
	im   Rational
}

// NewComplex return a new complex number initialize to zero.
func NewComplex() Complex {
	return NewComplexFromInteger(0)
}

// NewComplexFromInteger return a new complex number with real part.
func NewComplexFromInteger(real int64) Complex {
	return NewComplexFromIntegers(real, 0)
}

// NewComplexFromIntegers return a new complex number
func NewComplexFromIntegers(real int64, im int64) Complex {
	return NewComplexFromRationals(NewRationalFromInteger(real), NewRationalFromInteger(im))
}

// NewComplexFromRational return a new complex number from rational.
func NewComplexFromRational(real Rational) Complex {
	return NewComplexFromRationals(real, NewRational())
}

// NewComplexFromRationals return a new complex number from rationals
func NewComplexFromRationals(real Rational, im Rational) Complex {
	newComplex := Complex{}

	newComplex.real = real
	newComplex.im = im

	return newComplex
}

// Add two complex numbers.
func (num Complex) Add(y Complex) Complex {
	num.real = num.real.Add(y.real)
	num.im = num.im.Add(y.im)
	return num
}

// Substract two complex numbers.
func (num Complex) Substract(y Complex) Complex {
	num.real = num.real.Substract(y.real)
	num.im = num.im.Substract(y.im)
	return num
}

// Multiply two complex numbers.
func (num Complex) Multiply(y Complex) Complex {
	realPart := num.real.Multiply(y.real).Substract(num.im.Multiply(y.im))
	imPart := num.real.Multiply(y.im).Add(num.im.Multiply(y.real))

	num.real = realPart
	num.im = imPart
	return num
}

// Divide two complex numbers.
func (num Complex) Divide(y Complex) Complex {
	divisor := y.real.Multiply(y.real).Add(y.im.Multiply(y.im))
	realPart := num.real.Multiply(y.real).Add(num.im.Multiply(y.im)).Divide(divisor)
	imPart := num.im.Multiply(y.real).Substract(num.real.Multiply(y.im)).Divide(divisor)

	num.real = realPart
	num.im = imPart
	return num
}

// Compare return 0 if complex numebr are equal.
func (num Complex) Compare(y Complex) (int, error) {

	if num.real.Compare(y.real) == 0 && num.im.Compare(y.im) == 0 {
		return 0, nil
	}

	if !num.IsReal() || !y.IsReal() {
		return 0, errors.New("Unable to compare complex numbers")
	}

	return num.real.Compare(y.real), nil
}

// IsReal return if complex number is a pure real number.
func (num Complex) IsReal() bool {
	return num.im.Compare(RationalZero) == 0
}

func (num Complex) String() string {
	if num.IsReal() {
		return num.real.String()
	}

	return fmt.Sprintf("%s + %s I", num.real, num.im)
}
