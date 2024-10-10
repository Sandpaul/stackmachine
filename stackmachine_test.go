package main

import (
	"errors"
	"testing"

)

func TestStartsWithEmptyStack(t *testing.T) {
	_, err := StackMachine("")

	if err == nil {
		t.Error("expected error due to no results")
	}
}

func TestReturnsTopValueFromStack(t *testing.T) {
	actual, err := StackMachine("3 2 1")
	if err != nil {
		t.Errorf("expected no error but got: %v", err)
	}

	expected := 1

	if actual != expected {
		t.Errorf("expected: %v but got: %v", expected, actual)
	}
}

func TestNegativeNumbersNotAddedToStack(t *testing.T) {
	_, err := StackMachine("-1")
	if err == nil {
		t.Error("expected error due to negative number")
	}
}

func TestNumbersOverFiftyThousandNotAddedToStack(t *testing.T) {
	_, err := StackMachine("50001")
	if err == nil {
		t.Error("expected number out of bounds error")
	}
}

func TestPopRemovesMostRecentIntegerFromStack(t *testing.T) {
	actual, err := StackMachine("1 2 3 POP")
	if err != nil {
		t.Errorf("expected no error but got: %v", err)
	}

	expected := 2

	if actual != expected {
		t.Errorf("expected: %v but got: %v", expected, actual)
	}
}

func TestPlusCommandAddsTwoNumbers(t *testing.T) {
	actual, err := StackMachine("1 2 +")
	if err != nil {
		t.Errorf("expected no error but got: %v", err)
	}

	expected := 3

	if actual != expected {
		t.Errorf("expected: %v but got: %v", expected, actual)
	}
}

func TestDupCommandDuplicatesTopNumberOfStack(t *testing.T) {
	actual, err := StackMachine("3 DUP +")
	if err != nil {
		t.Errorf("expected no error but got: %v", err)
	}

	expected := 6

	if actual != expected {
		t.Errorf("expected: %v but got: %v", expected, actual)
	}
}


func TestDupCommandDoesNothingIfStackIsEmpty(t *testing.T) {
	actual, err := StackMachine("DUP 99")
	if err != nil {
		t.Errorf("expected no error but got: %v", err)
	}

	expected := 99

	if actual != expected {
		t.Errorf("expected: %v but got: %v", expected, actual)
	}
}

func TestMultiplyCommandMultipliesTopTwoNumbersOnStack(t *testing.T) {
	actual, err := StackMachine("9 9 *")
	if err != nil {
		t.Errorf("expected no error but got: %v", err)
	}

	expected := 81

	if actual != expected {
		t.Errorf("expected: %v but got: %v", expected, actual)
	}
}

func TestMinusCommandDeductsPenultimateNumberOnStackFromTopNumber(t *testing.T) {
	actual, err := StackMachine("2 5 -")
	if err != nil {
		t.Errorf("expected no error but got: %v", err)
	}

	expected := 3

	if actual != expected {
		t.Errorf("expected: %v but got: %v", expected, actual)
	}
}

/*
  All these tests must pass for completion
*/
func TestAcceptanceTests(t *testing.T) {
	tests := []struct {
		name string
		commands string
		expected int
		expectedErr error		
	}{
		{name:"empty error", commands:"", expected:0, expectedErr: errors.New("")},
		{name:"add overflow", commands:"50000 DUP +", expected: 0, expectedErr: errors.New("") },
		{name:"too few add", commands:"99 +", expected: 0, expectedErr: errors.New("") },
		{name:"too few minus", commands:"99 -", expected: 0, expectedErr: errors.New("") },
		{name:"too few multiply", commands:"99 *", expected: 0, expectedErr: errors.New("") },
		{name:"empty stack", commands:"99 CLEAR", expected: 0, expectedErr: errors.New("") },
		{name:"sum single value", commands:"99 SUM", expected: 99, expectedErr: nil },
		{name:"sum empty", commands:"SUM", expected: 0, expectedErr: errors.New("") },
		{name:"normal +*", commands:"5 6 + 2 *", expected: 22, expectedErr: nil },
		{name:"clear too few", commands:"1 2 3 4 + CLEAR 12 +", expected: 0, expectedErr: errors.New("") },
		{name:"normal after clear", commands:"1 CLEAR 2 3 +", expected: 5, expectedErr: nil },
		{name:"single integer", commands:"9876", expected: 9876, expectedErr: nil },
		{name:"invalid command", commands:"DOGBANANA", expected: 0, expectedErr: errors.New("") },
		{name:"normal +-*", commands:"5 9 DUP + + 43 - 3 *", expected: 60, expectedErr: nil },
		{name:"minus", commands:"2 5 -", expected: 3, expectedErr: nil },
		{name:"underflow minus", commands:"5 2 -", expected: 0, expectedErr: errors.New("") },
		{name:"at overflow limit", commands:"25000 DUP +", expected: 50000, expectedErr: nil },
		{name:"at overflow limit single value", commands:"50000 0 +", expected: 50000, expectedErr: nil },
		{name:"overflow plus", commands:"50000 1 +", expected: 0, expectedErr: errors.New("") },
		{name:"overflow single value", commands:"50001", expected: 0, expectedErr: errors.New("") },
		{name:"times zero at overflow limit", commands:"50000 0 *", expected: 0, expectedErr: nil },
		{name:"too few at first", commands:"1 2 3 4 5 + + + + * 999", expected: 0, expectedErr: errors.New("") },
		{name:"normal simple", commands:"1 2 - 99 +", expected: 100, expectedErr: nil },
		{name:"at overflow minus to zero", commands:"50000 50000 -", expected: 0, expectedErr: nil },
		{name:"clear empties stack", commands:"CLEAR", expected: 0, expectedErr: errors.New("") },
		{name:"normal sum", commands:"3 4 3 5 5 1 1 1 SUM", expected: 23, expectedErr: nil },
		{name:"sum after clear stack", commands:"3 4 3 5 CLEAR 5 1 1 1 SUM", expected: 8, expectedErr: nil },
		{name:"sum then too few", commands:"3 4 3 5 5 1 1 1 SUM -", expected: 0, expectedErr: errors.New("") },
		{name:"fibonacci", commands:"1 2 3 4 5 * * * *", expected: 120, expectedErr: nil },
	}

	for _, test := range tests {
			
		got, err := StackMachine(test.commands)

		if (test.expectedErr != nil) {
			if err == nil {
				t.Errorf("%s (%s) Expected error, but received nil", test.name, test.commands)
			} else if err.Error() != test.expectedErr.Error()  {
				t.Errorf("%s (%s) got error %v, want %v", test.name, test.commands, err, test.expectedErr)
			}
		}  else if got != test.expected {
			t.Errorf("%s (%s) got %v, want %v", test.name, test.commands, got, test.expected)
		}
	}
}