package atoms

import (
	"testing"
)

func TestComplexCompare(t *testing.T) {
	num := NewComplexFromIntegers(1, 2)
	num2 := NewComplexFromIntegers(1, 3)

	_, err := num.Compare(num2)

	if err == nil {
		t.Errorf("Complex number should not be comparables")
	}

	real := NewComplexFromInteger(1)
	real2 := NewComplexFromInteger(2)

	cmp, err := real.Compare(real2)

	if err != nil || cmp > 0 {
		t.Errorf("Comparing two real should be possible")
	}

	cmp, err = num.Compare(num)

	if err != nil || cmp != 0 {
		t.Errorf("Comparing two equals complex should return 0")
	}

}

func TestComplexAdd(t *testing.T) {
	num := NewComplexFromIntegers(1, 2)
	num2 := NewComplexFromIntegers(1, 3)

	testNum := NewComplexFromIntegers(2, 5)
	result := num.Add(num2)

	cmp, err := testNum.Compare(result)

	if err != nil || cmp != 0 {
		t.Errorf("Expected %s but %s", testNum, result)
	}
}

func TestComplexSubstract(t *testing.T) {
	num := NewComplexFromIntegers(2, 5)
	num2 := NewComplexFromIntegers(1, 3)

	testNum := NewComplexFromIntegers(1, 2)
	result := num.Substract(num2)

	cmp, err := testNum.Compare(result)

	if err != nil || cmp != 0 {
		t.Errorf("Expected %s but %s", testNum, result)
	}
}

func TestComplexMultiply(t *testing.T) {
	num := NewComplexFromIntegers(1, 2)
	num2 := NewComplexFromIntegers(1, 3)

	testNum := NewComplexFromIntegers(-5, 5)
	result := num.Multiply(num2)

	cmp, err := testNum.Compare(result)

	if err != nil || cmp != 0 {
		t.Errorf("Expected %s but %s", testNum, result)
	}
}

func TestComplexDivide(t *testing.T) {
	num := NewComplexFromIntegers(-5, 5)
	num2 := NewComplexFromIntegers(1, 3)

	testNum := NewComplexFromIntegers(1, 2)
	result := num.Divide(num2)

	cmp, err := testNum.Compare(result)

	if err != nil || cmp != 0 {
		t.Errorf("Expected %s but %s", testNum, result)
	}
}
