package nfs

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"io"
	"strconv"
	"strings"
	"testing"
	"time"

	"github.com/mayckol/brfiscalfaker/pkg/br_documents"
	v2 "github.com/mayckol/brfiscalfaker/pkg/br_documents/v2"
)

// parsedXML maps each element path ("NFe/infNFe/emit/CNPJ") and attribute path
// ("NFe/infNFe@Id") to the values found there, in document order.
type parsedXML map[string][]string

func parseXML(t *testing.T, data []byte) parsedXML {
	t.Helper()
	doc := parsedXML{}
	dec := xml.NewDecoder(bytes.NewReader(data))
	var path []string
	var text strings.Builder
	for {
		tok, err := dec.Token()
		if err != nil {
			if err == io.EOF {
				break
			}
			t.Fatalf("generated XML does not parse: %v\n%s", err, data)
		}
		switch el := tok.(type) {
		case xml.StartElement:
			path = append(path, el.Name.Local)
			p := strings.Join(path, "/")
			for _, a := range el.Attr {
				doc[p+"@"+a.Name.Local] = append(doc[p+"@"+a.Name.Local], a.Value)
			}
			text.Reset()
		case xml.CharData:
			text.Write(el)
		case xml.EndElement:
			p := strings.Join(path, "/")
			doc[p] = append(doc[p], strings.TrimSpace(text.String()))
			text.Reset()
			path = path[:len(path)-1]
		}
	}
	return doc
}

// all returns the values of every path that ends with suffix.
func (d parsedXML) all(suffix string) []string {
	var out []string
	for p, v := range d {
		if p == suffix || strings.HasSuffix(p, "/"+suffix) {
			out = append(out, v...)
		}
	}
	return out
}

// one returns the single value at suffix, failing when there is not exactly one.
func (d parsedXML) one(t *testing.T, suffix string) string {
	t.Helper()
	v := d.all(suffix)
	if len(v) != 1 {
		t.Fatalf("want exactly one %q, got %q", suffix, v)
	}
	return v[0]
}

type templateCase struct {
	tt       TemplateType
	model    string
	idAttr   string // path of the element whose Id holds the key
	idPrefix string
	emit     string // path of the issuer's CNPJ
	required []string
}

var templateCases = []templateCase{
	{
		tt: CFe, model: "59", idAttr: "infCFe@Id", idPrefix: "CFe", emit: "infCFe/emit/CNPJ",
		required: []string{"ide/cUF", "ide/cNF", "ide/nserieSAT", "ide/nCFe", "ide/dEmi", "ide/hEmi",
			"ide/cDV", "prod/vItem", "prod/vProd", "total/vCFe", "dest/CPF"},
	},
	{
		tt: NFe, model: "55", idAttr: "infNFe@Id", idPrefix: "NFe", emit: "infNFe/emit/CNPJ",
		required: []string{"ide/cUF", "ide/cNF", "ide/serie", "ide/nNF", "ide/dhEmi", "ide/cDV",
			"prod/vProd", "ICMSTot/vNF", "ICMSTot/vProd"},
	},
	{
		tt: NFCe, model: "65", idAttr: "infNFe@Id", idPrefix: "NFe", emit: "infNFe/emit/CNPJ",
		required: []string{"ide/cUF", "ide/cNF", "ide/serie", "ide/nNF", "ide/dhEmi", "ide/cDV",
			"prod/vProd", "ICMSTot/vNF", "ICMSTot/vProd", "dest/CPF", "infProt/chNFe", "infProt/dhRecbto"},
	},
	{
		tt: NFeDevolucao, model: "55", idAttr: "infNFe@Id", idPrefix: "NFe", emit: "infNFe/emit/CNPJ",
		required: []string{"ide/cUF", "ide/cNF", "ide/serie", "ide/nNF", "ide/dhEmi", "ide/dhSaiEnt", "ide/cDV",
			"prod/vProd", "ICMSTot/vNF", "dest/CPF", "infProt/chNFe", "infProt/dhRecbto"},
	},
}

