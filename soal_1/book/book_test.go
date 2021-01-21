package book

import "testing"

func TestInvalidInput(t *testing.T) {
	t.Run("input empty string", func(t *testing.T) {
		actual, err := ParseAndGetSortedBooks("")
		if actual != "" && err == nil {
			t.Errorf("expected an error")
		}

	})

}

func TestBookOnly2(t *testing.T) {
	t.Run("input 2 books ", func(t *testing.T) {
		actual, err := ParseAndGetSortedBooks("3A13 AX19")
		if actual != "" && err == nil {
			t.Errorf("expected an error")
		}
	})

}

func TestInvalidBookNaming(t *testing.T) {
	t.Run("input invalid book ", func(t *testing.T) {
		actual, err := ParseAndGetSortedBooks("3A13 aaa bbb")
		if actual != "" && err == nil {
			t.Errorf("expected an error")
		}
	})

}
func TestSortBookValid(t *testing.T) {
	t.Run("input '3A13 5X19 9Y20 2C18 1N20 3N20 1M21 1F14 9A21 3N21 0E13 5G14 8A23 9E22 3N14'", func(t *testing.T) {
		result, _ := ParseAndGetSortedBooks("3A13 5X19 9Y20 2C18 1N20 3N20 1M21 1F14 9A21 3N21 0E13 5G14 8A23 9E22 3N14")
		expect := "0E13 9E22 9A21 9Y20 8A23 1M21 1N20 1F14 2C18 5X19 5G14 3N21 3N20 3A13"

		if result != expect {
			t.Errorf("got %v want %v", result, expect)
		}
	})
}
