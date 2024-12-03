package day2

import (
	"reflect"
	"testing"
)

const inputPath = "./test.txt"

func TestDay2(t *testing.T) {
	tests := []string{
		"7 6 4 2 1",
		"1 2 7 8 9",
		"9 7 6 2 1",
		"1 3 2 4 5",
		"8 6 4 4 1",
		"1 3 6 7 9",
	}

	t.Run("absDiff - different subtractions return absolute values", func(t *testing.T) {
		cases := []struct {
			x    int
			y    int
			want int
		}{
			{3, 3, 0},
			{7, 2, 5},
			{2, 7, 5},
			{-2, -3, 1},
			{-3, -2, 1},
		}

		for _, test := range cases {
			got := absDiff(test.x, test.y)
			if got != test.want {
				t.Errorf("got %d, want %d", got, test.want)
			}
		}
	})

	t.Run("isPairSafe - check if pairs pass invariants", func(t *testing.T) {
		cases := []struct {
			x    int
			y    int
			want bool
		}{
			{1, 2, true},
			{8, 2, false},
			{5, 5, false},
		}

		for _, test := range cases {
			got := isPairSafe(test.x, test.y)
			if got != test.want {
				t.Errorf("got %v, want %v - (%d, %d)", got, test.want, test.x, test.y)
			}
		}
	})

	t.Run("marshalReport - report is correctly converted to slice of ints", func(t *testing.T) {
		want := []int{7, 6, 4, 2, 1}
		got, err := marshalReport(tests[0])

		if err != nil {
			t.Fatal("got unexpected error")
		}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v, want %v", got, want)
		}
	})

	t.Run("isReportSafe - correctly analyses report safety", func(t *testing.T) {
		got := []bool{}
		want := []bool{true, false, false, true, true, true}

		for _, test := range tests {
			r, _ := marshalReport(test)
			got = append(got, isReportSafe(r, true))
		}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v, want %v", got, want)
		}
	})

	t.Run("Day2 returns correct count", func(t *testing.T) {
		want := 4
		got, err := Day2(inputPath)

		if err != nil {
			t.Fatal("got unexpected error")
		}

		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})
}
