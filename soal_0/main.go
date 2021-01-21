package main

import (
	"bufio"
	"fmt"
	"os"
	"technical_test/soal_0/palindrome"
)

func main() {
	var input string
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Insert Two Numbers between 1 and 10^9")

	i := 0
	for scanner.Scan(); i <= 0; i++ {
		input = scanner.Text()
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}
	totalPalindrome, err := palindrome.CountPalindromePossibilities(input)
	fmt.Println(totalPalindrome)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("total possible palindrome: %d\n", totalPalindrome)
}
