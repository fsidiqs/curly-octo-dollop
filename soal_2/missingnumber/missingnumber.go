package missingnumber

import (
	"errors"
	"strconv"
)

const (
	MINIMUMNUMBEROFSEQUENCES = 3
)

// here is the simple explanation
// actual[97,99,100,101,102]
// prediction[97,98,100,101,102]
// I made the solution by comparing the actual input, with the predition
// 1. currentActual 97, currentPrediction 97
// 2. currentActual 99, currentPrediction 98, does not match, then its the missing number

func FindMissingNumber(input string) (int, error) {
	inputLen := len(input)
	if inputLen < 3 || inputLen > 1000 {
		return 0, errors.New("invalid_input")
	}
	var missingNumber int
	// inputLen, is total length of the input string

	// let's say we have 9799100101102, the length is =13
	// the possible number of digits is = inputLen / MINIMUMNUMBEROFSEQUENCES => 13 /3 = 4
	// so we have at most 4 items per chunks, and the possibilities are:
	//* 1 item per chunk. 9,
	//* 2 items per chunk. 97,
	//* 3 items per chunk. 979,
	//* 4 items per chunk. 9799
	maxItemPerChunk := inputLen / MINIMUMNUMBEROFSEQUENCES

	// we traverse the input item according to the item per chunk, starting from 1 item in each chunk
	for itemPerChunk := 1; itemPerChunk <= maxItemPerChunk; itemPerChunk++ {
		// let's consider itemPerChunk is now = 2, since its the correct chunk
		// initialNumber is the firstNumber of the chunks for the example above, it is now = 97
		initialNumber, _ := strconv.Atoi(input[0:itemPerChunk])

		// secondNumberPredict is the next sequence after the initial number (97) then it would be 98
		secondNumberPredict := initialNumber + 1
		// we count the length of the next prediction number
		secondNumberPredictLen := len(strconv.Itoa(secondNumberPredict))

		// get the actual input by extracting the input by their their index: in this example would be 99
		secondNumberActual, _ := strconv.Atoi(input[itemPerChunk : itemPerChunk+secondNumberPredictLen])
		if secondNumberPredict+1 == secondNumberActual*10 {
			secondNumberActual *= 10
		}

		// this logic is intended to decided whether the items for each chunk is considered a sequential number.
		// -1, is considered a valid since for example "1,2,3,4" would give -1
		// -2, is also possibly a sequence number, since its might be contain the missing number for example "97,99,100" 97-99 = -2
		if initialNumber-secondNumberActual == -1 || initialNumber-secondNumberActual == -2 {

			// here currentNumberPredict be initially assigned to the initialNumber.
			currentNumberPredict := initialNumber
			//	inputCusrsorPosition is the position of input item being scanned
			// example "9799100101102", if the currentNumberPredict is 99, so the inputCursorPosition should be at "97[99]100101102"
			inputCusrsorPosition := 0
			// here we are trying to compare the predition of currentNumber with the actual input
			for {
				// convert the currentNumberPredict to a string, so currentText = currentNumberPredict.tostring()
				currentText := strconv.Itoa(currentNumberPredict)
				// count the numberof digiit of the currentNumber
				currentTextLen := len(currentText)
				// getActualInput, the current actual input

				endIndex := inputCusrsorPosition + currentTextLen
				// if endIndex exceeds the actual Input length then break the loop
				if endIndex > inputLen {

					break
				}
				getActualInput := input[inputCusrsorPosition:endIndex]
				// moving the input cursor position for next scan
				inputCusrsorPosition += currentTextLen
				// convert input to integer
				getActualInputInt, _ := strconv.Atoi(getActualInput)
				//* if the currentNumberPredict is not equal to the actualInput then it is possibily the missing number

				if currentNumberPredict != getActualInputInt {

					// if the actual input is 97,
					// currentNumberPredic+1(the next sequence number) is 99, and missing number is never found,
					// then the currentNUmberPredict is the missingNumber, then return
					// if not then the input/chunk is considered as invalid input,
					if getActualInputInt == currentNumberPredict+1 && missingNumber == 0 {

						missingNumber = currentNumberPredict
						// fmt.Println("test")
						// fmt.Println(missingNumber)
						return missingNumber, nil
					} else if currentNumberPredict+1 == getActualInputInt*10 && missingNumber == 0 {
						// case for current 9, next actual need more digit
						missingNumber = currentNumberPredict
						return missingNumber, nil
					}
					break
				}

				currentNumberPredict++
			}
		}
	}
	return 0, errors.New("invalid_input")
}

type Service struct {
}

func (Service) RunService(input string) (interface{}, error) {
	return FindMissingNumber(input)
}
