// Package v2 implements Brazil's alphanumeric CNPJ format, effective July 2026.
//
// A CNPJ has 14 positions: the first 12 are alphanumeric (0-9, A-Z) and the
// last 2 are numeric check digits. Check digits use Módulo 11 where every
// character is converted to a number via ASCII - 48 (so '0'-'9' map to 0-9 and
// 'A'-'Z' map to 17-42). Legacy numeric-only CNPJs remain valid.
package v2

import (
	"math/rand"
	"strings"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

const (
	cnpjLength      = 14
	cnpjBaseLength  = 12
	asciiZeroOffset = 48
)

// ambiguousLetters are excluded from generation by default to avoid visual
// confusion (I/1, O/0, Q/O, U/V, F/E). They remain valid on input.
const ambiguousLetters = "IOQUF"

var dv1Weights = []int{5, 4, 3, 2, 9, 8, 7, 6, 5, 4, 3, 2}
var dv2Weights = []int{6, 5, 4, 3, 2, 9, 8, 7, 6, 5, 4, 3, 2}

// CNPJv2Config holds options for generating an alphanumeric CNPJ.
type CNPJv2Config struct {
	Masked                bool
	AllowAmbiguousLetters bool
}

// CNPJ generates a valid random alphanumeric CNPJ.
// Masked output is formatted as XX.XXX.XXX/YYYY-ZZ.
func CNPJ(configs ...CNPJv2Config) string {
	config := CNPJv2Config{}
	if len(configs) > 0 {
		config = configs[0]
	}

	base := make([]byte, cnpjBaseLength)
	charset := generationCharset(config.AllowAmbiguousLetters)
	for i := range base {
		base[i] = charset[rand.Intn(len(charset))]
	}

	dv1 := checkDigit(base, dv1Weights)
	dv2 := checkDigit(append(base, dv1), dv2Weights)

	cnpj := string(base) + string(dv1) + string(dv2)
	if config.Masked {
		return mask(cnpj)
	}
	return cnpj
}

// ValidateCNPJ reports whether s is a valid CNPJ in either the numeric (legacy)
// or alphanumeric format. Masks, surrounding whitespace and lower case are
// accepted and normalized before validation.
func ValidateCNPJ(s string) bool {
	cnpj := sanitize(s)
	if len(cnpj) != cnpjLength {
		return false
	}
	if cnpj == "00000000000000" {
		return false
	}

	for i := 0; i < cnpjBaseLength; i++ {
		if !isAlphanumericUpper(cnpj[i]) {
			return false
		}
	}
	if !isDigit(cnpj[12]) || !isDigit(cnpj[13]) {
		return false
	}

	base := []byte(cnpj[:cnpjBaseLength])
	dv1 := checkDigit(base, dv1Weights)
	dv2 := checkDigit(append(base, dv1), dv2Weights)

	return cnpj[12] == dv1 && cnpj[13] == dv2
}

// checkDigit computes a single Módulo 11 check digit and returns it as an ASCII
// byte ('0'-'9'). Each input character is valued as ASCII - 48.
func checkDigit(chars []byte, weights []int) byte {
	sum := 0
	for i, c := range chars {
		sum += (int(c) - asciiZeroOffset) * weights[i]
	}

	remainder := sum % 11
	if remainder < 2 {
		return '0'
	}
	return byte('0' + (11 - remainder))
}

// sanitize uppercases s and strips every character outside [A-Z0-9].
func sanitize(s string) string {
	var b strings.Builder
	for _, r := range strings.ToUpper(s) {
		if (r >= '0' && r <= '9') || (r >= 'A' && r <= 'Z') {
			b.WriteRune(r)
		}
	}
	return b.String()
}

func generationCharset(allowAmbiguous bool) string {
	const digits = "0123456789"
	const letters = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	if allowAmbiguous {
		return digits + letters
	}

	var b strings.Builder
	b.WriteString(digits)
	for i := 0; i < len(letters); i++ {
		if !strings.ContainsRune(ambiguousLetters, rune(letters[i])) {
			b.WriteByte(letters[i])
		}
	}
	return b.String()
}

func mask(cnpj string) string {
	if len(cnpj) != cnpjLength {
		return cnpj
	}
	return cnpj[0:2] + "." + cnpj[2:5] + "." + cnpj[5:8] + "/" + cnpj[8:12] + "-" + cnpj[12:14]
}

func isDigit(c byte) bool { return c >= '0' && c <= '9' }

func isAlphanumericUpper(c byte) bool {
	return isDigit(c) || (c >= 'A' && c <= 'Z')
}

// Normalize uppercases s and strips every character outside [A-Z0-9], turning
// a masked CNPJ (numeric or alphanumeric) into its raw 14-character form.
func Normalize(s string) string {
	return sanitize(s)
}

// CompleteCNPJ appends the two check digits to a 12-character base (numeric or
// alphanumeric; masks and lower case are accepted). It returns "" when the
// normalized base is not 12 characters long.
func CompleteCNPJ(base string) string {
	b := []byte(sanitize(base))
	if len(b) != cnpjBaseLength {
		return ""
	}
	dv1 := checkDigit(b, dv1Weights)
	dv2 := checkDigit(append(b, dv1), dv2Weights)
	return string(b) + string(dv1) + string(dv2)
}

// HasValidFormat reports whether s, once normalized, has the CNPJ layout:
// 12 alphanumeric characters followed by 2 digits. It does not check the
// check digits; use ValidateCNPJ for that.
func HasValidFormat(s string) bool {
	cnpj := sanitize(s)
	if len(cnpj) != cnpjLength {
		return false
	}
	return isDigit(cnpj[12]) && isDigit(cnpj[13])
}
