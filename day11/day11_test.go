package day11

import (
	"fmt"
	"testing"
)

var test = []int{125, 17}

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

func TestAlterStone(t *testing.T) {
	cases := []struct {
		num, stones, max int
	}{
		{test[0], 19025, 25},
		{test[1], 36287, 25},
		{test[0], 661984358, 50},
		{test[1], 1238449243, 50},
		{test[0], 22840618691206, 75},
		{test[1], 42760419959276, 75},
	}

	for _, c := range cases {
		t.Run(fmt.Sprintf("%d-%d", c.num, c.max), func(t *testing.T) {
			cache := map[string]int{}
			got := alterStone(c.num, 0, c.max, cache)
			fmt.Println(got)

			if got != c.stones {
				t.Errorf("wanted %d, got %d", c.stones, got)
			}
		})
	}
}

func TestDay11(t *testing.T) {
	got := Day11(test)
	want := 65601038650482

	if got != want {
		t.Errorf("expected %d stones, but got %d", want, got)
	}
}
