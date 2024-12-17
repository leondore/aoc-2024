package day9

import (
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

const blankChar = "."

type Space struct {
	Blocks, Start, End, Id int
	Literal                string
	Moved                  bool
}

type SpaceList []Space

func (sl *SpaceList) FindRoom(s *Space) {
	for i := len(*sl) - 1; i >= 0; i-- {
		if (*sl)[i].Start >= s.Start {
			b := *sl
			b = b[:len(b)-1]
			*sl = b
		}
	}

	for i := 0; i < len(*sl); i++ {
		if s.Blocks <= (*sl)[i].Blocks {
			s.Moved = true
			s.Start = (*sl)[i].Start
			s.End = s.Start + s.Blocks

			if (*sl)[i].Blocks-s.Blocks != 0 {
				newBlank := Space{
					Blocks:  (*sl)[i].Blocks - s.Blocks,
					Start:   s.End,
					End:     s.End + ((*sl)[i].Blocks - s.Blocks),
					Id:      -1,
					Literal: blankChar,
					Moved:   false,
				}
				(*sl)[i] = newBlank
			} else {
				temp := *sl
				*sl = append(temp[:i], temp[i+1:]...)
			}
			break
		}
	}
}

func CompactFiles(path string) (int, string, error) {
	disk, err := os.ReadFile(path)
	if err != nil {
		return 0, "", fmt.Errorf("could not process file: %w", err)
	}

	used, free := MapDisk(disk)

	for i := len(used) - 1; i >= 0; i-- {
		if len(free) == 0 {
			break
		}

		free.FindRoom(&used[i])
	}

	slices.SortStableFunc(used, func(a, b Space) int {
		return a.Start - b.Start
	})

	var compacted strings.Builder
	pos := 0
	for _, f := range used {
		if f.Start > pos {
			compacted.WriteString(strings.Repeat(blankChar, f.Start-pos))
		}
		compacted.WriteString(strings.Repeat(f.Literal, f.Blocks))
		pos = f.End
	}

	checksum := 0
	for _, f := range used {
		checksum += calculateChecksum(f.Id, f.Start, f.End)
	}

	return checksum, compacted.String(), nil
}

func MapDisk(disk []byte) (SpaceList, SpaceList) {
	usedSpace := make(SpaceList, len(disk)/2+1)
	freeSpace := make(SpaceList, len(disk)/2)

	pos := 0
	for i := 0; i < len(disk); i++ {
		blockCount := DigitMap[disk[i]]
		space := Space{
			Blocks: blockCount,
			Start:  pos,
			End:    pos + blockCount,
			Moved:  false,
		}
		if i%2 == 0 {
			space.Id = i / 2
			space.Literal = strconv.Itoa(space.Id)
			usedSpace[i/2] = space
		} else {
			space.Id = -1
			space.Literal = blankChar
			freeSpace[i/2] = space
		}
		pos += blockCount
	}

	return usedSpace, freeSpace
}
