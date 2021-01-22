package palindrome

import (
	"errors"
	"strconv"
	"strings"
)

var (
	MAXIMUM int = int(1e9)
	MINIMUM int = 1
)

type TwoNumbers struct {
	FirstNumber  int
	SecondNumber int
}

const (
	ErrInvalidInput        = "invalid_input"
	ErrInputOutOfRange     = "input_out_of_range"
	ErrSecondLessThanFirst = "secound_input_less_than_first_input"
)

// CountPalindromePossibilities will parse the input to a TwoNumbers Struct and then
// count the total palindrome possibilities, return it as integer
func CountPalindromePossibilities(input string) (int, error) {
	//parse the input
	parsedInput, err := ParseInput(input)
	if err != nil {
		return 0, err
	}

	// loop and check if the current number is palindrome, if so then increment
	// the countPalindrome variable
	countPalindrome := 0
	for currentNumber := parsedInput.FirstNumber; currentNumber <= parsedInput.SecondNumber; currentNumber++ {

		if IsPalindrome(currentNumber) {
			countPalindrome++
		}
	}

	// convert to string
	return countPalindrome, nil
}

// IsPalindrome is a function with parameter: number(the palindrome to be checked)
// returns: bool
// this function checks whether the given parameter is a palindrome or not
// returns true if the result is palindome, otherwise returns false
func IsPalindrome(number int) bool {
	// convert number into string
	firstNumberStr := strconv.Itoa(number)
	// make reversed string of the number in string type
	secondNumberStr := reverse(firstNumberStr)
	// if the source and the reversed are equals then it is a palindrome then return true
	if firstNumberStr == secondNumberStr {
		return true
	}
	// if not then return false
	return false
}

// reverse function is used for reversing a string
func reverse(str string) string {
	// prepare result as string
	var result string
	for _, v := range str {
		//  append the last result in the end of the current value
		result = string(v) + result
	}
	return result
}

func ParseInput(input string) (*TwoNumbers, error) {
	inputSlice := strings.Split(strings.TrimSpace(input), " ")
	inputLen := len(inputSlice)
	if inputLen > 2 || inputLen < 2 {
		return nil, errors.New(ErrInvalidInput)
	}

	firstNumber, err := strconv.Atoi(inputSlice[0])
	secondNumber, err := strconv.Atoi(inputSlice[1])

	if err != nil {
		return nil, errors.New(ErrInvalidInput)
	}

	if firstNumber < MINIMUM || secondNumber > MAXIMUM {
		return nil, errors.New(ErrInputOutOfRange)
	} else if secondNumber <= firstNumber {
		return nil, errors.New(ErrSecondLessThanFirst)
	}

	return &TwoNumbers{firstNumber, secondNumber}, nil
}

type Service struct {
}

func (Service) RunService(input string) (interface{}, error) {
	return CountPalindromePossibilities(input)
}
