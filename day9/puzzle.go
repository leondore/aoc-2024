package day9

import (
	"os"
	"strconv"
	"strings"
)

var digitMap = map[byte]int{
	'0': 0,
	'1': 1,
	'2': 2,
	'3': 3,
	'4': 4,
	'5': 5,
	'6': 6,
	'7': 7,
	'8': 8,
	'9': 9,
}

func CalculateSpace(diskmap []byte) (total int, used int, free int) {
	for i := 0; i < len(diskmap); i++ {
		total += digitMap[diskmap[i]]

		if i%2 == 0 {
			used += digitMap[diskmap[i]]
		} else {
			free += digitMap[diskmap[i]]
		}
	}
	return
}

func Day9(path string) (int, string) {
	diskmap, _ := os.ReadFile(path)

	checksum := 0
	emptySpace := 0
	var compact strings.Builder

	pos := 0
	queue := [][2]int{}

	dequeue := func() (int, int) {
		queued := queue[0]
		queue = queue[1:]

		return queued[0], queued[1]
	}

	for i := 0; i < len(diskmap); i++ {
		space := digitMap[diskmap[i]]
		id := i / 2

		if i%2 == 0 {
			checksum += calculateChecksum(id, pos, pos+space)
			compact.WriteString(strings.Repeat(strconv.Itoa(id), space))

			pos += space
		} else {
			emptySpace += space

			if space == 0 {
				continue
			}

			freeSpace := space
			for freeSpace > 0 {
				var tailSpace int

				if len(queue) > 0 {
					id, tailSpace = dequeue()
				} else {
					tailIdx := len(diskmap) - 1
					id = tailIdx / 2
					tailSpace = digitMap[diskmap[tailIdx]]

					emptySpace += digitMap[diskmap[tailIdx-1]]
					diskmap = diskmap[:len(diskmap)-2]
				}

				allocatedSpace := tailSpace

				freeSpace -= tailSpace
				if freeSpace < 0 {
					allocatedSpace = tailSpace + freeSpace
					queue = append(queue, [2]int{id, -freeSpace})
				}

				checksum += calculateChecksum(id, pos, pos+allocatedSpace)
				compact.WriteString(strings.Repeat(strconv.Itoa(id), allocatedSpace))

				pos += allocatedSpace
			}
		}
	}

	for len(queue) > 0 {
		id, tailSpace := dequeue()
		checksum += calculateChecksum(id, pos, pos+tailSpace)
		compact.WriteString(strings.Repeat(strconv.Itoa(id), tailSpace))

		pos += tailSpace
	}

	compact.WriteString(strings.Repeat(".", emptySpace))

	return checksum, compact.String()
}

func calculateChecksum(id, start, end int) (res int) {
	for i := start; i < end; i++ {
		res += i * id
	}
	return
}
