package day9

import (
	"os"
	"reflect"
	"testing"
)

func TestDay9(t *testing.T) {
	disk, _ := os.ReadFile("./test.txt")
	diskShort := disk[:5]

	t.Run("correctly calculates used and free space in disk", func(t *testing.T) {
		used, free := CalculateSpace(disk)
		want := struct {
			used int
			free int
		}{
			used: 28,
			free: 14,
		}

		if used != want.used {
			t.Errorf("got %d, want %d", used, want.used)
		}

		if free != want.free {
			t.Errorf("got %d, want %d", free, want.free)
		}
	})

	t.Run("generates slices with free blocks and used blocks", func(t *testing.T) {
		used, free := MapDisk(diskShort)
		wantUsed := SpaceList{
			{2, 0, 2, 0, "0", false},
			{3, 5, 8, 1, "1", false},
			{1, 11, 12, 2, "2", false},
		}
		wantFree := SpaceList{
			{3, 2, 5, -1, blankChar, false},
			{3, 8, 11, -1, blankChar, false},
		}

		if !reflect.DeepEqual(used, wantUsed) {
			t.Errorf("got %v, want %v", used, wantUsed)
		}

		if !reflect.DeepEqual(free, wantFree) {
			t.Errorf("got %v, want %v", free, wantFree)
		}
	})

	t.Run("finds the correct index in which to insert", func(t *testing.T) {
		used, free := MapDisk(disk)

		toMove := &used[len(used)-1]
		free.FindRoom(toMove)

		if !toMove.Moved {
			t.Errorf("file should have moved, but didn't")
		}

		if toMove.Start != 2 {
			t.Errorf("file should have move to position %d, but is in position %d", 2, toMove.Start)
		}
	})
}
