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
		report, err := marshalReport(r)
		if err != nil {
			return 0, fmt.Errorf("could not marshall report: %w", err)
		}

		wg.Add(1)
		go func() {
			if isReportSafe(report, true) {
				count.Add(1)
			}
			wg.Done()
		}()
	}

	wg.Wait()

	return int(count.Load()), nil
}

func isReportSafe(report []int, initial bool) bool {
	ascCount, descCount := 0, 0

	if len(report) < 2 {
		return true
	}

	for i := 1; i < len(report); i++ {
		prev, curr := report[i-1], report[i]

		if !isPairSafe(prev, curr) {
			if initial {
				return isReportSafe(append(report[:i], report[i+1:]...), false) || isReportSafe(append(report[:i-1], report[i:]...), false)
			}
			return false
		}

		if curr > prev {
			ascCount += 1
		} else {
			descCount += 1
		}
	}

	comparisons := len(report) - 1
	if initial {
		return ascCount >= (comparisons-1) || descCount >= (comparisons-1)
	}
	return ascCount == comparisons || descCount == comparisons
}

func isPairSafe(x, y int) bool {
	return x != y && absDiff(x, y) <= maxDiff
}

func marshalReport(report string) ([]int, error) {
	marshaled := []int{}

	for _, level := range strings.Split(report, " ") {
		x, err := strconv.Atoi(level)
		if err != nil {
			return marshaled, err
		}
		marshaled = append(marshaled, x)
	}

	return marshaled, nil
}

func absDiff(x, y int) (z int) {
	z = x - y
	if z < 0 {
		z = -z
	}
	return
}
