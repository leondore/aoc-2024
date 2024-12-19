package day11

import "testing"

func TestCountDigits(t *testing.T) {
	cases := []struct {
		num  int
		want int
	}{
		{2097446912, 10},
		{2, 1},
		{2024, 4},
		{14168, 5},
	}

	for _, c := range cases {
		got := countDigits(c.num)

		if got != c.want {
			t.Errorf("expected %d to have %d digits, but got %d", c.num, c.want, got)
		}
	}
}

func TestSplitNumber(t *testing.T) {
	cases := []struct {
		num, digits, left, right int
	}{
		{2097446912, 10, 20974, 46912},
		{2024, 4, 20, 24},
		{99, 2, 9, 9},
		{910056, 6, 910, 56},
	}

	for _, c := range cases {
		l, r := splitNumber(c.num, c.digits)

		if l != c.left {
			t.Errorf("first half of %d is %d, but got %d", c.num, c.left, l)
		}

		if r != c.right {
			t.Errorf("second half of %d is %d, but got %d", c.num, c.right, r)
		}
	}
}
