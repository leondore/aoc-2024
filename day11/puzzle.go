package day11

import (
	"math"
	"strconv"
)

const (
	maxBlinks  = 75
	multiplier = 2024
)

func Day11(stones []int) int {
	cache := map[string]int{}
	result := 0

	for _, stone := range stones {
		result += alterStone(stone, 0, maxBlinks, cache)
	}

	return result
}

func countDigits(num int) int {
	digits := 1
	for num/int(math.Pow10(digits)) > 0 {
		digits++
	}
	return digits
}

func splitNumber(num, digits int) (int, int) {
	left := num / int(math.Pow10(digits/2))
	right := num % int(math.Pow10(digits/2))
	return left, right
}

func alterStone(num, blinks, max int, cache map[string]int) int {
	if blinks == max {
		return 1
	}

	blinks++
	digits := countDigits(num)

	key := strconv.Itoa(num) + "-" + strconv.Itoa(blinks)

	if _, ok := cache[key]; !ok {
		switch {
		case num == 0:
			cache[key] = alterStone(1, blinks, max, cache)
		case digits%2 == 0:
			l, r := splitNumber(num, digits)
			cache[key] = alterStone(l, blinks, max, cache) + alterStone(r, blinks, max, cache)
		default:
			cache[key] = alterStone(num*multiplier, blinks, max, cache)
		}
	}

	return cache[key]
}
