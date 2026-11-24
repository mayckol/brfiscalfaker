package nfs

import (
	"fmt"
	"math"
	"regexp"
	"strings"
	"time"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/mayckol/brfiscalfaker/pkg/br_documents"
	v2 "github.com/mayckol/brfiscalfaker/pkg/br_documents/v2"
)

// brasilia is the fixed UTC-03:00 offset used for every fiscal date
// (Brazil has had no daylight saving time since 2019).
var brasilia = time.FixedZone("-03", -3*60*60)

// omitValue marks an element that must be removed from the output, such as
// the unused CNPJ/CPF choice of the NF-e recipient.
const omitValue = "__brfiscalfaker_omit__"

var omitRe = regexp.MustCompile(`\s*<\w+>` + omitValue + `</\w+>`)

var modelRe = regexp.MustCompile(`<mod>(\d{2})</mod>`)

// ufByCode maps IBGE state codes (cUF) to state abbreviations.
var ufByCode = map[string]string{
	"11": "RO", "12": "AC", "13": "AM", "14": "RR", "15": "PA", "16": "AP", "17": "TO",
	"21": "MA", "22": "PI", "23": "CE", "24": "RN", "25": "PB", "26": "PE", "27": "AL",
	"28": "SE", "29": "BA", "31": "MG", "32": "ES", "33": "RJ", "35": "SP",
	"41": "PR", "42": "SC", "43": "RS", "50": "MS", "51": "MT", "52": "GO", "53": "DF",
}

var ufCodes = []string{
	"11", "12", "13", "14", "15", "16", "17", "21", "22", "23", "24", "25", "26", "27",
	"28", "29", "31", "32", "33", "35", "41", "42", "43", "50", "51", "52", "53",
}

// alphanumericCNPJChars is the generation charset of alphanumeric CNPJs,
// without the ambiguous letters I, O, Q, U and F.
const alphanumericCNPJChars = "0123456789ABCDEGHJKLMNPRSTVWXYZ"

// buildInvoiceValues computes the placeholders that must agree with each other
// (issue date, access key, identifiers, item values and totals).
func buildInvoiceValues(template string, cfg *generationConfig) (map[string]string, error) {
	emitCNPJ, err := issuerCNPJ(cfg)
	if err != nil {
		return nil, err
	}
	values := map[string]string{}
	if err := addRecipient(values, cfg); err != nil {
		return nil, err
	}

	issuedAt := issueTime(cfg)
	cUF := gofakeit.RandomString(ufCodes)
	addIdentification(values, template, cUF, emitCNPJ, issuedAt)
	addDates(values, issuedAt)
	addAmounts(values)
	addFixedFields(values)
	return values, nil
}

func issuerCNPJ(cfg *generationConfig) (string, error) {
	if cfg.CNPJ != "" {
		return normalizedCNPJ(cfg.CNPJ)
	}
	if cfg.alphanumericCNPJ {
		return fakeAlphanumericCNPJ(), nil
	}
	return fakeCNPJ(), nil
}

func normalizedCNPJ(cnpj string) (string, error) {
	normalized := v2.Normalize(cnpj)
	if !v2.HasValidFormat(normalized) {
		return "", fmt.Errorf("invalid CNPJ %q: want 12 alphanumeric characters and 2 check digits", cnpj)
	}
	return normalized, nil
}

// addRecipient fills the recipient's document. NF-e has both a CNPJ and a CPF
// element; the one not used is removed after replacement.
func addRecipient(values map[string]string, cfg *generationConfig) error {
	cpf := digitsOnly(cfg.CPF)
	if cpf == "" {
		cpf = fakeCPF()
	}
	values["CPF"] = cpf

	switch {
	case cfg.destCNPJ != "":
		cnpj, err := normalizedCNPJ(cfg.destCNPJ)
		if err != nil {
			return err
		}
		values["destCNPJ"], values["destCPF"] = cnpj, omitValue
	case cfg.CPF != "":
		values["destCNPJ"], values["destCPF"] = omitValue, cpf
	default:
		values["destCNPJ"], values["destCPF"] = fakeCNPJ(), omitValue
	}
	return nil
}

func issueTime(cfg *generationConfig) time.Time {
	t := cfg.issuedAt
	if t.IsZero() {
		back := time.Duration(gofakeit.Number(60, 30*24*60*60)) * time.Second
		t = time.Now().Add(-back)
	}
	return t.In(brasilia).Truncate(time.Second)
}

