<div align="center">
  <img src="logo.svg" alt="BrFiscalFaker Logo" width="400" height="400">
  <h1>BrFiscalFaker</h1>
</div>

## Overview

brfiscalfaker is a Go-based command-line tool designed to generate mock Brazilian fiscal invoices (NF-e, NFC-e, CFe, NFeDevolucao) for testing and development purposes. It allows users to create realistic invoice XML files with customizable data, facilitating the development of applications that interact with Brazilian fiscal systems.

## Features

- **Supports Multiple Invoice Types:** Generate NF-e, NFC-e, CFe, and NFeDevolucao invoices.
- **Customizable Data:** Provide custom CPF and CNPJ numbers.
- **Block Specific Tags:** Remove or block specific XML tags using the `--block-tags` flag.
- **Dependency Management:** Ensures dependent placeholders are processed in the correct order.
- **Cross-Platform:** Works seamlessly on various operating systems.
- **Comprehensive Logging:** Provides detailed logs for debugging and transparency.
- **Br documents:** This project includes utilities to generate random yet valid Brazilian fiscal identifiers such as Access Key (Chave de Acesso), CPF, and CNPJ. These are essential for creating mock data for testing purposes.
- **Unit Tested:** Robust unit tests to ensure reliability.

## Installation

Ensure you have Go installed on your system (version 1.16 or higher).

1. **Clone the Repository:**

   ```bash
   git clone https://github.com/mayckol/brfiscalfaker.git
   cd brfiscalfaker
    ```
   
2. **Build the Application:**

   ```bash
   go build -o bfiscalfaker ./cmd/bfiscalfaker/main.go
   ```
   
## Usage

