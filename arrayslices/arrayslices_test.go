package arrayslices

import (
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {
	//numbers := [5]int{1, 2, 3, 4, 5}
	//got := Sum(numbers)
	//want := 15
	//
	//if got != want {
	//	t.Errorf("got %d want %d given %v", got, want, numbers)
	//}

	t.Run("collection of 5 numbers", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5}
		got := Sum(numbers)
		want := 15

		if got != want {
			t.Errorf("got %d want %d given %v", got, want, numbers)
		}
	})

	t.Run("collection of any size", func(t *testing.T) {
		numbers := []int{4, 10, 23}
		got := Sum(numbers)
		want := 37

		if got != want {
			t.Errorf("got %d want %d given %v", got, want, numbers)
		}
	})
}

//func TestSumAll(t *testing.T) {
//	got := SumAll([]int{1, 3}, []int{2, 4})
//	want := []int{4, 6}
//
//	if !reflect.DeepEqual(got, want) {
//		t.Errorf("got %v want %v", got, want)
//	}
//}

func TestSumAllTails(t *testing.T) {
	checkSums := func(t *testing.T, got, want []int) {
		t.Helper()
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	}
	t.Run("make the sums of tails of", func(t *testing.T) {
		got := SumAllTails([]int{1, 3}, []int{2, 4})
		want := []int{3, 4}

		checkSums(t, got, want)
	})

	t.Run("safely sum empty slices", func(t *testing.T) {
		got := SumAllTails([]int{}, []int{6, 10})
		want := []int{0, 10}

		checkSums(t, got, want)
	})
}