// addIdentification fills the access key and the fields it is made of.
func addIdentification(values map[string]string, template, cUF, emitCNPJ string, issuedAt time.Time) {
	model := "55"
	if m := modelRe.FindStringSubmatch(template); m != nil {
		model = m[1]
	}
	aamm := issuedAt.Format("0601")

	var partial string
	if model == "59" {
		nserieSAT := gofakeit.Number(1, 999999999)
		nCFe := gofakeit.Number(1, 999999)
		cNF := gofakeit.Number(0, 999999)
		partial = fmt.Sprintf("%s%s%s59%09d%06d%06d", cUF, aamm, emitCNPJ, nserieSAT, nCFe, cNF)
		values["nserieSAT"] = fmt.Sprintf("%09d", nserieSAT)
		values["nCFe"] = fmt.Sprintf("%06d", nCFe)
		values["cNF"] = fmt.Sprintf("%06d", cNF)
	} else {
		serie := gofakeit.Number(1, 999)
		nNF := gofakeit.Number(1, 999999)
		cNF := gofakeit.Number(10000000, 99999999)
		partial = fmt.Sprintf("%s%s%s%s%03d%09d1%08d", cUF, aamm, emitCNPJ, model, serie, nNF, cNF)
		values["serie"] = fmt.Sprintf("%d", serie)
		values["nNF"] = fmt.Sprintf("%d", nNF)
		values["cNF"] = fmt.Sprintf("%08d", cNF)
	}

	dv := br_documents.AccessKeyDV(partial)
	key := fmt.Sprintf("%s%d", partial, dv)
	uf := ufByCode[cUF]

	values["accessKey"] = key
	values["chNFe"] = key
	values["cDV"] = fmt.Sprintf("%d", dv)
	values["cUF"] = cUF
	values["UF"] = uf
	values["enderEmitUF"] = uf
	values["emitCNPJ"] = emitCNPJ
	values["nProt"] = fmt.Sprintf("1%s%s%s", cUF, issuedAt.Format("06"), gofakeit.Numerify("##########"))
	values["qrCode"] = fmt.Sprintf("https://www.fazenda.%s.gov.br/nfce/qrcode?p=%s|2|1|1|%s",
		strings.ToLower(uf), key, gofakeit.Numerify("########################################"))
}

func addDates(values map[string]string, issuedAt time.Time) {
	const layout = "2006-01-02T15:04:05-07:00"
	received := issuedAt.Add(time.Duration(gofakeit.Number(1, 120)) * time.Second)

	values["dhEmi"] = issuedAt.Format(layout)
	values["dhSaiEnt"] = issuedAt.Format(layout)
	values["dSaiEnt"] = issuedAt.Format("2006-01-02")
	values["dhRecbto"] = received.Format(layout)
	values["dEmi"] = issuedAt.Format("20060102")
	values["hEmi"] = issuedAt.Format("150405")
}

// addAmounts fills one item and totals that add up: the total is the item's
// value, and taxes are computed from it.
func addAmounts(values map[string]string) {
	qty := gofakeit.Number(1, 5)
	unitCents := gofakeit.Number(100, 50000)
	prodCents := qty * unitCents
	icms := percentOf(prodCents, 18)
	pis := percentOf(prodCents, 1.65)
	cofins := percentOf(prodCents, 7.6)
	approxTaxes := percentOf(prodCents, 30)
	unit := gofakeit.RandomString([]string{"UN", "PC", "KG", "CX", "LT"})

	set(values, fmt.Sprintf("%d.0000", qty), "qCom", "qTrib", "detProdQCom", "detProdQTrib")
	set(values, money(unitCents), "vUnCom", "vUnTrib", "detProdVUnCom", "detProdVUnTrib")
	set(values, unit, "uCom", "uTrib", "detProdUCom", "detProdUTrib")
	set(values, money(prodCents),
		"vProd", "detProdVProd", "vProd_total", "totalICMSTotvProd",
		"vItem", "vCFe", "vNF", "totalICMSTotvNF", "vPag", "vMP",
		"vBC_PIS", "vBC_COFINS", "impostoPISAliqvBC", "impostoCOFINSAliqvBC",
		"impostoICMS00vBC", "totalICMSTotvBC", "vBC_IPI")
	set(values, money(icms), "vICMS", "impostoICMS00vICMS", "totalICMSTotvICMS")
	set(values, money(pis), "vPIS", "impostoPISAliqvPIS", "vPIS_total", "totalICMSTotvPIS")
	set(values, money(cofins), "vCOFINS", "impostoCOFINSAliqvCOFINS", "vCOFINS_total", "totalICMSTotvCOFINS")
	set(values, money(approxTaxes), "vTotTrib", "vTotTrib_total", "vItem12741", "vCFeLei12741")
	set(values, "18.00", "pICMS", "impostoICMS00pICMS")
	set(values, "1.65", "pPIS", "impostoPISAliqpPIS")
	set(values, "7.60", "pCOFINS", "impostoCOFINSAliqpCOFINS")
	set(values, "0.00",
		"vDesc", "vDesc_total", "totalICMSTotvDesc", "vBC_total", "vICMS_total",
		"vICMSDeson", "vFCP", "vBCST", "vST", "vFCPST", "vFCPSTRet", "vFrete", "vSeg",
		"vII", "vIPI", "vIPIDevol", "vOutro", "vFCPUFDest", "vICMSUFDest", "vICMSUFRemet",
		"vIPI_total", "vIPIDevol_total", "vPISST", "vCOFINSST", "vTroco", "pIPI",
		"totalICMSTotvBCST", "totalICMSTotvST", "totalICMSTotvFrete", "totalICMSTotvSeg",
		"totalICMSTotvII", "totalICMSTotvIPI", "totalICMSTotvOutro")
}

