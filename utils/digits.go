package utils

import (
	"math/rand"
	"strconv"
	"strings"
)

// DigitsToString converts a slice of integers to a concatenated string of digits.
func DigitsToString(digits []int) string {
	var sb strings.Builder
	for _, digit := range digits {
		sb.WriteString(strconv.Itoa(digit))
	}
	return sb.String()
}

// GenerateRandomDigits generates a slice of random digits of specified length.
func GenerateRandomDigits(length int) []int {
	digits := make([]int, length)
	for i := 0; i < length; i++ {
		digits[i] = rand.Intn(10)
	}
	return digits
}
