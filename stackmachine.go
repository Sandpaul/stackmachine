package main

import (
	"errors"
	"strconv"
	"strings"
)


func StackMachine(commands string)(int, error) {
	stack := []int{}

	splitCommands := strings.Fields(commands)

	for _, command := range(splitCommands) {
		integer, err := checkCommandAndConverterToInteger(command)
		if err == nil {
			checkIntegerAndPushToStack(integer, &stack)
		} else {
			switch command {
			case "POP":
				popIntegerFromStack(&stack)
			case "+":
				sumTopTwoIntegersOfStack(&stack)
			}
			
		}
	}

	if len(stack) == 0 {
		return 0, errors.New("")
	}

	topmostValueOfStack := stack[len(stack)-1]

	return topmostValueOfStack, nil
}


func checkCommandAndConverterToInteger(command string) (int, error) {
	integer, err := strconv.Atoi(command)
	if err == nil {
		return integer, nil
	} else {
		return 0, err
	}
}

func integerInBounds(integer int) bool {
	lowerBound := 0
	upperBound := 50000
	return integer >= lowerBound && integer <= upperBound
}

func checkIntegerAndPushToStack(integer int, stack *[]int) error {
	if integerInBounds(integer) {
		*stack = append(*stack, integer)
		return nil
	} else {
		return errors.New("integer out of bounds error")
	}
}

func popIntegerFromStack(stack *[]int) (int, error) {
	if len(*stack) == 0 {
		return 0, errors.New("stack empty, noting to pop")
	}
	
	poppedInteger := (*stack)[len(*stack)-1]
	*stack = (*stack)[:len(*stack)-1]
	return poppedInteger, nil
}

func sumTopTwoIntegersOfStack(stack *[]int) {
	integer1, err1 := popIntegerFromStack(stack)
	integer2, err2 := popIntegerFromStack(stack)
	if err1 == nil && err2 == nil {
		sum := integer1 + integer2
		checkIntegerAndPushToStack(sum, stack)
	}
}


// Commands:
// - POP - removes most recently pushed integer from the stack
// - DUP - duplicate the last number pushed onto the stack (or keep the stack empty)
// - `+` - pop the the most recent two numbers, add them together and push the result. If an overflow occurs return an error
// - `-` - pop the most recent two numbers, subtract the second from the most recent, push the result. Return error if result is below 0
// - `*` - pop the top two elements from the stack, multiply them together, push the result onto the stack
// - CLEAR - empties the stack
// - SUM - pops all elements off the stack, adds them together, pushes result onto the stack. SUM on an empty stack returns an error.
// - Any other input is invalid - machine must stop and return an error

// After processing all the operators without errors, the machine returns the topmost value from the stack. 


func main() {
	// main is unused - run using 
	// go test ./...
}