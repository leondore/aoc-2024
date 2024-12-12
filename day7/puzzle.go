package day7

import (
	"strconv"
	"strings"
	"sync"

	"github.com/leondore/aoc-2024/utils"
)

func Day7(path string) (int, error) {
	equations, err := utils.ProcessInput(path)
	if err != nil {
		return 0, err
	}

	var wg sync.WaitGroup
	resultsChan := make(chan int, len(equations))

	for _, eq := range equations {
		result, operands, err := parseEquation(eq)
		if err != nil {
			return 0, err
		}

		wg.Add(1)
		go func() {
			if calculate(operands, result) {
				resultsChan <- result
			}
			wg.Done()
		}()
	}
	wg.Wait()
	close(resultsChan)

	sum := 0
	for num := range resultsChan {
		sum += num
	}
	return sum, nil
}

func calculate(operands []int, result int) bool {
	sum := operands[0] + operands[1]
	mult := operands[0] * operands[1]
	concat := concatenate(operands[0], operands[1])

	if len(operands) == 2 {
		return sum == result || mult == result || concat == result
	}

	return calculate(append([]int{sum}, operands[2:]...), result) ||
		calculate(append([]int{mult}, operands[2:]...), result) ||
		calculate(append([]int{concat}, operands[2:]...), result)
}

func parseEquation(eq string) (int, []int, error) {
	eqParts := strings.Split(eq, ": ")

	result, err := strconv.Atoi(eqParts[0])
	if err != nil {
		return 0, nil, err
	}

	operands := []int{}
	for _, op := range strings.Split(eqParts[1], " ") {
		num, err := strconv.Atoi(op)
		if err != nil {
			return 0, nil, err
		}

		operands = append(operands, num)
	}

	return result, operands, nil
}

func concatenate(x, y int) int {
	if y < 10 {
		return x*10 + y
	}
	if y < 100 {
		return x*100 + y
	}
	return x*1000 + y
}
