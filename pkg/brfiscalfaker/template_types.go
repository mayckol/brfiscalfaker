package brfiscalfaker

import "fmt"

// TemplateType is a custom type representing the type of XML template.
type TemplateType int

const (
	// Define the enum values using iota for automatic incrementing.
	CFe TemplateType = iota
	NFe
	NFCe
	NFeDevolucao
)

// String returns the string representation of the TemplateType.
func (tt TemplateType) String() string {
	switch tt {
	case CFe:
		return "CFe"
	case NFe:
		return "NFe"
	case NFCe:
		return "NFCe"
	case NFeDevolucao:
		return "NFeDevolucao"
	default:
		return "Unknown"
	}
}

// ParseTemplateType converts a string to a TemplateType enum.
// It returns an error if the input string does not match any known template type.
func ParseTemplateType(s string) (TemplateType, error) {
	switch s {
	case "CFe":
		return CFe, nil
	case "NFe":
		return NFe, nil
	case "NFCe":
		return NFCe, nil
	case "NFeDevolucao":
		return NFeDevolucao, nil
	default:
		return -1, fmt.Errorf("invalid TemplateType: %s", s)
	}
}
