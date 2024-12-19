package day11

import (
	"math"
)

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
