package main

import (
	"bufio"
	"fmt"
	"os"
	"technical_test/soal_1/book"
)

func main() {
	var input string
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Insert Books")

	i := 0
	for scanner.Scan(); i <= 0; i++ {
		input = scanner.Text()
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}
	output, err := book.ParseAndGetSortedBooks(input)
	if err != nil {
		fmt.Printf("error %+v", err)
		return
	}
	fmt.Println(output)

}
