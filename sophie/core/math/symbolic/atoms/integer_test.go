package atoms

import (
	"fmt"
	"testing"
)

func TestNewIntegerFromString(t *testing.T) {
	var err error

	_, err = NewIntegerFromString("10000", 10)

	if err != nil {
		t.Errorf("Valid Integer with valid radix should be OK")
	}

	_, err = NewIntegerFromString("a12", 16)

	if err != nil {
		t.Errorf("Valid Integer with valid radix should be OK")
	}

	_, err = NewIntegerFromString("a12", 10)

	if err == nil {
		t.Errorf("Invalid Integer with invalid radix should be KO")
	}
}

func TestCompare(t *testing.T) {
	testNum := NewIntegerFromInt(1)
	num := NewIntegerFromInt(0)

	if num.Compare(testNum) > 0 {
		t.Errorf("Expected %s but %s", testNum, num)
	}

	if testNum.Compare(num) < 0 {
		t.Errorf("Expected %s but %s", testNum, num)
	}

	if num.Compare(num) != 0 {
		t.Errorf("Expected %s but %s", testNum, num)
	}
}

func TestIntegerIncrement(t *testing.T) {
	testNum := NewIntegerFromInt(1)
	num := NewIntegerFromInt(0)

	num = num.Increment()

	if num.Compare(testNum) != 0 {
		t.Errorf("Expected %s but %s", testNum, num)
	}
}

func TestDecrement(t *testing.T) {
	testNum := NewIntegerFromInt(1)
	num := NewIntegerFromInt(2)

	num = num.Decrement()

	if num.Compare(testNum) != 0 {
		t.Errorf("Expected %s but %s", testNum, num)
	}
}

func TestAdd(t *testing.T) {
	testNum := NewIntegerFromInt(2)
	num := NewIntegerFromInt(1)

	num = num.Add(num)

	if num.Compare(testNum) != 0 {
		t.Errorf("Expected %s but %s", testNum, num)
	}
}

func TestSubstract(t *testing.T) {
	testNum := NewIntegerFromInt(0)
	num := NewIntegerFromInt(1)

	num = num.Substract(num)

	if num.Compare(testNum) != 0 {
		t.Errorf("Expected %s but %s", testNum, num)
	}
}

func TestMultiply(t *testing.T) {
	testNum := NewIntegerFromInt(4)
	num := NewIntegerFromInt(2)

	num = num.Multiply(num)

	if num.Compare(testNum) != 0 {
		t.Errorf("Expected %s but %s", testNum, num)
	}
}

func TestAbs(t *testing.T) {
	testNum := NewIntegerFromInt(2)
	num := NewIntegerFromInt(-2)

	num = num.Abs()

	if num.Compare(testNum) != 0 {
		t.Errorf("Expected %s but %s", testNum, num)
	}
}

func TestGCD(t *testing.T) {
	num1 := NewIntegerFromInt(25)
	num2 := NewIntegerFromInt(15)
	testNum := NewIntegerFromInt(5)

	result := num1.GCD(num2)
	if result.Compare(testNum) != 0 {
		t.Errorf("Expected %s but %s", testNum, result)
	}
}

func TestDivide(t *testing.T) {
	num1 := NewIntegerFromInt(125)
	num2 := NewIntegerFromInt(5)

	result := num1.Divide(num2)

	fmt.Println(result)
	fmt.Println(num1)
	fmt.Println(num2)
}