func generate(t *testing.T, tt TemplateType, options ...Option) parsedXML {
	t.Helper()
	g, err := NewTemplateGenerator(tt)
	if err != nil {
		t.Fatal(err)
	}
	data, err := g.Generate(options...)
	if err != nil {
		t.Fatalf("Generate: %v", err)
	}
	if bytes.Contains(data, []byte("{%")) || bytes.Contains(data, []byte(omitValue)) {
		t.Fatalf("placeholders left in output:\n%s", data)
	}
	return parseXML(t, data)
}

func TestAccessKeyIsConsistent(t *testing.T) {
	issuedAt := time.Date(2026, 7, 15, 13, 45, 10, 0, brasilia)
	cnpjs := []struct {
		name    string
		options []Option
		want    string // expected issuer CNPJ, "" when random
		alnum   bool
	}{
		{name: "random numeric"},
		{name: "numeric masked", options: []Option{WithCNPJ("11.222.333/0001-81")}, want: "11222333000181"},
		{name: "alphanumeric masked lower case", options: []Option{WithCNPJ("12.abc.345/01de-35")}, want: "12ABC34501DE35", alnum: true},
		{name: "random alphanumeric", options: []Option{WithAlphanumericCNPJ()}, alnum: true},
	}

	for _, tc := range templateCases {
		for _, c := range cnpjs {
			t.Run(tc.tt.String()+"/"+c.name, func(t *testing.T) {
				opts := append([]Option{WithIssuedAt(issuedAt)}, c.options...)
				doc := generate(t, tc.tt, opts...)

				key := strings.TrimPrefix(doc.one(t, tc.idAttr), tc.idPrefix)
				emit := doc.one(t, tc.emit)
				if len(key) != 44 {
					t.Fatalf("key %q has length %d, want 44", key, len(key))
				}
				if !br_documents.ValidateAccessKey(key) {
					t.Errorf("key %q has a wrong check digit", key)
				}
				if !v2.ValidateCNPJ(emit) {
					t.Errorf("issuer CNPJ %q is not valid", emit)
				}
				if c.want != "" && emit != c.want {
					t.Errorf("issuer CNPJ = %q, want %q", emit, c.want)
				}
				if c.alnum && strings.IndexFunc(emit, func(r rune) bool { return r >= 'A' && r <= 'Z' }) < 0 {
					t.Errorf("issuer CNPJ %q should be alphanumeric", emit)
				}

				checks := map[string][2]string{
					"cUF":   {key[0:2], doc.one(t, "ide/cUF")},
					"AAMM":  {key[2:6], "2607"},
					"CNPJ":  {key[6:20], emit},
					"model": {key[20:22], tc.model},
					"mod":   {doc.one(t, "ide/mod"), tc.model},
					"cDV":   {key[43:], doc.one(t, "ide/cDV")},
				}
				if tc.tt == CFe {
					checks["nserieSAT"] = [2]string{key[22:31], doc.one(t, "ide/nserieSAT")}
					checks["nCFe"] = [2]string{key[31:37], doc.one(t, "ide/nCFe")}
					checks["cNF"] = [2]string{key[37:43], doc.one(t, "ide/cNF")}
				} else {
					checks["serie"] = [2]string{key[22:25], pad(t, doc.one(t, "ide/serie"), 3)}
					checks["nNF"] = [2]string{key[25:34], pad(t, doc.one(t, "ide/nNF"), 9)}
					checks["tpEmis"] = [2]string{key[34:35], doc.one(t, "ide/tpEmis")}
					checks["cNF"] = [2]string{key[35:43], doc.one(t, "ide/cNF")}
				}
				for name, c := range checks {
					if c[0] != c[1] {
						t.Errorf("%s: key has %q, document has %q", name, c[0], c[1])
					}
				}
				for _, ch := range doc.all("infProt/chNFe") {
					if ch != key {
						t.Errorf("protNFe chNFe = %q, want %q", ch, key)
					}
				}
				for _, uri := range doc.all("Reference@URI") {
					if uri != "#"+tc.idPrefix+key {
						t.Errorf("Signature Reference URI = %q, want #%s%s", uri, tc.idPrefix, key)
					}
				}
				for _, p := range tc.required {
					for _, v := range doc.all(p) {
						if v == "" {
							t.Errorf("%s is empty", p)
						}
					}
					if len(doc.all(p)) == 0 {
						t.Errorf("%s is missing", p)
					}
				}
			})
		}
	}
}

