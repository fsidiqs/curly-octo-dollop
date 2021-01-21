package book

import "errors"

const (
	Category_Not_Found = "Category_Not_Found"
)

// Category has ID as the Category ID, Name as the Category Name, And Level according to the scenario given from the question
type Category struct {
	ID    int
	Name  string
	Level int
}

type Categories []Category

var MyCategories Categories

// creates a predefined category lists
func init() {
	MyCategories = Categories{
		Category{ID: 6, Name: "Applied Sciences (600)", Level: 1},
		Category{ID: 7, Name: "Arts (700)", Level: 2},
		Category{ID: 0, Name: "General (000)", Level: 3},
		Category{ID: 9, Name: "Geography, History (900)", Level: 4},
		Category{ID: 4, Name: "Language (400)", Level: 5},
		Category{ID: 8, Name: "Literature (800)", Level: 6},
		Category{ID: 1, Name: "Philosophy, Psychology (100)", Level: 7},
		Category{ID: 2, Name: "Religion (200)", Level: 8},
		Category{ID: 5, Name: "Science, Math (500)", Level: 9},
		Category{ID: 3, Name: "Social Sciences (300)", Level: 10},
	}
}

// GetCategoryByID method takes an ID as argument and returns the Category struct
func (categories Categories) GetCategoryByID(ID int) (Category, error) {
	for _, v := range categories {
		if v.ID == ID {
			return v, nil
		}
	}
	return Category{}, errors.New(Category_Not_Found)
}
