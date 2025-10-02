// Description: Command-line tool to generate fake Brazilian fiscal invoices.
package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/mayckol/brfiscalfaker/pkg/nfs"

	"golang.org/x/term"
)

func main() {
	cpf := flag.String("cpf", "", "Optional CPF to include in the invoice")
	cnpj := flag.String("cnpj", "", "Optional CNPJ to include in the invoice")
	templateType := flag.String("type", "NFCe", "Type of invoice to generate (CFe, NFe, NFCe, NFeDevolucao)")
	blockTags := flag.String("block-tags", "", "Comma-separated list of placeholders to block (e.g., emitCNPJ,CNPJ,CPF)")

	flag.Parse()

	var tt nfs.TemplateType
	switch *templateType {
	case "CFe":
		tt = nfs.CFe
	case "NFe":
		tt = nfs.NFe
	case "NFCe":
		tt = nfs.NFCe
	case "NFeDevolucao":
		tt = nfs.NFeDevolucao
	default:
		log.Fatalf("Unsupported template type: %s", *templateType)
	}

	// Create the generator
	generator, err := nfs.NewTemplateGenerator(tt)
	if err != nil {
		log.Fatalf("Failed to create generator: %v", err)
	}

	// Prepare options
	var options []nfs.Option
	if *cpf != "" {
		option := nfs.WithCPF(*cpf)
		options = append(options, option)
	}
	if *cnpj != "" {
		option := nfs.WithCNPJ(*cnpj)
		options = append(options, option)
	}
	if *blockTags != "" {
		// Split the blockTags by comma and trim any whitespace
		tags := splitAndTrim(*blockTags, ",")
		if len(tags) > 0 {
			options = append(options, nfs.WithBlockedPlaceholders(tags...))
		}
	}

	// Generate the invoice
	xmlBytes, err := generator.Generate(options...)
	if err != nil {
		log.Fatalf("Failed to generate invoice: %v", err)
	}

	// Determine if running via Docker by checking if stdout is a terminal
	if isTerminal(os.Stdout) {
		// Print the XML
		fmt.Println(string(xmlBytes))
	} else {
		// If not a terminal (e.g., piped), write the bytes to stdout
		os.Stdout.Write(xmlBytes)
	}
}

// isTerminal checks if the file descriptor is a terminal.
func isTerminal(f *os.File) bool {
	return term.IsTerminal(int(f.Fd()))
}

// splitAndTrim splits a string by a separator and trims whitespace from each element.
func splitAndTrim(s string, sep string) []string {
	raw := strings.Split(s, sep)
	var result []string
	for _, item := range raw {
		trimmed := strings.TrimSpace(item)
		if trimmed != "" {
			result = append(result, trimmed)
		}
	}
	return result
}
