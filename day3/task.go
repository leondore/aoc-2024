package day3

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

const (
	instructionPrefixSize = 4
	instructionSuffixSize = 1

	enableInstruction  = "do()"
	disableInstruction = "don't()"
)

func Day3(path string) (int, error) {
	dat, err := os.ReadFile(path)
	if err != nil {
		return 0, fmt.Errorf("error reading file: %w", err)
	}

	instructions, err := findAllMatches(string(dat))
	if err != nil {
		return 0, err
	}

	sum := 0
	enabled := true

	for i := 0; i < len(instructions); i++ {
		if instructions[i] == enableInstruction {
			enabled = true
			continue
		}
		if instructions[i] == disableInstruction {
			enabled = false
			continue
		}

		if enabled {
			sum += parseInstructions(instructions[i])
		}
	}

	return sum, nil
}

func findAllMatches(str string) ([]string, error) {
	r, err := regexp.Compile(`mul\(\d{1,3},\d{1,3}\)|don't\(\)|do\(\)`)
	if err != nil {
		return nil, fmt.Errorf("bad regexp: %w", err)
	}

	return r.FindAllString(str, -1), nil
}

func parseInstructions(instruction string) int {
	parts := strings.Split(instruction, ",")
	if len(parts) < 2 {
		return 0
	}

	x, err := strconv.Atoi(parts[0][instructionPrefixSize:])
	if err != nil {
		return 0
	}
	y, err := strconv.Atoi(parts[1][:len(parts[1])-instructionSuffixSize])
	if err != nil {
		return 0
	}

	return x * y
}
