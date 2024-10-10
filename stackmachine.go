package main

import (
	"errors"
	"strconv"
	"strings"
)

func StackMachine(commands string) (int, error) {
	stack := []int{}

	result, err := executeCommands(commands, &stack)
	if err != nil {
		return 0, err
	}

	return result, nil
}


func executeCommands(commands string, stack *[]int) (int, error) {
	splitCommands := strings.Fields(commands)

	for _, command := range splitCommands {
		err := executeCommand(command, stack)
		if err != nil {
			return 0, err
		}
	}

	topmostIntegerOfStack, err := popFromStack(stack)
	if err != nil {
		return 0, err
	}

	return topmostIntegerOfStack, nil
}


func executeCommand(command string, stack *[]int) error {
	var err error

	switch command {
	case "POP":
		_, err = popFromStack(stack)
	case "DUP":
		duplicateTop(stack)
	case "SUM":
		err = sumAll(stack)
	case "CLEAR":
		clearStack(stack)
	case "+":
		err = sumTopTwo(stack)
	case "*":
		err = multiplyTopTwo(stack)
	case "-":
		err = subtractSecondFromTop(stack)
	default:
		err = pushToStack(command, stack)
	}

	if err != nil {
		return err
	}
	return nil
}


func parseIntegerCommand(command string) (int, error) {
	integer, err := strconv.Atoi(command)
	if err == nil {
		return integer, nil
	}

	return 0, err
}

func isIntegerInBounds(integer int) bool {
	lowerBound := 0
	upperBound := 50000
	return integer >= lowerBound && integer <= upperBound
}

func pushIfValid(integer int, stack *[]int) error {
	if isIntegerInBounds(integer) {
		*stack = append(*stack, integer)
		return nil
	}

	return errors.New("")
}

func pushToStack(command string, stack *[]int) error {
	integer, err := parseIntegerCommand(command)
	if err != nil {
		return errors.New("")
	}
	err = pushIfValid(integer, stack)
	if err != nil {
		return errors.New("")
	}
	return nil
}


func popFromStack(stack *[]int) (int, error) {
	if len(*stack) == 0 {
		return 0, errors.New("")
	}

	poppedInteger := (*stack)[len(*stack)-1]
	*stack = (*stack)[:len(*stack)-1]
	return poppedInteger, nil
}

func sumTopTwo(stack *[]int) error {
	integer1, err1 := popFromStack(stack)
	integer2, err2 := popFromStack(stack)
	if err1 == nil && err2 == nil {
		sum := integer1 + integer2
		pushIfValid(sum, stack)
		return nil
	}
	return errors.New("")
}

func multiplyTopTwo(stack *[]int) error {
	integer1, err := popFromStack(stack)
	integer2, err2 := popFromStack(stack)
	if err == nil && err2 == nil {
		product := integer1 * integer2
		pushIfValid(product, stack)
		return nil
	}
	return errors.New("")
}

func subtractSecondFromTop(stack *[]int) error {
	integer1, err := popFromStack(stack)
	integer2, err2 := popFromStack(stack)
	if err == nil && err2 == nil {
		difference := integer1 - integer2
		pushIfValid(difference, stack)
		return nil
	}
	return errors.New("")
}

func getTop(stack *[]int) (int, error) {
	if len(*stack) == 0 {
		return -1, errors.New("")
	}

	topmostValueOfStack := (*stack)[len(*stack)-1]
	return topmostValueOfStack, nil
}

func duplicateTop(stack *[]int) {
	topmostValueOfStack, err := getTop(stack)
	if err != nil {
		return
	}

	pushIfValid(topmostValueOfStack, stack)
}

func clearStack(stack *[]int) {
	*stack = (*stack)[:0]
}

func sumAll(stack *[]int) error {
	if len(*stack) == 0 {
		return errors.New("")
	}
	total := 0
	for _, integer := range *stack {
		total += integer
	}
	clearStack(stack)
	pushIfValid(total, stack)
	return nil
}

func main() {
	// main is unused - run using
	// go test ./...
}