func pad(t *testing.T, s string, width int) string {
	t.Helper()
	n, err := strconv.Atoi(s)
	if err != nil {
		t.Fatalf("%q is not a number", s)
	}
	return fmt.Sprintf("%0*d", width, n)
}

func TestWithIssuedAt(t *testing.T) {
	// 02:30 UTC on Dec 1st is still Nov 30th in Brasília.
	issuedAt := time.Date(2026, 12, 1, 2, 30, 15, 0, time.UTC)
	tests := []struct {
		tt     TemplateType
		fields map[string]string
	}{
		{CFe, map[string]string{"ide/dEmi": "20261130", "ide/hEmi": "233015"}},
		{NFe, map[string]string{"ide/dhEmi": "2026-11-30T23:30:15-03:00", "ide/dSaiEnt": "2026-11-30"}},
		{NFCe, map[string]string{"ide/dhEmi": "2026-11-30T23:30:15-03:00"}},
		{NFeDevolucao, map[string]string{"ide/dhEmi": "2026-11-30T23:30:15-03:00", "ide/dhSaiEnt": "2026-11-30T23:30:15-03:00"}},
	}
	for _, tc := range tests {
		t.Run(tc.tt.String(), func(t *testing.T) {
			doc := generate(t, tc.tt, WithIssuedAt(issuedAt))
			for path, want := range tc.fields {
				if got := doc.one(t, path); got != want {
					t.Errorf("%s = %q, want %q", path, got, want)
				}
			}
			for _, recv := range doc.all("infProt/dhRecbto") {
				got, err := time.Parse(time.RFC3339, recv)
				if err != nil {
					t.Fatalf("dhRecbto %q: %v", recv, err)
				}
				if d := got.Sub(issuedAt); d <= 0 || d > 2*time.Minute {
					t.Errorf("dhRecbto %s should be shortly after the issue time", recv)
				}
			}
		})
	}
}

func TestDefaultIssueDateIsRecent(t *testing.T) {
	for _, tc := range templateCases[1:] {
		doc := generate(t, tc.tt)
		got, err := time.Parse(time.RFC3339, doc.one(t, "ide/dhEmi"))
		if err != nil {
			t.Fatal(err)
		}
		if age := time.Since(got); age < 0 || age > 31*24*time.Hour {
			t.Errorf("%s: default dhEmi %s is not within the last 30 days", tc.tt, got)
		}
	}
}

func TestNFeRecipientDocument(t *testing.T) {
	tests := []struct {
		name     string
		options  []Option
		wantCPF  string
		wantCNPJ string // "" means a random valid CNPJ is expected when CPF is empty
	}{
		{name: "default is a CNPJ"},
		{name: "WithCPF gives a CPF", options: []Option{WithCPF("529.982.247-25")}, wantCPF: "52998224725"},
		{name: "WithDestCNPJ wins over WithCPF", options: []Option{WithCPF("52998224725"), WithDestCNPJ("12ABC34501DE35")}, wantCNPJ: "12ABC34501DE35"},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			doc := generate(t, NFe, tc.options...)
			cpfs, cnpjs := doc.all("dest/CPF"), doc.all("dest/CNPJ")
			if tc.wantCPF != "" {
				if len(cnpjs) != 0 || len(cpfs) != 1 || cpfs[0] != tc.wantCPF {
					t.Fatalf("dest CPF %q CNPJ %q, want only CPF %q", cpfs, cnpjs, tc.wantCPF)
				}
				return
			}
			if len(cpfs) != 0 || len(cnpjs) != 1 || !v2.ValidateCNPJ(cnpjs[0]) {
				t.Fatalf("dest CPF %q CNPJ %q, want only a valid CNPJ", cpfs, cnpjs)
			}
			if tc.wantCNPJ != "" && cnpjs[0] != tc.wantCNPJ {
				t.Errorf("dest CNPJ = %q, want %q", cnpjs[0], tc.wantCNPJ)
			}
		})
	}
}

