package day6

import "testing"

func TestDay6(t *testing.T) {
	grid, _ := NewGrid("./test.txt")

	t.Run("can find guard in grid", func(t *testing.T) {
		got := grid.FindGuard()
		want := Guard{
			X: 4,
			Y: 6,
		}

		if got != want {
			t.Errorf("got %v, want %v", got, want)
		}
	})

	t.Run("counts all the visited positions", func(t *testing.T) {
		got, _, err := Day6("./test.txt")
		want := 41

		if err != nil {
			t.Fatal(err)
		}

		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})

	t.Run("counts all times guard got stuck in loop", func(t *testing.T) {
		_, got, err := Day6("./test.txt")
		want := 6

		if err != nil {
			t.Fatal(err)
		}

		if got != want {
			t.Errorf("got %d, want %d", got, want)
		}
	})
}
