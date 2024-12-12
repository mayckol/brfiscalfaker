package brfiscalfaker

import (
	"strings"
	"testing"
)

func TestNFeGenerator_Generate_NoBlocking(t *testing.T) {
	generator := NewNFeGenerator()

	xmlBytes, err := generator.Generate()
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	xmlContent := string(xmlBytes)

	// Ensure no placeholders remain
	if strings.Contains(xmlContent, "{%") {
		t.Errorf("Expected all placeholders to be replaced, but some remain")
	}

	// Optionally, perform more checks like presence of certain tags or structure
	if !strings.Contains(xmlContent, "<infNFe") || !strings.Contains(xmlContent, "</infNFe>") {
		t.Errorf("Expected <infNFe> tags to be present")
	}
}

func TestNFeGenerator_Generate_BlockSinglePlaceholder(t *testing.T) {
	generator := NewNFeGenerator()

	// Block "CNPJ"
	xmlBytes, err := generator.Generate(WithBlockedPlaceholders([]string{"CNPJ"}...))
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	xmlContent := string(xmlBytes)

	// Check that <CNPJ> tags are removed
	if strings.Contains(xmlContent, "<CNPJ>") || strings.Contains(xmlContent, "</CNPJ>") {
		t.Errorf("Expected <CNPJ> tags to be removed")
	}

	// Ensure other placeholders are replaced
	if strings.Contains(xmlContent, "{%") {
		t.Errorf("Expected all non-blocked placeholders to be replaced")
	}

	// Ensure no blank lines are left
	lines := strings.Split(xmlContent, "\n")
	for i, line := range lines {
		if strings.TrimSpace(line) == "" && i != len(lines)-1 {
			t.Errorf("Found blank line at line %d", i+1)
		}
	}
}

func TestNFeGenerator_Generate_BlockMultiplePlaceholders(t *testing.T) {
	generator := NewNFeGenerator()

	// Block "CNPJ", "IE", and "transporta"
	xmlBytes, err := generator.Generate(WithBlockedPlaceholders([]string{"CNPJ", "IE", "transporta"}...))
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	xmlContent := string(xmlBytes)

	// Check that <CNPJ>, <IE>, and <transporta> tags are removed
	if strings.Contains(xmlContent, "<CNPJ>") || strings.Contains(xmlContent, "</CNPJ>") ||
		strings.Contains(xmlContent, "<IE>") || strings.Contains(xmlContent, "</IE>") ||
		strings.Contains(xmlContent, "<transporta>") || strings.Contains(xmlContent, "</transporta>") {
		t.Errorf("Expected <CNPJ>, <IE>, and <transporta> tags to be removed")
	}

	// Ensure other placeholders are replaced
	if strings.Contains(xmlContent, "{%") {
		t.Errorf("Expected all non-blocked placeholders to be replaced")
	}

	// Ensure no blank lines are left
	lines := strings.Split(xmlContent, "\n")
	for i, line := range lines {
		if strings.TrimSpace(line) == "" && i != len(lines)-1 {
			t.Errorf("Found blank line at line %d", i+1)
		}
	}
}

func TestNFeGenerator_Generate_BlockNonExistentPlaceholder(t *testing.T) {
	generator := NewNFeGenerator()

	// Block a non-existent placeholder "NonExistent"
	xmlBytes, err := generator.Generate(WithBlockedPlaceholders([]string{"NonExistent"}...))
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	xmlContent := string(xmlBytes)

	// Ensure existing placeholders are replaced
	if strings.Contains(xmlContent, "{%") {
		t.Errorf("Expected all existing placeholders to be replaced")
	}

	// Ensure non-specified placeholders are unaffected
	// (Since "NonExistent" doesn't exist, the template should remain unchanged except for placeholder replacement)

	// Ensure no blank lines are left
	lines := strings.Split(xmlContent, "\n")
	for i, line := range lines {
		if strings.TrimSpace(line) == "" && i != len(lines)-1 {
			t.Errorf("Found blank line at line %d", i+1)
		}
	}
}