// addFixedFields fills codes that must hold a valid, stable value and the
// fields that had no generator before.
func addFixedFields(values map[string]string) {
	set(values, "1", "tpEmis", "tpAmb", "tpAmbProt", "indTot", "qVol", "nItem", "detNItem")
	values["cStat"] = "100"
	values["xMotivo"] = "Autorizado o uso da NF-e"
	values["CST_ICMS"] = "00"
	values["Orig"] = "0"
	values["CST_IPI"] = "99"
	values["cEnq"] = "999"
	values["CEST"] = gofakeit.Numerify("#######")
	values["pDevol"] = "100.00"
	values["infAdFisco"] = gofakeit.Sentence(6)
	values["infCpl"] = gofakeit.Sentence(6)
	values["infRespTecCNPJ"] = fakeCNPJ()
	values["infRespTecXContato"] = gofakeit.Name()
	values["infRespTecEmail"] = gofakeit.Email()
	values["infRespTecFone"] = gofakeit.Numerify("###########")

	// CF-e (SAT) specific fields.
	values["acCNPJ"] = fakeCNPJ()
	values["signAC"] = gofakeit.LetterN(64)
	values["assinaturaQRCODE"] = gofakeit.LetterN(64)
	values["numeroCaixa"] = fmt.Sprintf("%03d", gofakeit.Number(1, 999))
	values["cRegTrib"] = "3"
	values["indRatISSQN"] = "N"
	values["indRegra"] = "A"
	values["cMP"] = "01"
	values["cAdmC"] = "999"
	values["xCampo"] = "ICMS"
	values["xTexto"] = "Comete crime quem sonega"
}

func set(values map[string]string, value string, keys ...string) {
	for _, k := range keys {
		values[k] = value
	}
}

func percentOf(cents int, percent float64) int {
	return int(math.Round(float64(cents) * percent / 100))
}

func money(cents int) string {
	return fmt.Sprintf("%d.%02d", cents/100, cents%100)
}

func digitsOnly(s string) string {
	var b strings.Builder
	for _, r := range s {
		if r >= '0' && r <= '9' {
			b.WriteRune(r)
		}
	}
	return b.String()
}

// fakeCNPJ returns a valid numeric CNPJ drawn from gofakeit, so WithSeed
// makes it reproducible.
func fakeCNPJ() string {
	return v2.CompleteCNPJ(gofakeit.Numerify("############"))
}

// fakeAlphanumericCNPJ returns a valid alphanumeric CNPJ (July 2026 format).
func fakeAlphanumericCNPJ() string {
	base := make([]byte, 12)
	for i := range base {
		base[i] = alphanumericCNPJChars[gofakeit.Number(0, len(alphanumericCNPJChars)-1)]
	}
	return v2.CompleteCNPJ(string(base))
}

// fakeCPF returns a valid CPF drawn from gofakeit.
func fakeCPF() string {
	return br_documents.CompleteCPF(gofakeit.Numerify("#########"))
}
