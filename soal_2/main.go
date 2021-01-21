package main

import (
	"bufio"
	"fmt"
	"os"
	"technical_test/soal_2/missingnumber"
)

func main() {
	var input string
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Find missing number:")

	i := 0
	for scanner.Scan(); i <= 0; i++ {
		input = scanner.Text()
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}

	missingNumber, _ := missingnumber.FindMissingNumber(input)
	// fmt.Println(missingNumber)
	if missingNumber != 0 {

		fmt.Printf("The missing number is: %v \n", missingNumber)
		return
	}
	fmt.Println("invalid_input")
}
