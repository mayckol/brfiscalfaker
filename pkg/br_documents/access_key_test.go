package br_documents

import (
	"strings"
	"testing"
	"time"

	v2 "github.com/mayckol/brfiscalfaker/pkg/br_documents/v2"
)

// referenceDV computes the access key check digit independently: modulo 11
// with weights 2..9 from the right and every character valued ASCII - 48.
func referenceDV(partial string) int {
	sum, weight := 0, 2
	for i := len(partial) - 1; i >= 0; i-- {
		sum += (int(partial[i]) - '0') * weight
		weight++
		if weight > 9 {
			weight = 2
		}
	}
	if r := sum % 11; r > 1 {
		return 11 - r
	}
	return 0
}

func TestAccessKeyDV(t *testing.T) {
	tests := []struct {
		name    string
		partial string
		want    int
	}{
		// Real numeric CF-e key (DV 2), unchanged by the alphanumeric rule.
		{"numeric CF-e", "3524100634654500153659001197341007639729954", 2},
		{"alphanumeric CNPJ", "35260712ABC34501DE35550010000000011" + "12345678", referenceDV("35260712ABC34501DE35550010000000011" + "12345678")},
		{"lower case is upper-cased", "35260712abc34501de3555001000000001112345678", referenceDV("35260712ABC34501DE35550010000000011" + "12345678")},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			if got := AccessKeyDV(tc.partial); got != tc.want {
				t.Errorf("AccessKeyDV = %d, want %d", got, tc.want)
			}
		})
	}
}

func TestAccessKeyWithAlphanumericCNPJ(t *testing.T) {
	issued := time.Date(2026, 8, 3, 10, 0, 0, 0, time.UTC)
	for i := 0; i < 50; i++ {
		cnpj := v2.CNPJ(v2.CNPJv2Config{Masked: true})
		key := AccessKey(AccessKeyConfig{CNPJ: cnpj, Model: "65", UF: "35", IssuedAt: issued})
		if len(key) != 44 || !ValidateAccessKey(key) {
			t.Fatalf("key %q is not a valid 44-character key", key)
		}
		if key[:6] != "352608" || key[6:20] != v2.Normalize(cnpj) || key[20:22] != "65" {
			t.Fatalf("key %q does not embed UF, AAMM, CNPJ %s and model", key, cnpj)
		}
	}
}

func TestValidateAccessKey(t *testing.T) {
	valid := "35241006346545001536590011973410076397299542"
	if !ValidateAccessKey(valid) {
		t.Errorf("%s should be valid", valid)
	}
	for _, bad := range []string{valid[:43] + "3", valid[:43], strings.Repeat("A", 44)} {
		if ValidateAccessKey(bad) {
			t.Errorf("%s should be invalid", bad)
		}
	}
}
