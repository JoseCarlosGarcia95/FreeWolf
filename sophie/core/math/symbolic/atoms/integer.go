package atoms

import (
	"errors"
	"fmt"
	"math/big"
)

// Integer return a high level number
type Integer struct {
	val *big.Int
}

// NewInteger create a new integer with value 0
func NewInteger() Integer {
	return NewIntegerFromInt(0)
}

// NewIntegerFromString construct a new integer from string.
func NewIntegerFromString(s string, radix int) (Integer, error) {
	var integer Integer
	var success bool
	var err error

	if radix == 0 {
		radix = 10
	}

	integer.val = new(big.Int)
	integer.val, success = integer.val.SetString(s, radix)

	if !success {
		err = errors.New("unable to read integer")
	}

	return integer, err
}

// NewIntegerFromInt construct a new integer from int.
func NewIntegerFromInt(n int64) Integer {
	var integer Integer

	integer.val = new(big.Int)
	integer.val = integer.val.SetInt64(n)

	return integer
}

// Add two integers.
func (num Integer) Add(y Integer) Integer {
	num.val = new(big.Int).Add(num.val, y.val)
	return num
}

// Substract two integers
func (num Integer) Substract(y Integer) Integer {
	num.val = new(big.Int).Sub(num.val, y.val)
	return num
}

// Multiply two integers
func (num Integer) Multiply(y Integer) Integer {
	num.val = new(big.Int).Mul(num.val, y.val)
	return num
}

// Divide two integers
func (num Integer) Divide(y Integer) Integer {
	num.val = new(big.Int).Div(num.val, y.val)
	return num
}

// Mod of two integeres.
func (num Integer) Mod(y Integer) Integer {
	num.val = new(big.Int).Mod(num.val, y.val)
	return num
}

// Pow of current number
func (num Integer) Pow(y Integer) Integer {
	num.val = new(big.Int).Exp(num.val, y.val, nil)
	return num
}

// GCD return GCD from two integers.
func (num Integer) GCD(y Integer) Integer {
	num.val = new(big.Int).GCD(nil, nil, num.val, y.val)
	return num
}

// Abs calculate the absolute value of integer
func (num Integer) Abs() Integer {
	num.val = new(big.Int).Abs(num.val)

	return num
}

// Sqrt of current number
func (num Integer) Sqrt() Integer {
	newInteger := NewIntegerFromInt(0)
	newInteger.val = newInteger.val.Sqrt(num.val)

	return newInteger
}

// Increment add one to current integer
func (num Integer) Increment() Integer {
	num.val = new(big.Int).Add(num.val, ONE.val)
	return num
}

// Decrement substract one from current integer
func (num Integer) Decrement() Integer {
	num.val = new(big.Int).Sub(num.val, ONE.val)
	return num
}

// Compare return -1 if num<y, 0 if x==1, 1 if num>y
func (num Integer) Compare(y Integer) int {
	return num.val.Cmp(y.val)
}

// String generate a human readable string
func (num Integer) String() string {
	return fmt.Sprintf("%v", num.val)
}
