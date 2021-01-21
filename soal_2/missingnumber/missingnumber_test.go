package missingnumber

import (
	"testing"
)

func TestEmptyString(t *testing.T) {
	t.Run("input empty string", func(t *testing.T) {
		actual, err := FindMissingNumber("")
		if actual != 0 && err == nil {
			t.Errorf("expected an error")
		}
	})

}

func TestInvalidInput(t *testing.T) {
	t.Run("input '12'", func(t *testing.T) {
		actual, err := FindMissingNumber("12")
		if actual != 0 && err == nil {
			t.Errorf("expected an error")
		}
	})

}
func TestFindMissingNumberResult(t *testing.T) {

	t.Run("input '23242526272830'", func(t *testing.T) {
		result, _ := FindMissingNumber("23242526272830")
		expect := 29

		if result != expect {
			t.Errorf("got %v want %v", result, expect)
		}
	})

	t.Run("when argument is '101102103104105106107108109111112113'", func(t *testing.T) {
		result, _ := FindMissingNumber("101102103104105106107108109111112113")
		expect := 110

		if result != expect {
			t.Errorf("got %v want %v", result, expect)
		}
	})

	t.Run("when argument is '12346789'", func(t *testing.T) {
		result, _ := FindMissingNumber("12346789")
		expect := 5

		if result != expect {
			t.Errorf("got %v want %v", result, expect)
		}
	})

	t.Run("when argument is '79101112'", func(t *testing.T) {
		result, _ := FindMissingNumber("79101112")
		expect := 8

		if result != expect {
			t.Errorf("got %v want %v", result, expect)
		}
	})

	t.Run("when argument is '7891012'", func(t *testing.T) {
		result, _ := FindMissingNumber("7891012")
		expect := 11

		if result != expect {
			t.Errorf("got %v want %v", result, expect)
		}
	})

	t.Run("when argument is '9799100101102'", func(t *testing.T) {
		result, _ := FindMissingNumber("9799100101102")
		expect := 98

		if result != expect {
			t.Errorf("got %v want %v", result, expect)
		}
	})

	t.Run("when argument is '100001100002100003100004100006'", func(t *testing.T) {
		result, _ := FindMissingNumber("100001100002100003100004100006")
		expect := 100005

		if result != expect {
			t.Errorf("got %v want %v", result, expect)
		}
	})
}