func TestNFeGenerator_Generate_AllPlaceholdersBlocked(t *testing.T) {
	generator := NewNFeGenerator()

	// Block all possible placeholders by listing them (for the sake of the test, assume some)
	blockedPlaceholders := []string{
		"cUF", "cNF", "natOp", "indPag", "serie", "nNF",
		"dEmi", "dSaiEnt", "tpNF", "cMunFG", "tpImp", "tpEmis",
		"cDV", "tpAmb", "finNFe", "procEmi", "verProc",
		"emitCNPJ", "emitXNome", "emitXFant", "enderEmitXLgr",
		"enderEmitNro", "enderEmitXCpl", "enderEmitXBairro", "enderEmitCMun",
		"enderEmitXMun", "enderEmitUF", "enderEmitCEP", "enderEmitCPais",
		"enderEmitXPais", "enderEmitFone", "emitIE",
		"destCNPJ", "destXNome", "enderDestXLgr", "enderDestNro",
		"enderDestXCpl", "enderDestXBairro", "enderDestCMun",
		"enderDestXMun", "enderDestUF", "enderDestCEP", "enderDestCPais",
		"enderDestXPais", "enderDestFone", "destIE",
		"retiradaCNPJ", "retiradaXLgr", "retiradaNro", "retiradaXCpl",
		"retiradaXBairro", "retiradaCMun", "retiradaXMun", "retiradaUF",
		"entregaCNPJ", "entregaXLgr", "entregaNro", "entregaXCpl",
		"entregaXBairro", "entregaCMun", "entregaXMun", "entregaUF",
		"detNItem", "detProdCProd", "detProdCEAN", "detProdXProd",
		"detProdCFOP", "detProdUCom", "detProdQCom", "detProdVUnCom",
		"detProdVProd", "detProdCEANTrib", "detProdUTrib", "detProdQTrib",
		"detProdVUnTrib", "impostoICMS00orig", "impostoICMS00CST",
		"impostoICMS00modBC", "impostoICMS00vBC", "impostoICMS00pICMS",
		"impostoICMS00vICMS", "impostoPISAliqCST", "impostoPISAliqvBC",
		"impostoPISAliqpPIS", "impostoPISAliqvPIS", "impostoCOFINSAliqCST",
		"impostoCOFINSAliqvBC", "impostoCOFINSAliqpCOFINS",
		"impostoCOFINSAliqvCOFINS", "totalICMSTotvBC", "totalICMSTotvICMS",
		"totalICMSTotvBCST", "totalICMSTotvST", "totalICMSTotvProd",
		"totalICMSTotvFrete", "totalICMSTotvSeg", "totalICMSTotvDesc",
		"totalICMSTotvII", "totalICMSTotvIPI", "totalICMSTotvPIS",
		"totalICMSTotvCOFINS", "totalICMSTotvOutro", "totalICMSTotvNF",
		"transpModFrete", "transpTransportaCNPJ", "transpTransportaXNome",
		"transpTransportaIE", "transpTransportaXEnder", "transpTransportaXMun",
		"transpTransportaUF", "transpVeicTranspPlaca", "transpVeicTranspUF",
		"transpVeicTranspRNTC", "transpReboquePlaca", "transpReboqueUF",
		"transpReboqueRNTC", "transpVolQVol", "transpVolEsp", "transpVolMarca",
		"transpVolNVol", "transpVolPesoL", "transpVolPesoB", "transpVolLacresNLacre",
		"infAdicInfAdFisco", "DigestValue", "SignatureValue", "X509Certificate",
	}

	xmlBytes, err := generator.Generate(WithBlockedPlaceholders(blockedPlaceholders...))
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	xmlContent := string(xmlBytes)

	// Ensure all specified placeholders are removed
	for _, placeholder := range blockedPlaceholders {
		if strings.Contains(xmlContent, "<"+placeholder+">") || strings.Contains(xmlContent, "</"+placeholder+">") {
			t.Errorf("Expected <%s> tags to be removed", placeholder)
		}
	}

	// Ensure no placeholders remain
	if strings.Contains(xmlContent, "{%") {
		t.Errorf("Expected all placeholders to be replaced or removed")
	}

	// Ensure no blank lines are left
	lines := strings.Split(xmlContent, "\n")
	for i, line := range lines {
		if strings.TrimSpace(line) == "" && i != len(lines)-1 {
			t.Errorf("Found blank line at line %d", i+1)
		}
	}
}

func TestNFeGenerator_Generate_InvalidOption(t *testing.T) {
	generator := NewNFeGenerator()

	// Pass an empty string as a blocked placeholder (assuming it's allowed and simply does nothing)
	xmlBytes, err := generator.Generate(WithBlockedPlaceholders(""))
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	xmlContent := string(xmlBytes)

	// Ensure that the generator still works correctly
	if strings.Contains(xmlContent, "{%") {
		t.Errorf("Expected all placeholders to be replaced")
	}
}

func TestNFeGenerator_Generate_EmptyTemplate(t *testing.T) {
	// Create a generator with an empty template
	emptyGenerator := &NFeGenerator{template: ""}

	xmlBytes, err := emptyGenerator.Generate()
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	xmlContent := string(xmlBytes)

	// The result should be empty
	if xmlContent != "" {
		t.Errorf("Expected empty string, got %s", xmlContent)
	}
}
