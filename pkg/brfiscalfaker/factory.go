package brfiscalfaker

import (
	"fmt"
)

// TemplateGenerator is an interface for generating XML templates.
type TemplateGenerator interface {
	Generate(options ...Option) ([]byte, error)
}

// NewTemplateGenerator creates a TemplateGenerator based on the provided TemplateType.
// It returns an error if the TemplateType is unsupported.
func NewTemplateGenerator(templateType TemplateType) (TemplateGenerator, error) {
	switch templateType {
	case CFe:
		return NewCFeGenerator(), nil
	case NFe:
		return NewNFeGenerator(), nil
	case NFCe:
		return NewNFCeGenerator(), nil
	case NFeDevolucao:
		return NewNFeDevolucaoGenerator(), nil
	default:
		return nil, fmt.Errorf("unsupported template type: %v", templateType)
	}
}
