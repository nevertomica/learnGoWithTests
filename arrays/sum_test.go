package arrays

import (
	"slices"
	"testing"
)

func TestSumAll(t *testing.T) {
	sliceA := []int{1, 2}
	sliceB := []int{0, 9}
	want := []int{3, 9}
	got := SumAll(sliceA, sliceB)

	if !slices.Equal(want, got) {
		t.Errorf("got %v want %v, given slices %v, %v", got, want, sliceA, sliceB)
	}

}

func TestSum(t *testing.T) {
	numbers := []int{1, 2, 3, 4, 5}
	want := 15
	got := Sum(numbers)

	if want != got {
		t.Errorf("got %d want %d, given %v", got, want, numbers)
	}
}

func TestSumAllTails(t *testing.T) {

	checkSum := func(t testing.TB, got, want []int, input ...[]int) {
		t.Helper()
		if !slices.Equal(want, got) {
			t.Errorf("got %d want %d, given %v", got, want, input)
		}

	}

	t.Run(" len = 1 subslice ", func(t *testing.T) {
		numbers := []int{3}
		want := []int{0}
		got := SumAllTails(numbers)
		checkSum(t, got, want, numbers)
	})

	t.Run("len = 0 slice", func(t *testing.T) {
		numbers := []int{}
		want := []int{0}
		got := SumAllTails(numbers)
		checkSum(t, got, want, numbers)
	})

	t.Run(" len > 1 subslice ", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5}
		want := []int{14}
		got := SumAllTails(numbers)
		checkSum(t, got, want, numbers)
	})

	t.Run(" mixed subslices ", func(t *testing.T) {
		input := [][]int{[]int{1, 2, 3, 4, 5}, []int{2, 3}, []int{100}, []int{}}
		want := []int{14, 3, 0, 0}
		got := SumAllTails(input...)
		checkSum(t, got, want, input...)
	})
}
