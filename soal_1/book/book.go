package book

import (
	"errors"
	"sort"
)

const (
	Err_Invalid_Max_Book    = "Invalid_Max_Book"
	Err_Invalid_Book_Format = "Invalid_Book_Format"
	Err_Invalid_Book_Length = "Invalid_Book_Length"
)

type Book struct {
	Category Category
	Name     string
	Height   int
}

type Books []Book

// SortByCategoryThenHeight method will sort the books based on the following priority
// 1. Book Category
// 2. Book Height
func (books Books) SortByCategoryThenHeight() {
	sort.Slice(books, func(i, j int) bool {
		if books[i].Category.Level == books[j].Category.Level {
			return books[i].Height > books[j].Height
		}
		return books[i].Category.Level < books[j].Category.Level
	})
}

// MaxSortedBooksEachCateory (books must have already been sorted) takes an integer as the argument it measure the  total masimum of books allowed for each category
func (books Books) MaxSortedBooksEachCategory(max int) (Books, error) {
	// maximum value must not less or equeal than 0
	if max <= 0 {
		return nil, errors.New(Err_Invalid_Max_Book)
	}
	// result to be returned
	var results Books

	var prevBook struct {
		name  string
		count int
	}

	for _, book := range books {
		// checking whether the previous book name is different than current book, if they are different
		// the previous book will be replaced by the current book, and start over the counter
		if prevBook.name != book.Name {
			prevBook.name = book.Name
			prevBook.count = 1
		}
		// checking each book name, if it is occuring more than 2 times, it wont be appended
		if prevBook.count <= 2 {
			results = append(results, book)
		}
		prevBook.count++
	}
	return results, nil
}
