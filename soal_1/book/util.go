package book

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

// ParseBook takes a raw string as a paramter then returns the Books as the result
//
func ParseBook(booksStr string) (Books, error) {
	var result Books
	// clearing trailing spaces, and then turns the book into a slice of string
	bookSlices := strings.Split(strings.TrimSpace(booksStr), " ")
	if len(bookSlices) < 3 {
		return nil, errors.New(Err_Invalid_Book_Length)
	}
	for _, currentBook := range bookSlices {
		currentBookLen := len(currentBook)
		if currentBookLen < 4 || currentBookLen > 4 {
			return nil, errors.New(Err_Invalid_Book_Format)
		}
		categoryID, _ := strconv.Atoi(currentBook[:1])
		bookCategory, err := MyCategories.GetCategoryByID(categoryID)
		if err != nil {
			return nil, err
		}
		bookHeight, _ := strconv.Atoi(currentBook[2:])

		result = append(result, Book{
			Category: bookCategory,
			Name:     currentBook[1:2],
			Height:   bookHeight,
		})
	}
	return result, nil
}

// Parse AndGetSortedBooks takes an input as parameter for
// 1. beinng parsed, the parsed book then
// 2. sorted by Category and Height
// 3. filter books with maximum books 2 for each category
// the results is returned as string
func ParseAndGetSortedBooks(input string) (string, error) {
	// 1. parsing book
	books, err := ParseBook(input)
	if err != nil {
		return "", err
	}

	// 2. sort booksby category and height
	books.SortByCategoryThenHeight()
	// 3. filter maximum books for each category
	results, err := books.MaxSortedBooksEachCategory(2)
	if err != nil {
		return "", err
	}
	// convert the books into a string
	return results.toString(), nil
}

// toString method convert books into string which return a string
func (books Books) toString() string {
	// prepare slice of string that will be used to list according to format "[categoryID][bookName][bookHeight]"
	var s []string
	for _, v := range books {
		s = append(s, fmt.Sprintf("%d%s%d", v.Category.ID, v.Name, v.Height))
	}
	// join into a string with space as the delimiter
	return strings.Join(s, " ")
}

type Service struct {
}

func (Service) RunService(input string) (interface{}, error) {
	return ParseAndGetSortedBooks(input)
}
