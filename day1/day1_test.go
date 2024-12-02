package day1

import (
	"reflect"
	"testing"
)

func TestDay1(t *testing.T) {
	pairs := []string{
		"3 4",
		"4 3",
		"2 5",
		"1 3",
		"3 9",
		"3 3",
	}

	t.Run("processPairs - test it returns two ordered lists", func(t *testing.T) {
		wantLeft := []int{1, 2, 3, 3, 3, 4}
		wantRight := []int{3, 3, 3, 4, 5, 9}
		wantSimilarities := map[int]int{3: 3, 4: 1, 5: 1, 9: 1}

		got, err := processPairs(pairs)

		if err != nil {
			t.Fatal("got unexpected error")
		}

		if !reflect.DeepEqual(got.left, wantLeft) {
			t.Errorf("got %v, want %v", got.left, wantLeft)
		}

		if !reflect.DeepEqual(got.right, wantRight) {
			t.Errorf("got %v, want %v", got.right, wantRight)
		}

		if !reflect.DeepEqual(got.similarities, wantSimilarities) {
			t.Errorf("got %v, want %v", got.similarities, wantSimilarities)
		}
	})

	t.Run("calculateDistance - distance is calculated correctly", func(t *testing.T) {
		lists, _ := processPairs(pairs)

		got, err := calculateDistance(lists.left, lists.right)
		want := 11

		if err != nil {
			t.Fatal("got unexpected error")
		}

		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})

	t.Run("calculateSimilarityScore - similarity score is calculated correctly", func(t *testing.T) {
		lists, _ := processPairs(pairs)

		got := calculateSimilarityScore(lists.left, lists.similarities)
		want := 31

		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})
}
