package day7

import (
	"testing"
)

func TestCalculate(t *testing.T) {
	cases := []struct {
		operands []int
		result   int
		want     bool
	}{
		{[]int{10, 19}, 190, true},
		{[]int{81, 40, 27}, 3267, true},
		{[]int{6, 8, 6, 15}, 7290, false},
		{[]int{16, 10, 13}, 161011, false},
		{[]int{11, 6, 16, 20}, 292, true},
	}

	for _, c := range cases {
		got := calculate(c.operands, c.result)

		if got != c.want {
			t.Errorf("calculate(%v, %d) = %v; want %v", c.operands, c.result, got, c.want)
		}
	}
}

func TestParseEquation(t *testing.T) {
	cases := []struct {
		eq       string
		result   int
		operands []int
	}{
		{"190: 10 19", 190, []int{10, 19}},
		{"3267: 81 40 27", 3267, []int{81, 40, 27}},
		{"7290: 6 8 6 15", 7290, []int{6, 8, 6, 15}},
		{"161011: 16 10 13", 161011, []int{16, 10, 13}},
		{"292: 11 6 16 20", 292, []int{11, 6, 16, 20}},
	}

	for _, c := range cases {
		result, operands, err := parseEquation(c.eq)

		if err != nil {
			t.Fatalf("parseEquation(%s) returned error: %v", c.eq, err)
		}

		if result != c.result {
			t.Errorf("parseEquation(%s) returned result %d; want %d", c.eq, result, c.result)
		}

		if len(operands) != len(c.operands) {
			t.Errorf("parseEquation(%s) returned operands %v; want %v", c.eq, operands, c.operands)
		}
	}
}

func TestConcatenate(t *testing.T) {
	cases := []struct {
		x    int
		y    int
		want int
	}{
		{1, 2, 12},
		{3, 4, 34},
		{5, 60, 560},
		{70, 8, 708},
		{9, 100, 9100},
	}

	for _, c := range cases {
		got := concatenate(c.x, c.y)

		if got != c.want {
			t.Errorf("concatenate(%d, %d) = %d; want %d", c.x, c.y, got, c.want)
		}
	}
}

func TestDay7(t *testing.T) {
	inputPath := "./test.txt"

	want := 11387
	got, err := Day7(inputPath)

	if err != nil {
		t.Fatalf("Day7() returned error: %v", err)
	}

	if got != want {
		t.Errorf("Day7() = %d; want %d", got, want)
	}
}