func TestWithCNPJSetsOnlyTheIssuer(t *testing.T) {
	const cnpj = "11222333000181"
	for _, tc := range templateCases {
		t.Run(tc.tt.String(), func(t *testing.T) {
			doc := generate(t, tc.tt, WithCNPJ(cnpj))
			if got := doc.one(t, tc.emit); got != cnpj {
				t.Fatalf("issuer CNPJ = %q, want %q", got, cnpj)
			}
			count := 0
			for _, v := range doc.all("CNPJ") {
				if v == cnpj {
					count++
				}
			}
			if count != 1 {
				t.Fatalf("CNPJ %s appears %d times, want only in emit", cnpj, count)
			}
		})
	}
}

func TestWithCNPJRejectsMalformedValues(t *testing.T) {
	for _, bad := range []string{"123", "1234567890123X", "ABCDEFGHIJKLMN"} {
		g := NewNFCeGenerator()
		if _, err := g.Generate(WithCNPJ(bad)); err == nil {
			t.Errorf("WithCNPJ(%q): want an error", bad)
		}
		if _, err := NewNFeGenerator().Generate(WithDestCNPJ(bad)); err == nil {
			t.Errorf("WithDestCNPJ(%q): want an error", bad)
		}
	}
}

func TestTotalsMatchItems(t *testing.T) {
	tests := []struct {
		tt     TemplateType
		item   string
		totals []string
	}{
		{CFe, "prod/vItem", []string{"total/vCFe", "pgto/MP/vMP"}},
		{NFe, "prod/vProd", []string{"ICMSTot/vProd", "ICMSTot/vNF"}},
		{NFCe, "prod/vProd", []string{"ICMSTot/vProd", "ICMSTot/vNF", "detPag/vPag"}},
		{NFeDevolucao, "prod/vProd", []string{"ICMSTot/vProd", "ICMSTot/vNF", "detPag/vPag"}},
	}
	for _, tc := range tests {
		t.Run(tc.tt.String(), func(t *testing.T) {
			for i := 0; i < 20; i++ {
				doc := generate(t, tc.tt)
				item := doc.one(t, tc.item)
				value, err := strconv.ParseFloat(item, 64)
				if err != nil || value <= 0 || value > 10000 {
					t.Fatalf("%s = %q, want a positive amount up to 10000", tc.item, item)
				}
				for _, total := range tc.totals {
					if got := doc.one(t, total); got != item {
						t.Errorf("%s = %q, want %q", total, got, item)
					}
				}
				if got := doc.one(t, "ide/tpAmb"); got != "1" {
					t.Errorf("tpAmb = %q, want 1", got)
				}
			}
		})
	}
}

func TestWithSeedIsReproducible(t *testing.T) {
	issuedAt := time.Date(2026, 11, 20, 9, 0, 0, 0, brasilia)
	for _, tc := range templateCases {
		t.Run(tc.tt.String(), func(t *testing.T) {
			g, _ := NewTemplateGenerator(tc.tt)
			a, _ := g.Generate(WithSeed(42), WithIssuedAt(issuedAt))
			b, _ := g.Generate(WithSeed(42), WithIssuedAt(issuedAt))
			c, _ := g.Generate(WithSeed(43), WithIssuedAt(issuedAt))
			if !bytes.Equal(a, b) {
				t.Error("same seed produced different output")
			}
			if bytes.Equal(a, c) {
				t.Error("different seeds produced the same output")
			}
		})
	}
}
