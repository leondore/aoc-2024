package day3

import (
	"reflect"
	"testing"
)

func TestDay3(t *testing.T) {
	test := "xmul(2,4)&mul[3,7]!^don't()_mul(5,5)+mul(32,64](mul(11,8)undo()?mul(8,5))"

	t.Run("findAllMatches", func(t *testing.T) {
		want := []string{
			"mul(2,4)",
			"don't()",
			"mul(5,5)",
			"mul(11,8)",
			"do()",
			"mul(8,5)",
		}
		got, err := findAllMatches(test)

		if err != nil {
			t.Fatal("got unexpected error")
		}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %+v, want %+v", got, want)
		}
	})

	t.Run("parseInstructions", func(t *testing.T) {
		cases := []struct {
			input  string
			result int
		}{
			{"mul(2,4)", 8},
			{"mul(5,5)", 25},
			{"mul(11,8)", 88},
			{"mul(8,5)", 40},
		}

		for _, test := range cases {
			got := parseInstructions(test.input)

			if got != test.result {
				t.Errorf("got %d, want %d", got, test.result)
			}
		}
	})
}
