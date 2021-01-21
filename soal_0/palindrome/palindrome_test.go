package palindrome

import (
	"testing"
)

func TestEmptyString(t *testing.T) {

	// if the input is empty, returns an invalid_input message
	t.Run("input an empty sring", func(t *testing.T) {
		actual, err := CountPalindromePossibilities("")

		if actual != 0 && err == nil {
			t.Errorf("expected an error")
		}
	})
}

func TestInvalidInput(t *testing.T) {

	t.Run("input is not valid", func(t *testing.T) {
		actual, err := CountPalindromePossibilities("1 a")
		if actual != 0 && err == nil {
			t.Errorf("expected an error")
		}
	})
}

// valid input
func TestValidInput(t *testing.T) {

	t.Run("input 1 10", func(t *testing.T) {
		result, err := CountPalindromePossibilities("1 10")
		expect := 9
		if err != nil {
			t.Error(err)
			return
		}
		if result != expect {
			t.Errorf("result %v expect %v", result, expect)
		}
	})

	t.Run("input 99 100", func(t *testing.T) {
		result, err := CountPalindromePossibilities("99 100")
		expect := 1
		if err != nil {
			t.Error(err)
			return
		}
		if result != expect {
			t.Errorf("result %v expect %v", result, expect)
		}
	})

	t.Run("input 21 31", func(t *testing.T) {
		result, err := CountPalindromePossibilities("21 31")
		expect := 1
		if err != nil {
			t.Error(err)
			return
		}
		if result != expect {
			t.Errorf("result %v expect %v", result, expect)
		}
	})
}
