package br_documents

import (
	"github.com/mayckol/brfiscalfaker/utils"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

// init seeds the random number generator to ensure varied outputs.
func init() {
	rand.Seed(time.Now().UnixNano())
}

// CNPJConfig holds configuration options for generating a CNPJ.
type CNPJConfig struct {
	Masked bool
}

// CNPJ generates a valid random CNPJ number.
// If Masked is true, it returns the formatted CNPJ (e.g., XX.XXX.XXX/XXXX-XX).
// If Masked is false, it returns the raw CNPJ digits (e.g., XXXXXXXXXXXXXXX).
func CNPJ(configs ...CNPJConfig) string {
	// Initialize default configuration
	config := CNPJConfig{
		Masked: false,
	}

	// Override defaults with provided configurations
	if len(configs) > 0 {
		config = configs[0]
	}

	// Generate the first 12 digits of the CNPJ
	cnpjDigits := utils.GenerateRandomDigits(12)

	// Calculate the first check digit
	firstCheckDigit := calculateCNPJCheckDigit(cnpjDigits)
	cnpjDigits = append(cnpjDigits, firstCheckDigit)

	// Calculate the second check digit
	secondCheckDigit := calculateCNPJCheckDigit(cnpjDigits)
	cnpjDigits = append(cnpjDigits, secondCheckDigit)

	if config.Masked {
		return formatCNPJ(cnpjDigits)
	}

	return utils.DigitsToString(cnpjDigits)
}

// calculateCNPJCheckDigit calculates a CNPJ check digit using the modulo 11 algorithm.
func calculateCNPJCheckDigit(cnpj []int) int {
	var weights []int

	// Define weights based on the position
	if len(cnpj) <= 12 {
		// Weights for the first check digit
		weights = []int{5, 4, 3, 2, 9, 8, 7, 6, 5, 4, 3, 2}
	} else {
		// Weights for the second check digit
		weights = []int{6, 5, 4, 3, 2, 9, 8, 7, 6, 5, 4, 3, 2}
	}

	sum := 0
	for i, weight := range weights {
		sum += cnpj[i] * weight
	}

	remainder := sum % 11
	if remainder < 2 {
		return 0
	}
	return 11 - remainder
}

// formatCNPJ formats a slice of CNPJ digits into the standard format XX.XXX.XXX/XXXX-XX.
func formatCNPJ(cnpj []int) string {
	if len(cnpj) != 14 {
		return utils.DigitsToString(cnpj)
	}

	var formatted strings.Builder
	for i, digit := range cnpj {
		switch i {
		case 2, 5:
			formatted.WriteString(".")
		case 8:
			formatted.WriteString("/")
		case 12:
			formatted.WriteString("-")
		}
		formatted.WriteString(strconv.Itoa(digit))
	}
	return formatted.String()
}
