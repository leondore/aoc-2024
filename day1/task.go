package day1

import (
	"errors"
	"fmt"
	"slices"
	"strconv"
	"strings"

	"github.com/leondore/aoc-2024/utils"
)

type Lists struct {
	left         []int
	right        []int
	similarities map[int]int
}

func Day1(path string) (int, int, error) {
	input, err := utils.ProcessInput(path)
	if err != nil {
		return 0, 0, fmt.Errorf("error processing input file: %w", err)
	}

	lists, err := processPairs(input)
	if err != nil {
		return 0, 0, fmt.Errorf("error with pairs: %w", err)
	}

	distance, err := calculateDistance(lists.left, lists.right)
	if err != nil {
		return 0, 0, fmt.Errorf("errors with pairs: %w", err)
	}

	score := calculateSimilarityScore(lists.left, lists.similarities)

	return distance, score, nil
}

func processPairs(input []string) (*Lists, error) {
	left := make([]int, len(input))
	right := make([]int, len(input))
	similarities := make(map[int]int)

	for i := 0; i < len(input); i++ {
		pair := strings.Split(input[i], " ")
		if len(pair) < 2 {
			return &Lists{}, errors.New("malformed input")
		}

		f, err := strconv.Atoi(pair[0])
		if err != nil {
			return &Lists{}, errors.New("pair is not a number")
		}
		s, err := strconv.Atoi(pair[1])
		if err != nil {
			return &Lists{}, errors.New("pair is not a number")
		}

		left[i], right[i] = f, s
		similarities[s]++
	}

	slices.Sort(left)
	slices.Sort(right)

	return &Lists{left, right, similarities}, nil
}

func calculateDistance(left []int, right []int) (int, error) {
	if len(left) != len(right) {
		return 0, errors.New("pairs don't match")
	}

	sumDistance := 0

	for i := 0; i < len(left); i++ {
		if left[i] < right[i] {
			sumDistance += (right[i] - left[i])
		} else {
			sumDistance += (left[i] - right[i])
		}
	}

	return sumDistance, nil
}

func calculateSimilarityScore(left []int, similarities map[int]int) int {
	score := 0

	for _, v := range left {
		score += (v * similarities[v])
	}

	return score
}
