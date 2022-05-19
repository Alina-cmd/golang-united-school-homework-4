package string_sum

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

//use these errors as appropriate, wrapping them with fmt.Errorf function
var (
	// Use when the input is empty, and input is considered empty if the string contains only whitespace
	errorEmptyInput = errors.New("input is empty")
	// Use when the expression has number of operands not equal to two
	errorNotTwoOperands = errors.New("expecting two operands, but received more or less")
)

// Implement a function that computes the sum of two int numbers written as a string
// For example, having an input string "3+5", it should return output string "8" and nil error
// Consider cases, when operands are negative ("-3+5" or "-3-5") and when input string contains whitespace (" 3 + 5 ")
//
//For the cases, when the input expression is not valid(contains characters, that are not numbers, +, - or whitespace)
// the function should return an empty string and an appropriate error from strconv package wrapped into your own error
// with fmt.Errorf function
//
// Use the errors defined above as described, again wrapping into fmt.Errorf

func StringSum(input string) (output string, err error) {

	input = deleteWhitespaces(input)

	input, err = checkIsEmpty(input)
	if err != nil {
		return input, fmt.Errorf(" Empty Input in StringSum(): %w", err)
	}

	input, err = checkContainExtraSymbol(input)
	if err != nil {
		return input, fmt.Errorf("Contain extra symbol: %w", err)
	}

	amount := getAmountOfNumbers(input)
	if amount != 2 {
		return "", fmt.Errorf(" Error in StringSum: %w", errorNotTwoOperands)
	}

	first, second := getFirstAndSecond(input)

	output = strconv.Itoa(first + second)

	return output, nil
}

func deleteWhitespaces(input string) string {
	return strings.ReplaceAll(input, " ", "")
}

func isContain(what rune, where []rune) (exist bool) {
	for _, v := range where {
		if v == what {
			return true
		}
	}
	return
}

func checkIsEmpty(input string) (string, error) {
	if input == "" {
		return "", errorEmptyInput
	}

	return input, nil
}

func checkContainExtraSymbol(input string) (string, error) {
	symbols := []rune{45, 43, 49, 50, 51, 52, 53, 54, 55, 56, 57, 32} // { '-', '+', '1', '2', '3', '4', '5', '6', '7', '8', '9', ' '}
	inputInRunes := []rune(input)
	for _, v := range inputInRunes {
		if isContain(v, symbols) == false {
			_, errAtoi := strconv.Atoi(string(v)) // get error from strconv package
			return "", errAtoi
		}
	}

	return input, nil
}

func getAmountOfNumbers(input string) (amount int) {
	input = strings.ReplaceAll(input, "+", "/+")
	input = strings.ReplaceAll(input, "-", "/-")
	count := strings.Split(input, "/")
	if count[0] == "" { // case "/-3/-5"
		amount-- // amount == -1
	}
	amount += len(count)

	return
}

func getFirstAndSecond(input string) (first, second int) {
	input = strings.ReplaceAll(input, "+", "/+") // "3+5" -> "3/+5"
	input = strings.ReplaceAll(input, "-", "/-")
	count := strings.Split(input, "/") // count = { "3", "+5"}
	start := 0
	if count[0] == "" {
		start = 1
	}
	first, _ = strconv.Atoi(count[start])
	second, _ = strconv.Atoi(count[start+1])

	return
}
