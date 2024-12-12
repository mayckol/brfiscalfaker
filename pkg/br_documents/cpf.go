package br_documents

import (
	"github.com/mayckol/brfiscalfaker/utils"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

// init seeds the global random number generator to ensure varied outputs.
func init() {
	rand.Seed(time.Now().UnixNano())
}

// CPFConfig holds configuration options for generating a CPF.
type CPFConfig struct {
	Masked bool
}

// CPF generates a valid random CPF number.
// If Masked is true, it returns the formatted CPF (e.g., XXX.XXX.XXX-XX).
// If Masked is false, it returns the raw CPF digits (e.g., XXXXXXXXXXX).
func CPF(configs ...CPFConfig) string {
	config := CPFConfig{
		Masked: false,
	}

	if len(configs) > 0 {
		config = configs[0]
	}

	// Generate the first 9 random digits of the CPF
	cpfDigits := utils.GenerateRandomDigits(9)

	// Calculate the first check digit with a starting weight of 10
	firstCheckDigit := calculateCPFCheckDigit(cpfDigits, 10)
	cpfDigits = append(cpfDigits, firstCheckDigit)

	// Calculate the second check digit with a starting weight of 11
	secondCheckDigit := calculateCPFCheckDigit(cpfDigits, 11)
	cpfDigits = append(cpfDigits, secondCheckDigit)

	if config.Masked {
		return formatCPF(cpfDigits)
	}

	return utils.DigitsToString(cpfDigits)
}

// calculateCPFCheckDigit calculates a CPF check digit using the modulo 11 algorithm.
// The weightStart parameter is 10 for the first check digit and 11 for the second.
func calculateCPFCheckDigit(cpf []int, weightStart int) int {
	sum := 0
	for i, digit := range cpf {
		sum += digit * (weightStart - i)
	}

	remainder := sum % 11
	if remainder < 2 {
		return 0
	}
	return 11 - remainder
}

// formatCPF formats a slice of CPF digits into the standard format XXX.XXX.XXX-XX.
func formatCPF(cpf []int) string {
	if len(cpf) != 11 {
		return utils.DigitsToString(cpf)
	}

	var formatted strings.Builder
	for i, digit := range cpf {
		switch i {
		case 3, 6:
			formatted.WriteString(".")
		case 9:
			formatted.WriteString("-")
		}
		formatted.WriteString(strconv.Itoa(digit))
	}
	return formatted.String()
}