Run the application using the `go run command or the built binary.

### Command-Line Flags

- **`--cpf` (`optional`):** Recipient CPF. In NF-e it makes the recipient a CPF instead of a CNPJ.
- **`--cnpj` (`optional`):** Issuer (emit) CNPJ, numeric or alphanumeric, raw or masked. It is also embedded in the access key.
- **`--dest-cnpj` (`optional`):** Recipient CNPJ (NF-e only).
- **`--alnum-cnpj` (`optional`):** Use a random alphanumeric issuer CNPJ (July 2026 format).
- **`--issued-at` (`optional`):** Issue date and time in RFC 3339 (e.g. `2026-11-24T10:30:00-03:00`).
- **`--seed` (`optional`):** Non-zero seed for reproducible output.
- **`--block-tags` (`optional`):** --block-tags: (Optional) Block specific XML tags from being included in the invoice.
- **`--type` (`default NFCe`):** --type: (Optional) Specify the type of invoice to generate (NF-e, NFC-e, CFe, NFeDevolucao).

### Examples
* **Generate a Basic NFC-e Invoice:**

   ```bash
   go run cmd/bfiscalfaker/main.go
   ```
* **Generate an NF-e Invoice with Custom CPF and CNPJ:**

   ```bash
    go run cmd/bfiscalfaker/main.go --type NFe --cpf 12345678900 --cnpj 12345678901234
    ```
* **Generate a CFe Invoice with Blocked Tags:**

   ```bash
   go run cmd/bfiscalfaker/main.go --type CFe --block-tags "nItem, vProd"
   ```
* **Generate an NFCe Invoice with Custom Data and Blocked Tags:**

   ```bash
   go run cmd/bfiscalfaker/main.go --type NFCe --cpf 12345678900 --cnpj 12345678901234 --block-tags "nItem, vProd"
   ```

## Library Usage

### Download the Library

```bash
go get github.com/mayckol/brfiscalfaker
```

### Generation Options

| Option | What it does |
| --- | --- |
| `WithCNPJ(cnpj string)` | Issuer (emit) CNPJ, also embedded in the access key. Numeric or alphanumeric, raw or masked (normalized to 14 upper-case characters). `Generate` returns an error if it does not have the CNPJ layout. |
| `WithAlphanumericCNPJ()` | Random valid alphanumeric issuer CNPJ (July 2026 format). `WithCNPJ` wins over it. |
| `WithCPF(cpf string)` | Recipient CPF (mask accepted). In NF-e it switches `<dest>` from `<CNPJ>` to `<CPF>`. |
| `WithDestCNPJ(cnpj string)` | Recipient CNPJ in NF-e (numeric or alphanumeric). Wins over `WithCPF`. |
| `WithIssuedAt(t time.Time)` | Issue date/time: `dhEmi`, `dhSaiEnt`/`dSaiEnt`, `dhRecbto` (a few seconds later), CF-e `dEmi`/`hEmi` and the key's AAMM. Written in Brasília time (UTC-03:00). Default: a random moment in the last 30 days. |
| `WithSeed(seed int64)` | Reproducible output (same seed and options give the same XML). Combine with `WithIssuedAt`, since the default date depends on the current time. `0` means random. |
| `WithBlockedPlaceholders(tags ...string)` | Removes the given XML elements from the output. |

Every generated document is internally consistent:

- The access key has the right layout and a valid mod-11 check digit, and it matches `cUF`, the issue month, the issuer CNPJ, the model (55 NF-e, 65 NFC-e, 59 CF-e), `serie`/`nNF`/`tpEmis`/`cNF` (or `nserieSAT`/`nCFe`/`cNF` for CF-e) and `cDV`. The `Id` attribute, the signature reference and `protNFe/infProt/chNFe` all carry the same key.
- `tpEmis` and `tpAmb` are `1`; `cStat` is `100`.
- One item per document, with `vProd = qCom × vUnCom`; totals (`vProd`, `vNF`, `vPag`, CF-e `vItem`/`vCFe`/`vMP`) equal the item value, and taxes are computed from it.

```go
xmlBytes, err := generator.Generate(
    nfs.WithIssuedAt(time.Now().Add(-2 * time.Hour)),
    nfs.WithCNPJ("12.ABC.345/01DE-35"), // alphanumeric issuer CNPJ
    nfs.WithCPF("529.982.247-25"),      // NF-e recipient becomes a CPF
    nfs.WithSeed(42),
)
```

#### Alphanumeric CNPJ in invoices

From July 2026 a CNPJ may hold letters in its first 12 positions. `WithCNPJ` and `WithDestCNPJ` accept them, and
`WithAlphanumericCNPJ()` draws a random one for the issuer. The access key embeds the alphanumeric CNPJ at positions
7-20, and its check digit follows the official rule: modulo 11 with every character valued as its ASCII code minus 48
(digits keep their value, so numeric keys are unchanged). `br_documents.AccessKeyDV` and `br_documents.ValidateAccessKey`
expose that rule.

### Create a Template Generator
* **Random CNPJ, CPF, and Access Key Generation:**

   ```go
   package main
    import (
        "fmt"
        "github.com/mayckol/brfiscalfaker/pkg/br_documents"
    )
    func main() {
        cpf := br_documents.CPF()
        cnpj := br_documents.CNPJ()
        accessKey := br_documents.AccessKey()
        fmt.Println("CPF:", cpf)
        fmt.Println("CNPJ:", cnpj)
        fmt.Println("Access Key:", accessKey)
    }
    ```  
**Parameters Description:**

- **CPF()**
   - **Config**: Optional parameter to control output format.
      - `Masked`: `bool`. If `true`, the CPF is formatted as `XXX.XXX.XXX-XX`. Default is `false` (raw digits).

- **CNPJ()**
   - **Config**: Optional parameter to control output format.
      - `Masked`: `bool`. If `true`, the CNPJ is formatted as `XX.XXX.XXX/XXXX-XX`. Default is `false` (raw digits).

- **AccessKey()**
   - **Config**: Optional parameter to customize the generated key.
      - `Masked`: `bool`. If `true`, the Access Key is formatted with separators for readability. Default is `false` (raw digits).
      - `CNPJ`: `string`. Issuer CNPJ, numeric or alphanumeric (raw or masked). Default is a random numeric CNPJ.
      - `Model`, `UF`, `IssuedAt`: optional model (default `55`), UF code (default random) and issue date (default now).

```go
package main

import (
   "fmt"
   "log"
   "time"

   "github.com/mayckol/brfiscalfaker/pkg/nfs"
)

func main() {
   // Define the desired template type
   templateType := nfs.NFCe

   // Create the generator using the helper
   generator, err := nfs.NewTemplateGenerator(templateType)
   if err != nil {
      log.Fatalf("Failed to create generator: %v", err)
   }
   options := []nfs.Option{
	   nfs.WithCPF("01234567890"),
	   nfs.WithCNPJ("11222333000181"),
	   nfs.WithIssuedAt(time.Now()),
   }

   // Generate the XML template
   xmlBytes, err := generator.Generate(options...)
   if err != nil {
      log.Fatalf("Failed to generate XML template: %v", err)
   }

   fmt.Println(string(xmlBytes))

}

