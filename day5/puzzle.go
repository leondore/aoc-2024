package day5

import (
	"bufio"
	"os"
	"slices"
	"strconv"
	"strings"
)

func Day5(instructionsPath, updatesPath string) (int, int, error) {
	instructions, err := ParseInstructions(instructionsPath)
	if err != nil {
		return 0, 0, err
	}

	file, err := os.Open(updatesPath)
	if err != nil {
		return 0, 0, err
	}
	defer file.Close()

	sumOrdered := 0
	sumUnordered := 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		update := strings.Split(line, ",")

		if IsUpdateSorted(instructions, update) {
			num, err := getUpdateMid(update)
			if err != nil {
				return 0, 0, err
			}
			sumOrdered += num
		} else {
			slices.SortStableFunc(update, sortFunc(instructions))
			num, err := getUpdateMid(update)
			if err != nil {
				return 0, 0, err
			}
			sumUnordered += num
		}
	}

	return sumOrdered, sumUnordered, nil
}

func getUpdateMid(update []string) (int, error) {
	mid := update[len(update)/2]
	return strconv.Atoi(mid)
}

func IsUpdateSorted(instructions map[string]map[string]int, update []string) bool {
	return slices.IsSortedFunc(update, sortFunc(instructions))
}

func ParseInstructions(path string) (map[string]map[string]int, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	instructions := map[string]map[string]int{}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		instruction := scanner.Text()
		x, y := instruction[:2], instruction[3:]

		if _, ok := instructions[x]; !ok {
			instructions[x] = map[string]int{}
		}
		if _, ok := instructions[y]; !ok {
			instructions[y] = map[string]int{}
		}

		instructions[x][y], instructions[y][x] = 1, -1
	}

	return instructions, nil
}

func sortFunc(instructions map[string]map[string]int) func(a, b string) int {
	return func(a, b string) int {
		if instructions[b] == nil {
			return 0
		}
		return instructions[b][a]
	}
}
