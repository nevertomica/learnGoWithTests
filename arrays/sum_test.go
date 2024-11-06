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
