package day2

import (
	"fmt"
	"strconv"
	"strings"
	"sync"
	"sync/atomic"

	"github.com/leondore/aoc-2024/utils"
)

const maxDiff = 3

func Day2(path string) (int, error) {
	reports, err := utils.ProcessInput(path)
	if err != nil {
		return 0, fmt.Errorf("error processing input file: %w", err)
	}

	var count atomic.Uint32
	var wg sync.WaitGroup

	for _, r := range reports {
		wg.Add(1)

		go func() {
			if isReportSafe(r) {
				count.Add(1)
			}
			wg.Done()
		}()
	}

	wg.Wait()

	return int(count.Load()), nil
}

func isReportSafe(report string) bool {
	asc := true
	s := strings.Split(report, " ")

	if len(s) < 2 {
		return true
	}

	for i := 1; i < len(s); i++ {
		x, err := strconv.Atoi(s[i-1])
		if err != nil {
			return false
		}
		y, err := strconv.Atoi(s[i])
		if err != nil {
			return false
		}

		if i == 1 && y < x {
			asc = false
		}

		if x == y || y > x != asc || absDiff(x, y) > maxDiff {
			return false
		}
	}

	return true
}

func absDiff(x, y int) (z int) {
	z = x - y
	if z < 0 {
		z = -z
	}
	return
}
