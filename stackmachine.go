package main

import (
	"errors"
	"strconv"
	"strings"
)

func StackMachine(commands string) (int, error) {
	stack := []int{}

	splitCommands := strings.Fields(commands)

	for _, command := range splitCommands {
		var err error

		switch command {
		case "POP":
			_, err = popIntegerFromStack(&stack)
		case "DUP":
			duplicateTopmostIntegerOfStack(&stack)
		case "+":
			err = sumTopTwoIntegersOfStack(&stack)
		case "*":
			err = multiplyTopTwoIntegersOfStack(&stack)
		case "-":
			err = minusPenultimateIntegerFromTopIntegerOfStack(&stack)
		case "CLEAR":
			clearStack(&stack)
		case "SUM":
			err = sumAllIntegersOnStack(&stack)
		default:
			err = pushIntegerToStack(command, &stack)
		}

		if err != nil {
			return 0, err
		}
	}

	topmostIntegerOfStack, err := getTopmostIntegerOfStack(&stack)
	if err != nil {
		return 0, err
	}

	return topmostIntegerOfStack, nil
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

	return errors.New("")
}

func pushIntegerToStack(command string, stack *[]int) error {
	integer, err := checkCommandAndConverterToInteger(command)
	if err != nil {
		return errors.New("")
	}
	err = checkIntegerAndPushToStack(integer, stack)
	if err != nil {
		return errors.New("")
	}
	return nil
}


func popIntegerFromStack(stack *[]int) (int, error) {
	if len(*stack) == 0 {
		return 0, errors.New("")
	}

	poppedInteger := (*stack)[len(*stack)-1]
	*stack = (*stack)[:len(*stack)-1]
	return poppedInteger, nil
}

func sumTopTwoIntegersOfStack(stack *[]int) error {
	integer1, err1 := popIntegerFromStack(stack)
	integer2, err2 := popIntegerFromStack(stack)
	if err1 == nil && err2 == nil {
		sum := integer1 + integer2
		checkIntegerAndPushToStack(sum, stack)
		return nil
	}
	return errors.New("")
}

func multiplyTopTwoIntegersOfStack(stack *[]int) error {
	integer1, err := popIntegerFromStack(stack)
	integer2, err2 := popIntegerFromStack(stack)
	if err == nil && err2 == nil {
		product := integer1 * integer2
		checkIntegerAndPushToStack(product, stack)
		return nil
	}
	return errors.New("")
}

func minusPenultimateIntegerFromTopIntegerOfStack(stack *[]int) error {
	integer1, err := popIntegerFromStack(stack)
	integer2, err2 := popIntegerFromStack(stack)
	if err == nil && err2 == nil {
		difference := integer1 - integer2
		checkIntegerAndPushToStack(difference, stack)
		return nil
	}
	return errors.New("")
}

func getTopmostIntegerOfStack(stack *[]int) (int, error) {
	if len(*stack) == 0 {
		return -1, errors.New("")
	}

	topmostValueOfStack := (*stack)[len(*stack)-1]
	return topmostValueOfStack, nil
}

func duplicateTopmostIntegerOfStack(stack *[]int) {
	topmostValueOfStack, err := getTopmostIntegerOfStack(stack)
	if err != nil {
		return
	}

	checkIntegerAndPushToStack(topmostValueOfStack, stack)
}

func clearStack(stack *[]int) {
	*stack = (*stack)[:0]
}

func sumAllIntegersOnStack(stack *[]int) error {
	if len(*stack) == 0 {
		return errors.New("")
	}
	total := 0
	for _, integer := range *stack {
		total += integer
	}
	clearStack(stack)
	checkIntegerAndPushToStack(total, stack)
	return nil
}

func main() {
	// main is unused - run using
	// go test ./...
}
