package arrsli

import (
	"slices"
	"testing"
)

func TestSum(t *testing.T) {
	t.Run("collection of any lenght", func(t *testing.T) {
		numbers := []int{10, 2, 3}
		got := Sum(numbers)
		want := 15

		if got != want {
			t.Errorf("got %d want %d,given %v", got, want, numbers)
		}
	})
}

func TestSumAll(t *testing.T) {
	got := SumAll([]int{1, 2}, []int{0, 9})
	want := []int{3, 9}

	if !slices.Equal(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}

func TestSumAllTails(t *testing.T) {

	checkFunc := func(t testing.TB, got, want []int) {
		t.Helper()
		if !slices.Equal(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	}

	t.Run("sums of slices", func(t *testing.T) {
		got := SumAllTails([]int{1, 2}, []int{0, 9})
		want := []int{2, 9}

		checkFunc(t, got, want)
	})

	t.Run("using an empty slice", func(t *testing.T) {
		got := SumAllTails([]int{}, []int{3, 4, 5})
		want := []int{0, 9}

		checkFunc(t, got, want)
	})
}
