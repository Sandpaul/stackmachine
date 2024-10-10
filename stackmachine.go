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
			case "DUP":
				duplicateTopmostValueOfStack(&stack)
			case "+":
				sumTopTwoIntegersOfStack(&stack)
			case "*":
				multiplyTopTwoIntegesOfStack(&stack)
			case "-":
				minusPenultimateIntegerFromTopIntegerOfStack(&stack)
			}
			
		}
	}

	topmostValueOfStack, err := getTopmostValueOfStack(&stack)
	if err != nil {
		return -1, err
	}
	
	return topmostValueOfStack, nil
}


func checkCommandAndConverterToInteger(command string) (int, error) {
	integer, err := strconv.Atoi(command)
	if err == nil {
		return integer, nil
	}

	return 0, err
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
	}
	
	return errors.New("integer out of bounds error")
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

func multiplyTopTwoIntegesOfStack(stack *[]int) {
	integer1, err := popIntegerFromStack(stack)
	integer2, err2 := popIntegerFromStack(stack)
	if err == nil && err2 == nil {
		product := integer1 * integer2
		checkIntegerAndPushToStack(product, stack)
	}
}

func minusPenultimateIntegerFromTopIntegerOfStack(stack *[]int) {
	integer1, err := popIntegerFromStack(stack)
	integer2, err2 := popIntegerFromStack(stack)
	if err == nil && err2 == nil {
		difference := integer1 - integer2
		checkIntegerAndPushToStack(difference, stack)
	}
}

func getTopmostValueOfStack(stack *[]int) (int, error){
	if len(*stack) == 0 {
		return -1, errors.New("")
	}

	topmostValueOfStack := (*stack)[len(*stack)-1]
	return topmostValueOfStack, nil
}

func duplicateTopmostValueOfStack(stack *[]int) {
	topmostValueOfStack, err := getTopmostValueOfStack(stack)
	if err != nil {
		return
	}
	
	checkIntegerAndPushToStack(topmostValueOfStack, stack)
}


// Commands:
// - `-` - pop the most recent two numbers, subtract the second from the most recent, push the result. Return error if result is below 0
// - CLEAR - empties the stack
// - SUM - pops all elements off the stack, adds them together, pushes result onto the stack. SUM on an empty stack returns an error.
// - Any other input is invalid - machine must stop and return an error


func main() {
	// main is unused - run using 
	// go test ./...
}