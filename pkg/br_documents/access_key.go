package br_documents

import (
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

// AllDigitsAreIdentical checks if all digits in the slice are identical.
func AllDigitsAreIdentical(digits []int) bool {
	if len(digits) == 0 {
		return false
	}
	first := digits[0]
	for _, digit := range digits[1:] {
		if digit != first {
			return false
		}
	}
	return true
}

// AccessKeyConfig holds configuration options for generating an Access Key.
type AccessKeyConfig struct {
	Masked bool
	CNPJ   string
}

// AccessKey generates a valid random Chave de Acesso for NF-e.
// If Masked is true, it returns the formatted Chave de Acesso (e.g., with separators).
// If Masked is false, it returns the raw 44-digit Chave de Acesso.
func AccessKey(configs ...AccessKeyConfig) string {
	// Initialize default configuration
	config := AccessKeyConfig{
		Masked: false,
		CNPJ:   CNPJ(), // Assuming GenerateCNPJ is defined elsewhere in the faker package
	}

	// Override defaults with provided configurations
	if len(configs) > 0 {
		config = configs[0]
		if config.CNPJ == "" {
			config.CNPJ = CNPJ()
		}
	}

	// List of valid UF (Federative Unit) codes
	ufCodes := []string{
		"11", "12", "13", "14", "15", "16", "17",
		"21", "22", "23", "24", "25", "26", "27",
		"28", "29", "31", "32", "33", "35",
		"41", "42", "43", "50", "51", "52", "53",
	}

	// 1. UF Code: 2 digits
	uf := ufCodes[rand.Intn(len(ufCodes))]

	// 2. Year and Month: 4 digits (AAMM)
	currentTime := time.Now()
	yearMonth := fmt.Sprintf("%02d%02d", currentTime.Year()%100, currentTime.Month())

	// 3. CNPJ: 14 digits (provided via config)

	// 4. Model: 2 digits (assuming "55" for NF-e)
	model := "55"

	// 5. Series: 3 digits (000 to 999)
	series := fmt.Sprintf("%03d", rand.Intn(1000))

	// 6. Invoice Number (nNF): 9 digits (000000001 to 999999999)
	invoiceNumber := fmt.Sprintf("%09d", rand.Intn(1000000000))

	// 7. Emission Type (tpEmis): 1 digit (1 to 7)
	emissionType := fmt.Sprintf("%d", rand.Intn(7)+1)

	// 8. Numeric Code (cNF): 8 digits (00000000 to 99999999)
	numericCode := fmt.Sprintf("%08d", rand.Intn(100000000))

	// Assemble the first 43 digits of the Access Key
	partialKey := uf + yearMonth + config.CNPJ + model + series + invoiceNumber + emissionType + numericCode

	// Calculate the Verification Digit (DV)
	verificationDigit := calculateAccessKeyDV(partialKey)

	// Complete the Access Key by appending the DV
	fullKey := partialKey + strconv.Itoa(verificationDigit)

	if config.Masked {
		return FormatAccessKey(fullKey)
	}

	return fullKey
}

// calculateAccessKeyDV calculates the Verification Digit (DV) for the Access Key.
// It uses the modulo 11 algorithm as specified.
func calculateAccessKeyDV(key string) int {
	// Convert key to a slice of integers
	digits := make([]int, len(key))
	for i, char := range key {
		digit, err := strconv.Atoi(string(char))
		if err != nil {
			digits[i] = 0
		} else {
			digits[i] = digit
		}
	}

	// Define multipliers: start from the right, repeating 2 to 9
	multipliers := []int{2, 3, 4, 5, 6, 7, 8, 9}
	sum := 0
	multiplierIndex := 0

	// Iterate from right to left
	for i := len(digits) - 1; i >= 0; i-- {
		sum += digits[i] * multipliers[multiplierIndex]
		multiplierIndex = (multiplierIndex + 1) % len(multipliers)
	}

	remainder := sum % 11
	if remainder == 0 || remainder == 1 {
		return 0
	}
	return 11 - remainder
}

// FormatAccessKey formats the Access Key with standard separators.
// The standard format typically doesn't include separators, but for demonstration,
// we'll group digits as follows: 4-4-4-4-4-4-4-4-4-4-4 (total 44 digits).
func FormatAccessKey(key string) string {
	if len(key) != 44 {
		return key
	}

	var formattedKey strings.Builder
	for i, char := range key {
		formattedKey.WriteRune(char)
		if (i+1)%4 == 0 && i != 43 {
			formattedKey.WriteString(" ")
		}
	}
	return formattedKey.String()
}

// init seeds the random number generator to ensure varied outputs.
func init() {
	rand.NewSource(time.Now().UnixNano())
}