```

### Alphanumeric CNPJ (v2) — July 2026 Format

Brazil's new alphanumeric CNPJ format becomes effective in July 2026. This package includes a v2 module with full support for the new Módulo 11 algorithm with dual check digits.

**v2 Features:**
- 14-character alphanumeric CNPJ format (8 root + 4 branch + 2 numeric check digits)
- Automatic check digit calculation using Módulo 11 with specified weight arrays
- Input validation for both numeric (legacy) and alphanumeric (new) formats
- Configurable character set (blacklist I, O, Q, U, F by default to avoid visual ambiguity)
- Masked format support: `XX.XXX.XXX/YYYY-ZZ`

**Generate Raw Alphanumeric CNPJ:**

```go
package main

import (
	"fmt"
	"github.com/mayckol/brfiscalfaker/pkg/br_documents/v2"
)

func main() {
	// Generate a valid raw alphanumeric CNPJ
	cnpj := v2.CNPJ()
	fmt.Println("CNPJ:", cnpj) // Output: AB1CD2E3F4GH90
}
```

**Generate Masked CNPJ:**

```go
package main

import (
	"fmt"
	"github.com/mayckol/brfiscalfaker/pkg/br_documents/v2"
)

func main() {
	// Generate a masked alphanumeric CNPJ
	cnpj := v2.CNPJ(v2.CNPJv2Config{Masked: true})
	fmt.Println("CNPJ:", cnpj) // Output: AB.1CD.2E3/F4GH-90
}
```

**Allow Ambiguous Letters (I, O, Q, U, F):**

```go
package main

import (
	"fmt"
	"github.com/mayckol/brfiscalfaker/pkg/br_documents/v2"
)

func main() {
	// Generate CNPJ allowing previously blacklisted letters
	cnpj := v2.CNPJ(v2.CNPJv2Config{AllowAmbiguousLetters: true})
	fmt.Println("CNPJ:", cnpj) // May contain I, O, Q, U, or F
	
	// Combine with masked format
	maskedCNPJ := v2.CNPJ(v2.CNPJv2Config{
		Masked:                true,
		AllowAmbiguousLetters: true,
	})
	fmt.Println("Masked CNPJ:", maskedCNPJ)
}
```

**Validate CNPJ (v1 and v2 formats):**

```go
package main

import (
	"fmt"
	"github.com/mayckol/brfiscalfaker/pkg/br_documents/v2"
)

func main() {
	// Validate alphanumeric CNPJ
	cnpj := v2.CNPJ()
	if v2.ValidateCNPJ(cnpj) {
		fmt.Println("Valid CNPJ:", cnpj)
	}
	
	// Validate masked CNPJ (automatically sanitized)
	maskedCNPJ := "AB.1CD.2E3/F4GH-90"
	if v2.ValidateCNPJ(maskedCNPJ) {
		fmt.Println("Valid masked CNPJ")
	}
	
	// Backward compatibility: numeric-only CNPJs also validate
	numericCNPJ := "12345678901234"
	if v2.ValidateCNPJ(numericCNPJ) {
		fmt.Println("Valid numeric CNPJ (backward compatible)")
	}
}
```

**Parameters:**

- **v2.CNPJ()**
  - **Config** (`CNPJv2Config`): Optional parameter to customize generation.
    - `Masked`: `bool`. If `true`, format as `XX.XXX.XXX/YYYY-ZZ`. Default is `false` (raw 14 characters).
    - `AllowAmbiguousLetters`: `bool`. If `true`, allow letters I, O, Q, U, F in generation. Default is `false` (excluded to reduce visual ambiguity).

- **v2.ValidateCNPJ()**
  - Accepts both raw and masked formats.
  - Handles mixed case and whitespace automatically.
  - Returns `true` if CNPJ passes structural and algorithmic validation (check digits).
  - Works with both numeric (legacy) and alphanumeric (v2) formats.

**Status:** This module is ready for testing and development. The new alphanumeric CNPJ format becomes official in July 2026.

## Docker

You can also run the application using Docker.
   ```shell
   docker pull mayckol/bfiscalfaker
   ```
   ```shell
   docker run mayckol/bfiscalfaker --type NFe --cpf 12345678900 --cnpj 12345678901234
   ```

## Contributing

Contributions are welcome! Feel free to open issues or submit pull requests to help improve this project.

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for more information.