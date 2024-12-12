package brfiscalfaker

// NFeGenerator generates a standard NFe XML.
type NFeGenerator struct {
	template string
}

// NewNFeGenerator creates a new instance of NFeGenerator with the NFe XML template.
func NewNFeGenerator() *NFeGenerator {
	return &NFeGenerator{
		template: NFeXMLMock,
	}
}

// Generate replaces placeholders in the NFe template, respecting blocked placeholders.
func (g *NFeGenerator) Generate(options ...Option) ([]byte, error) {
	return ReplaceTemplate(g.template, options...)
}

// NFCeGenerator generates a NFCe XML.
type NFCeGenerator struct {
	template string
}

// NewNFCeGenerator creates a new instance of NFCeGenerator with the NFCe XML template.
func NewNFCeGenerator() *NFCeGenerator {
	return &NFCeGenerator{
		template: NFCeXMLMock, // Define NFCeXMLMock similarly to CFeXMLMock
	}
}

type CFeGenerator struct {
	template string
}

func (c CFeGenerator) Generate(options ...Option) ([]byte, error) {
	return ReplaceTemplate(c.template, options...)
}

// NewCFeGenerator creates a new instance of CFeGenerator with the CFe XML template.
func NewCFeGenerator() *CFeGenerator {
	return &CFeGenerator{
		template: CFeXMLMock,
	}
}

// Generate replaces placeholders in the NFCe template, respecting blocked placeholders.
func (g *NFCeGenerator) Generate(options ...Option) ([]byte, error) {
	return ReplaceTemplate(g.template, options...)
}

// NFeDevolucaoGenerator generates an NFe Devolucao XML.
type NFeDevolucaoGenerator struct {
	template string
}

// NewNFeDevolucaoGenerator creates a new instance of NFeDevolucaoGenerator with the NFeDevolucao XML template.
func NewNFeDevolucaoGenerator() *NFeDevolucaoGenerator {
	return &NFeDevolucaoGenerator{
		template: NFeDevolucaoXMLMock,
	}
}

// Generate replaces placeholders in the NFeDevolucao template, respecting blocked placeholders.
func (g *NFeDevolucaoGenerator) Generate(options ...Option) ([]byte, error) {
	return ReplaceTemplate(g.template, options...)
}
