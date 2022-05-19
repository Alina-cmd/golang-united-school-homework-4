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

	// delete whitespaces
	input = strings.ReplaceAll(input, " ", "")

	// is empty
	if input == "" {
		return "", fmt.Errorf(" Empty Input in StringSum(): %w", errorEmptyInput)
	}

	// contain extra symbols
	symbols := []rune{45, 43, 49, 50, 51, 52, 53, 54, 55, 56, 57, 32}
	inputRunes := []rune(input)
	for _, v := range inputRunes {
		if find(v, symbols) == -1 {
			_, errAtoi := strconv.Atoi(string(v))
			return "", fmt.Errorf("Contain extra symbol: %w", errAtoi)
		}
	}

	// count of numbers
	amount := 0
	input = strings.ReplaceAll(input, "+", "/+")
	input = strings.ReplaceAll(input, "-", "/-")
	count := strings.Split(input, "/")
	start := 0
	if count[0] == "" {
		amount--
		start = 1
	}
	amount += len(count)
	if amount != 2 {
		return "", fmt.Errorf(" Error in StringSum: %w", errorNotTwoOperands)
	}
	first, _ := strconv.Atoi(count[start])
	second, _ := strconv.Atoi(count[start+1])
	output = strconv.Itoa(first + second)

	return output, nil
}

func find(what rune, where []rune) (idx int) {
	idx = 1
	for _, v := range where {
		if v == what {
			return
		}
	}
	return -1
}
