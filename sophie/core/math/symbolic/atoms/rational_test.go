package atoms

import (
	"math/big"
	"testing"
)

func TestNewFracFromString(t *testing.T) {
	var err error

	_, err = NewRationalFromString("2", "4", 10)

	if err != nil {
		t.Errorf("Valid Integer with valid radix should be OK")
	}
}

func TestRationalCompare(t *testing.T) {
	num := NewRationalFromIntegers(1, 2)
	num2 := NewRationalFromIntegers(1, 3)

	if num.Compare(num2) < 0 {
		t.Errorf("Unable to compare rational numbers")
	}

	if num2.Compare(num) > 0 {
		t.Errorf("Unable to compare rational numbers")
	}

	if num2.Compare(num2) != 0 {
		t.Errorf("Unable to compare rational numbers")
	}

}

func TestRationalAdd(t *testing.T) {
	num := NewRationalFromIntegers(1, 2)
	num2 := NewRationalFromIntegers(1, 3)

	testNum := NewRationalFromIntegers(5, 6)
	result := num.Add(num2)
	if result.Compare(testNum) != 0 {
		t.Errorf("Expected %s but %s", testNum, result)
	}
}

func TestRationalSubstract(t *testing.T) {
	num := NewRationalFromIntegers(5, 6)
	num2 := NewRationalFromIntegers(1, 3)

	testNum := NewRationalFromIntegers(1, 2)
	result := num.Substract(num2)

	if result.Compare(testNum) != 0 {
		t.Errorf("Expected %s but %s", testNum, result)
	}
}

func TestRationalMultiply(t *testing.T) {
	num := NewRationalFromIntegers(1, 2)
	num2 := NewRationalFromIntegers(1, 3)

	testNum := NewRationalFromIntegers(1, 6)
	result := num.Multiply(num2)
	if result.Compare(testNum) != 0 {
		t.Errorf("Expected %s but %s", testNum, result)
	}
}

func TestRationalDivide(t *testing.T) {
	num := NewRationalFromIntegers(1, 6)
	num2 := NewRationalFromIntegers(1, 3)

	testNum := NewRationalFromIntegers(1, 2)
	result := num.Divide(num2)
	if result.Compare(testNum) != 0 {
		t.Errorf("Expected %s but %s", testNum, result)
	}
}

func TestRationalN(t *testing.T) {
	num := NewRationalFromIntegers(1, 2)

	if num.N().Cmp(big.NewFloat(0.5)) != 0 {
		t.Errorf("Expected 0.5")
	}
}
