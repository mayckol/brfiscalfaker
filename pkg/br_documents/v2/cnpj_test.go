package v2

import (
	"strings"
	"testing"
)

func TestValidateCNPJ_OfficialRFExample(t *testing.T) {
	// Official Receita Federal example: 12.ABC.345/01DE-35
	cases := []string{
		"12ABC34501DE35",
		"12.ABC.345/01DE-35",
		"12abc34501de35",
		"  12.ABC.345/01DE-35  ",
	}
	for _, c := range cases {
		if !ValidateCNPJ(c) {
			t.Errorf("ValidateCNPJ(%q) = false, want true", c)
		}
	}
}

func TestValidateCNPJ_NumericBackwardCompatible(t *testing.T) {
	cases := []struct {
		cnpj string
		want bool
	}{
		{"11222333000181", true},
		{"11.222.333/0001-81", true},
		{"11222333000182", false},
	}
	for _, c := range cases {
		if got := ValidateCNPJ(c.cnpj); got != c.want {
			t.Errorf("ValidateCNPJ(%q) = %v, want %v", c.cnpj, got, c.want)
		}
	}
}

func TestValidateCNPJ_Invalid(t *testing.T) {
	cases := []string{
		"",
		"12ABC34501DE34",   // wrong DV
		"12ABC34501DE3",    // 13 chars
		"12ABC34501DE355",  // 15 chars
		"12ABC34501DEA5",   // non-numeric DV
		"12@BC34501DE35",   // invalid charset
		"00000000000000",   // structurally invalid DV
	}
	for _, c := range cases {
		if ValidateCNPJ(c) {
			t.Errorf("ValidateCNPJ(%q) = true, want false", c)
		}
	}
}

func TestCNPJ_GeneratesValid(t *testing.T) {
	for i := 0; i < 1000; i++ {
		cnpj := CNPJ()
		if len(cnpj) != cnpjLength {
			t.Fatalf("CNPJ() length = %d, want %d (%q)", len(cnpj), cnpjLength, cnpj)
		}
		if !ValidateCNPJ(cnpj) {
			t.Fatalf("CNPJ() generated invalid value %q", cnpj)
		}
	}
}

func TestCNPJ_ExcludesAmbiguousByDefault(t *testing.T) {
	for i := 0; i < 1000; i++ {
		cnpj := CNPJ()
		if strings.ContainsAny(cnpj[:cnpjBaseLength], ambiguousLetters) {
			t.Fatalf("CNPJ() contained ambiguous letter: %q", cnpj)
		}
	}
}

func TestCNPJ_Masked(t *testing.T) {
	cnpj := CNPJ(CNPJv2Config{Masked: true})
	if len(cnpj) != 18 {
		t.Fatalf("masked length = %d, want 18 (%q)", len(cnpj), cnpj)
	}
	if cnpj[2] != '.' || cnpj[6] != '.' || cnpj[10] != '/' || cnpj[15] != '-' {
		t.Errorf("masked separators wrong: %q", cnpj)
	}
	if !ValidateCNPJ(cnpj) {
		t.Errorf("masked CNPJ %q failed validation", cnpj)
	}
}

func TestCNPJ_AllowAmbiguousLetters(t *testing.T) {
	for i := 0; i < 1000; i++ {
		if !ValidateCNPJ(CNPJ(CNPJv2Config{AllowAmbiguousLetters: true})) {
			t.Fatal("CNPJ with ambiguous letters allowed generated invalid value")
		}
	}
}
