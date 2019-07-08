package atoms

import (
	"fmt"
	"math/big"
)

// Rational represent rational numbers
type Rational struct {
	numerator   Integer
	denominator Integer
}

// NewRational return a 0 rational
func NewRational() Rational {
	return NewRationalFromInteger(0)
}

// NewRationalFromInteger return a rational object.
func NewRationalFromInteger(n int64) Rational {
	return NewRationalFromIntegers(n, 1)
}

// NewRationalFromIntegers return a rational form Integers
func NewRationalFromIntegers(n int64, m int64) Rational {
	newrational := Rational{}

	numerator := NewIntegerFromInt(n)
	denominator := NewIntegerFromInt(m)

	newrational.numerator = numerator
	newrational.denominator = denominator

	return newrational.normalize()
}

// NewRationalFromString return a rational form Integers
func NewRationalFromString(s1 string, s2 string, radix int) (Rational, error) {
	newrational := Rational{}

	numerator, err := NewIntegerFromString(s1, radix)
	if err != nil {
		return newrational, err
	}
	denominator, err := NewIntegerFromString(s2, radix)

	newrational.numerator = numerator
	newrational.denominator = denominator

	return newrational.normalize(), err
}

// Compare two rational numbers
func (num Rational) Compare(y Rational) int {
	cmp := num.Substract(y)
	return cmp.numerator.Compare(NewIntegerFromInt(0))
}

// Add two rational numbers.
func (num Rational) Add(y Rational) Rational {
	num.numerator = num.numerator.Multiply(y.denominator).Add(y.numerator.Multiply(num.denominator))
	num.denominator = num.denominator.Multiply(y.denominator)

	return num.normalize()
}

// Substract two rational numbers.
func (num Rational) Substract(y Rational) Rational {
	num.numerator = num.numerator.Multiply(y.denominator).Substract(y.numerator.Multiply(num.denominator))
	num.denominator = num.denominator.Multiply(y.denominator)

	return num.normalize()
}

// Multiply two rational numbers.
func (num Rational) Multiply(y Rational) Rational {
	num.numerator = num.numerator.Multiply(y.numerator)
	num.denominator = num.denominator.Multiply(y.denominator)

	return num.normalize()
}

// Divide two rational numbers.
func (num Rational) Divide(y Rational) Rational {
	num.numerator = num.numerator.Multiply(y.denominator)
	num.denominator = num.denominator.Multiply(y.numerator)

	return num.normalize()
}

// N return numerical value.
func (num Rational) N() *big.Float {
	numeratorN := new(big.Float).SetInt(num.numerator.val)
	denominatorN := new(big.Float).SetInt(num.denominator.val)

	return new(big.Float).Quo(numeratorN, denominatorN)
}

// String return string representation of rational.
func (num Rational) String() string {
	if num.denominator.Compare(ONE) == 0 {
		return fmt.Sprintf("%s", num.numerator)
	}
	return fmt.Sprintf("%s/%s", num.numerator, num.denominator)
}

// ToLaTeX return a LaTeX representation of rational
func (num Rational) ToLaTeX() string {
	if num.denominator.Compare(ONE) == 0 {
		return fmt.Sprintf("%s", num.numerator)
	}
	return fmt.Sprintf("\\frac{%s}{%s}", num.numerator, num.denominator)
}

func (num Rational) normalize() Rational {
	if num.numerator.Compare(ZERO) == 0 {
		return num
	}

	gcd := num.numerator.Abs().GCD(num.denominator.Abs())

	num.numerator = num.numerator.Divide(gcd)
	num.denominator = num.denominator.Divide(gcd)
	return num
}
